package parser

import (
	"fmt"
	"sync"

	"github.com/antlr4-go/antlr/v4"
	regexp "github.com/wasilibs/go-re2"
)

type EvaluatorOptions struct {
	// DisableASTRegexCache makes the evaluator recompile regex literals on
	// every Process call instead of compiling them once when the rule is parsed.
	DisableASTRegexCache bool
}

// syntaxErrors collects lexer and parser errors instead of printing them.
type syntaxErrors struct {
	*antlr.DefaultErrorListener
	first error
}

func (l *syntaxErrors) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	if l.first == nil {
		l.first = fmt.Errorf("syntax error at %d:%d: %s", line, column, msg)
	}
}

// Evaluator is safe for concurrent use once created. With concurrent Process
// calls LastDebugErr reports whichever call finished last.
type Evaluator struct {
	mu           sync.Mutex // guards lastDebugErr
	lastDebugErr error

	rule    string
	tree    antlr.ParseTree
	doPanic bool

	regexCache map[*RegexValueContext]*regexp.Regexp

	testHookPanic func()
}

func NewEvaluator(rule string) (*Evaluator, error) {
	return NewEvaluatorWithOptions(rule, EvaluatorOptions{})
}

func NewEvaluatorWithOptions(rule string, opts EvaluatorOptions) (ret *Evaluator, retErr error) {
	// antlr lib has panics for exceptions so we have to put a recover here
	// in the unlikely case there is an exception
	defer func() {
		info := recover()
		if info != nil {
			retErr = fmt.Errorf("%q", info)
		}
	}()
	input := antlr.NewInputStream(rule)
	lex := NewJsonQueryLexer(input)
	errs := &syntaxErrors{DefaultErrorListener: antlr.NewDefaultErrorListener()}
	lex.RemoveErrorListeners()
	lex.AddErrorListener(errs)
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := NewJsonQueryParser(tokens)
	p.RemoveErrorListeners()
	p.AddErrorListener(errs)
	tree := p.Query()
	if errs.first != nil {
		return nil, errs.first
	}
	// the grammar has no EOF anchor, so leftover input is not a parser error
	for tokens.LA(1) == JsonQueryParserSP || tokens.LA(1) == JsonQueryParserNEWLINE {
		tokens.Consume()
	}
	if rest := tokens.LT(1); rest.GetTokenType() != antlr.TokenEOF {
		return nil, fmt.Errorf("syntax error at %d:%d: unexpected %q after end of rule",
			rest.GetLine(), rest.GetColumn(), rest.GetText())
	}

	ev := &Evaluator{
		rule: rule,
		tree: tree,
	}
	if !opts.DisableASTRegexCache {
		ev.regexCache = map[*RegexValueContext]*regexp.Regexp{}
		collectRegexes(tree, ev.regexCache)
	}
	return ev, nil
}

func NewEvaluatorWithPanicOnParseError(rule string) (ret *Evaluator, retErr error) {
	ret, retErr = NewEvaluator(rule)
	if retErr != nil {
		return nil, retErr
	}

	ret.doPanic = true
	return ret, nil
}

func (e *Evaluator) setLastDebugErr(err error) {
	e.mu.Lock()
	e.lastDebugErr = err
	e.mu.Unlock()
}

func (e *Evaluator) Reset() error {
	e.setLastDebugErr(nil)
	return nil
}

func (e *Evaluator) LastDebugErr() error {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.lastDebugErr
}

func (e *Evaluator) Process(items map[string]interface{}) (ret bool, retErr error) {
	e.setLastDebugErr(nil)
	// antlr lib has panics for exceptions so we have to put a recover here
	// in the unlikely case there is an exception
	if !e.doPanic {
		defer func() {
			info := recover()
			if info != nil {
				retErr = fmt.Errorf("%q", info)
				ret = false
			}
		}()
	}

	visitor := NewJsonQueryVisitorImpl(items)
	visitor.regexCache = e.regexCache
	if e.doPanic {
		visitor.doPanic = true
	}

	result := visitor.Visit(e.tree)
	e.setLastDebugErr(visitor.debugErr)
	if e.testHookPanic != nil && !e.doPanic {
		defer e.testHookPanic()
	}
	if result == nil || visitor.err != nil {
		return false, visitor.err
	}

	return result.(bool), visitor.err
}

func Evaluate(rule string, items map[string]interface{}) bool {
	ev, err := NewEvaluator(rule)
	if err != nil {
		return false
	}
	result, _ := ev.Process(items)
	return result
}
