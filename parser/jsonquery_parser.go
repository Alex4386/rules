// Code generated from JsonQuery.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // JsonQuery

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type JsonQueryParser struct {
	*antlr.BaseParser
}

var JsonQueryParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jsonqueryParserInit() {
	staticData := &JsonQueryParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "')'", "'pr'", "'.'", "'-'", "'['", "']'", "", "", "", "'null'",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "'\\n'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "NOT", "LOGICAL_OPERATOR", "BOOLEAN",
		"NULL", "IN", "EQ", "NE", "GT", "LT", "GE", "LE", "CO", "SW", "EW",
		"MT", "ATTRNAME", "VERSION", "STRING", "REGEX", "IP_ADDRESS", "IP_CIDR",
		"DOUBLE", "INT", "EXP", "NEWLINE", "COMMA", "SP",
	}
	staticData.RuleNames = []string{
		"query", "attrPath", "subAttr", "value", "regexValue", "ipValue", "listIPs",
		"subListOfIPs", "listStrings", "subListOfStrings", "listDoubles", "subListOfDoubles",
		"listInts", "subListOfInts",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 34, 150, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 3, 0, 31, 8,
		0, 1, 0, 3, 0, 34, 8, 0, 1, 0, 1, 0, 3, 0, 38, 8, 0, 1, 0, 1, 0, 3, 0,
		42, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 3, 0, 68, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 75, 8, 0,
		10, 0, 12, 0, 78, 9, 0, 1, 1, 1, 1, 3, 1, 82, 8, 1, 1, 2, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 93, 8, 3, 1, 3, 1, 3, 3, 3, 97,
		8, 3, 1, 3, 1, 3, 1, 3, 3, 3, 102, 8, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 118, 8, 7,
		1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 128, 8, 9, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 138, 8, 11, 1,
		12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 148, 8, 13,
		1, 13, 0, 1, 0, 14, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		0, 3, 1, 0, 12, 21, 1, 0, 12, 14, 1, 0, 27, 28, 159, 0, 67, 1, 0, 0, 0,
		2, 79, 1, 0, 0, 0, 4, 83, 1, 0, 0, 0, 6, 101, 1, 0, 0, 0, 8, 103, 1, 0,
		0, 0, 10, 105, 1, 0, 0, 0, 12, 107, 1, 0, 0, 0, 14, 117, 1, 0, 0, 0, 16,
		119, 1, 0, 0, 0, 18, 127, 1, 0, 0, 0, 20, 129, 1, 0, 0, 0, 22, 137, 1,
		0, 0, 0, 24, 139, 1, 0, 0, 0, 26, 147, 1, 0, 0, 0, 28, 30, 6, 0, -1, 0,
		29, 31, 5, 8, 0, 0, 30, 29, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 33, 1,
		0, 0, 0, 32, 34, 5, 34, 0, 0, 33, 32, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34,
		35, 1, 0, 0, 0, 35, 37, 5, 1, 0, 0, 36, 38, 5, 34, 0, 0, 37, 36, 1, 0,
		0, 0, 37, 38, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 41, 3, 0, 0, 0, 40, 42,
		5, 34, 0, 0, 41, 40, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0,
		43, 44, 5, 2, 0, 0, 44, 68, 1, 0, 0, 0, 45, 46, 3, 2, 1, 0, 46, 47, 5,
		34, 0, 0, 47, 48, 5, 3, 0, 0, 48, 68, 1, 0, 0, 0, 49, 50, 3, 2, 1, 0, 50,
		51, 5, 34, 0, 0, 51, 52, 7, 0, 0, 0, 52, 53, 5, 34, 0, 0, 53, 54, 3, 6,
		3, 0, 54, 68, 1, 0, 0, 0, 55, 56, 3, 2, 1, 0, 56, 57, 5, 34, 0, 0, 57,
		58, 5, 22, 0, 0, 58, 59, 5, 34, 0, 0, 59, 60, 3, 8, 4, 0, 60, 68, 1, 0,
		0, 0, 61, 62, 3, 2, 1, 0, 62, 63, 5, 34, 0, 0, 63, 64, 7, 1, 0, 0, 64,
		65, 5, 34, 0, 0, 65, 66, 3, 10, 5, 0, 66, 68, 1, 0, 0, 0, 67, 28, 1, 0,
		0, 0, 67, 45, 1, 0, 0, 0, 67, 49, 1, 0, 0, 0, 67, 55, 1, 0, 0, 0, 67, 61,
		1, 0, 0, 0, 68, 76, 1, 0, 0, 0, 69, 70, 10, 5, 0, 0, 70, 71, 5, 34, 0,
		0, 71, 72, 5, 9, 0, 0, 72, 73, 5, 34, 0, 0, 73, 75, 3, 0, 0, 6, 74, 69,
		1, 0, 0, 0, 75, 78, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0,
		77, 1, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 79, 81, 5, 23, 0, 0, 80, 82, 3,
		4, 2, 0, 81, 80, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 3, 1, 0, 0, 0, 83,
		84, 5, 4, 0, 0, 84, 85, 3, 2, 1, 0, 85, 5, 1, 0, 0, 0, 86, 102, 5, 10,
		0, 0, 87, 102, 5, 11, 0, 0, 88, 102, 5, 24, 0, 0, 89, 102, 5, 25, 0, 0,
		90, 102, 5, 29, 0, 0, 91, 93, 5, 5, 0, 0, 92, 91, 1, 0, 0, 0, 92, 93, 1,
		0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 96, 5, 30, 0, 0, 95, 97, 5, 31, 0, 0,
		96, 95, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 102, 1, 0, 0, 0, 98, 102, 3,
		24, 12, 0, 99, 102, 3, 20, 10, 0, 100, 102, 3, 16, 8, 0, 101, 86, 1, 0,
		0, 0, 101, 87, 1, 0, 0, 0, 101, 88, 1, 0, 0, 0, 101, 89, 1, 0, 0, 0, 101,
		90, 1, 0, 0, 0, 101, 92, 1, 0, 0, 0, 101, 98, 1, 0, 0, 0, 101, 99, 1, 0,
		0, 0, 101, 100, 1, 0, 0, 0, 102, 7, 1, 0, 0, 0, 103, 104, 5, 26, 0, 0,
		104, 9, 1, 0, 0, 0, 105, 106, 7, 2, 0, 0, 106, 11, 1, 0, 0, 0, 107, 108,
		5, 6, 0, 0, 108, 109, 3, 14, 7, 0, 109, 13, 1, 0, 0, 0, 110, 111, 3, 10,
		5, 0, 111, 112, 5, 33, 0, 0, 112, 113, 3, 14, 7, 0, 113, 118, 1, 0, 0,
		0, 114, 115, 3, 10, 5, 0, 115, 116, 5, 7, 0, 0, 116, 118, 1, 0, 0, 0, 117,
		110, 1, 0, 0, 0, 117, 114, 1, 0, 0, 0, 118, 15, 1, 0, 0, 0, 119, 120, 5,
		6, 0, 0, 120, 121, 3, 18, 9, 0, 121, 17, 1, 0, 0, 0, 122, 123, 5, 25, 0,
		0, 123, 124, 5, 33, 0, 0, 124, 128, 3, 18, 9, 0, 125, 126, 5, 25, 0, 0,
		126, 128, 5, 7, 0, 0, 127, 122, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 128,
		19, 1, 0, 0, 0, 129, 130, 5, 6, 0, 0, 130, 131, 3, 22, 11, 0, 131, 21,
		1, 0, 0, 0, 132, 133, 5, 29, 0, 0, 133, 134, 5, 33, 0, 0, 134, 138, 3,
		22, 11, 0, 135, 136, 5, 29, 0, 0, 136, 138, 5, 7, 0, 0, 137, 132, 1, 0,
		0, 0, 137, 135, 1, 0, 0, 0, 138, 23, 1, 0, 0, 0, 139, 140, 5, 6, 0, 0,
		140, 141, 3, 26, 13, 0, 141, 25, 1, 0, 0, 0, 142, 143, 5, 30, 0, 0, 143,
		144, 5, 33, 0, 0, 144, 148, 3, 26, 13, 0, 145, 146, 5, 30, 0, 0, 146, 148,
		5, 7, 0, 0, 147, 142, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 148, 27, 1, 0,
		0, 0, 14, 30, 33, 37, 41, 67, 76, 81, 92, 96, 101, 117, 127, 137, 147,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// JsonQueryParserInit initializes any static state used to implement JsonQueryParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJsonQueryParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JsonQueryParserInit() {
	staticData := &JsonQueryParserStaticData
	staticData.once.Do(jsonqueryParserInit)
}

// NewJsonQueryParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJsonQueryParser(input antlr.TokenStream) *JsonQueryParser {
	JsonQueryParserInit()
	this := new(JsonQueryParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &JsonQueryParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "JsonQuery.g4"

	return this
}

// JsonQueryParser tokens.
const (
	JsonQueryParserEOF              = antlr.TokenEOF
	JsonQueryParserT__0             = 1
	JsonQueryParserT__1             = 2
	JsonQueryParserT__2             = 3
	JsonQueryParserT__3             = 4
	JsonQueryParserT__4             = 5
	JsonQueryParserT__5             = 6
	JsonQueryParserT__6             = 7
	JsonQueryParserNOT              = 8
	JsonQueryParserLOGICAL_OPERATOR = 9
	JsonQueryParserBOOLEAN          = 10
	JsonQueryParserNULL             = 11
	JsonQueryParserIN               = 12
	JsonQueryParserEQ               = 13
	JsonQueryParserNE               = 14
	JsonQueryParserGT               = 15
	JsonQueryParserLT               = 16
	JsonQueryParserGE               = 17
	JsonQueryParserLE               = 18
	JsonQueryParserCO               = 19
	JsonQueryParserSW               = 20
	JsonQueryParserEW               = 21
	JsonQueryParserMT               = 22
	JsonQueryParserATTRNAME         = 23
	JsonQueryParserVERSION          = 24
	JsonQueryParserSTRING           = 25
	JsonQueryParserREGEX            = 26
	JsonQueryParserIP_ADDRESS       = 27
	JsonQueryParserIP_CIDR          = 28
	JsonQueryParserDOUBLE           = 29
	JsonQueryParserINT              = 30
	JsonQueryParserEXP              = 31
	JsonQueryParserNEWLINE          = 32
	JsonQueryParserCOMMA            = 33
	JsonQueryParserSP               = 34
)

// JsonQueryParser rules.
const (
	JsonQueryParserRULE_query            = 0
	JsonQueryParserRULE_attrPath         = 1
	JsonQueryParserRULE_subAttr          = 2
	JsonQueryParserRULE_value            = 3
	JsonQueryParserRULE_regexValue       = 4
	JsonQueryParserRULE_ipValue          = 5
	JsonQueryParserRULE_listIPs          = 6
	JsonQueryParserRULE_subListOfIPs     = 7
	JsonQueryParserRULE_listStrings      = 8
	JsonQueryParserRULE_subListOfStrings = 9
	JsonQueryParserRULE_listDoubles      = 10
	JsonQueryParserRULE_subListOfDoubles = 11
	JsonQueryParserRULE_listInts         = 12
	JsonQueryParserRULE_subListOfInts    = 13
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_query
	return p
}

func InitEmptyQueryContext(p *QueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_query
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) CopyAll(ctx *QueryContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CompareExpContext struct {
	QueryContext
	op antlr.Token
}

func NewCompareExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareExpContext {
	var p = new(CompareExpContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *CompareExpContext) GetOp() antlr.Token { return s.op }

func (s *CompareExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareExpContext) AttrPath() IAttrPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *CompareExpContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserSP)
}

func (s *CompareExpContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, i)
}

func (s *CompareExpContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *CompareExpContext) EQ() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserEQ, 0)
}

func (s *CompareExpContext) NE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserNE, 0)
}

func (s *CompareExpContext) GT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserGT, 0)
}

func (s *CompareExpContext) LT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLT, 0)
}

func (s *CompareExpContext) GE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserGE, 0)
}

func (s *CompareExpContext) LE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLE, 0)
}

func (s *CompareExpContext) CO() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCO, 0)
}

func (s *CompareExpContext) SW() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSW, 0)
}

func (s *CompareExpContext) EW() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserEW, 0)
}

func (s *CompareExpContext) IN() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserIN, 0)
}

func (s *CompareExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitCompareExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type RegexExpContext struct {
	QueryContext
	op antlr.Token
}

func NewRegexExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RegexExpContext {
	var p = new(RegexExpContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *RegexExpContext) GetOp() antlr.Token { return s.op }

func (s *RegexExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *RegexExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexExpContext) AttrPath() IAttrPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *RegexExpContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserSP)
}

func (s *RegexExpContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, i)
}

func (s *RegexExpContext) RegexValue() IRegexValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegexValueContext)
}

func (s *RegexExpContext) MT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserMT, 0)
}

func (s *RegexExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitRegexExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExpContext struct {
	QueryContext
}

func NewParenExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExpContext {
	var p = new(ParenExpContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *ParenExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExpContext) Query() IQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *ParenExpContext) NOT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserNOT, 0)
}

func (s *ParenExpContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserSP)
}

func (s *ParenExpContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, i)
}

func (s *ParenExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitParenExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type PresentExpContext struct {
	QueryContext
}

func NewPresentExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PresentExpContext {
	var p = new(PresentExpContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *PresentExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PresentExpContext) AttrPath() IAttrPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *PresentExpContext) SP() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, 0)
}

func (s *PresentExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitPresentExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type IpCompareExpContext struct {
	QueryContext
	op antlr.Token
}

func NewIpCompareExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IpCompareExpContext {
	var p = new(IpCompareExpContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *IpCompareExpContext) GetOp() antlr.Token { return s.op }

func (s *IpCompareExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *IpCompareExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IpCompareExpContext) AttrPath() IAttrPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *IpCompareExpContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserSP)
}

func (s *IpCompareExpContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, i)
}

func (s *IpCompareExpContext) IpValue() IIpValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIpValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIpValueContext)
}

func (s *IpCompareExpContext) EQ() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserEQ, 0)
}

func (s *IpCompareExpContext) NE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserNE, 0)
}

func (s *IpCompareExpContext) IN() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserIN, 0)
}

func (s *IpCompareExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitIpCompareExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicalExpContext struct {
	QueryContext
}

func NewLogicalExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalExpContext {
	var p = new(LogicalExpContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *LogicalExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalExpContext) AllQuery() []IQueryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IQueryContext); ok {
			len++
		}
	}

	tst := make([]IQueryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IQueryContext); ok {
			tst[i] = t.(IQueryContext)
			i++
		}
	}

	return tst
}

func (s *LogicalExpContext) Query(i int) IQueryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *LogicalExpContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserSP)
}

func (s *LogicalExpContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, i)
}

func (s *LogicalExpContext) LOGICAL_OPERATOR() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLOGICAL_OPERATOR, 0)
}

func (s *LogicalExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitLogicalExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) Query() (localctx IQueryContext) {
	return p.query(0)
}

func (p *JsonQueryParser) query(_p int) (localctx IQueryContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewQueryContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IQueryContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, JsonQueryParserRULE_query, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserNOT {
			{
				p.SetState(29)
				p.Match(JsonQueryParserNOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserSP {
			{
				p.SetState(32)
				p.Match(JsonQueryParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(35)
			p.Match(JsonQueryParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(37)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(36)
				p.Match(JsonQueryParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(39)
			p.query(0)
		}
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserSP {
			{
				p.SetState(40)
				p.Match(JsonQueryParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(43)
			p.Match(JsonQueryParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewPresentExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(45)
			p.AttrPath()
		}
		{
			p.SetState(46)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.Match(JsonQueryParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewCompareExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(49)
			p.AttrPath()
		}
		{
			p.SetState(50)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CompareExpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4190208) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CompareExpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(52)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.Value()
		}

	case 4:
		localctx = NewRegexExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(55)
			p.AttrPath()
		}
		{
			p.SetState(56)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)

			var _m = p.Match(JsonQueryParserMT)

			localctx.(*RegexExpContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)
			p.RegexValue()
		}

	case 5:
		localctx = NewIpCompareExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(61)
			p.AttrPath()
		}
		{
			p.SetState(62)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*IpCompareExpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&28672) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*IpCompareExpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(64)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.IpValue()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLogicalExpContext(p, NewQueryContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, JsonQueryParserRULE_query)
			p.SetState(69)

			if !(p.Precpred(p.GetParserRuleContext(), 5)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				goto errorExit
			}
			{
				p.SetState(70)
				p.Match(JsonQueryParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(71)
				p.Match(JsonQueryParserLOGICAL_OPERATOR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(72)
				p.Match(JsonQueryParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(73)
				p.query(6)
			}

		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttrPathContext is an interface to support dynamic dispatch.
type IAttrPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ATTRNAME() antlr.TerminalNode
	SubAttr() ISubAttrContext

	// IsAttrPathContext differentiates from other interfaces.
	IsAttrPathContext()
}

type AttrPathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrPathContext() *AttrPathContext {
	var p = new(AttrPathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_attrPath
	return p
}

func InitEmptyAttrPathContext(p *AttrPathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_attrPath
}

func (*AttrPathContext) IsAttrPathContext() {}

func NewAttrPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrPathContext {
	var p = new(AttrPathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_attrPath

	return p
}

func (s *AttrPathContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrPathContext) ATTRNAME() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserATTRNAME, 0)
}

func (s *AttrPathContext) SubAttr() ISubAttrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubAttrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubAttrContext)
}

func (s *AttrPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrPathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitAttrPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) AttrPath() (localctx IAttrPathContext) {
	localctx = NewAttrPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JsonQueryParserRULE_attrPath)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(JsonQueryParserATTRNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JsonQueryParserT__3 {
		{
			p.SetState(80)
			p.SubAttr()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISubAttrContext is an interface to support dynamic dispatch.
type ISubAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AttrPath() IAttrPathContext

	// IsSubAttrContext differentiates from other interfaces.
	IsSubAttrContext()
}

type SubAttrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubAttrContext() *SubAttrContext {
	var p = new(SubAttrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subAttr
	return p
}

func InitEmptySubAttrContext(p *SubAttrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subAttr
}

func (*SubAttrContext) IsSubAttrContext() {}

func NewSubAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubAttrContext {
	var p = new(SubAttrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_subAttr

	return p
}

func (s *SubAttrContext) GetParser() antlr.Parser { return s.parser }

func (s *SubAttrContext) AttrPath() IAttrPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *SubAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubAttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubAttrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitSubAttr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) SubAttr() (localctx ISubAttrContext) {
	localctx = NewSubAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JsonQueryParserRULE_subAttr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.Match(JsonQueryParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
		p.AttrPath()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyAll(ctx *ValueContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListOfDoublesContext struct {
	ValueContext
}

func NewListOfDoublesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListOfDoublesContext {
	var p = new(ListOfDoublesContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ListOfDoublesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOfDoublesContext) ListDoubles() IListDoublesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListDoublesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListDoublesContext)
}

func (s *ListOfDoublesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListOfDoubles(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListOfStringsContext struct {
	ValueContext
}

func NewListOfStringsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListOfStringsContext {
	var p = new(ListOfStringsContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ListOfStringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOfStringsContext) ListStrings() IListStringsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListStringsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListStringsContext)
}

func (s *ListOfStringsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListOfStrings(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanContext struct {
	ValueContext
}

func NewBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanContext {
	var p = new(BooleanContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserBOOLEAN, 0)
}

func (s *BooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullContext struct {
	ValueContext
}

func NewNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullContext {
	var p = new(NullContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *NullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullContext) NULL() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserNULL, 0)
}

func (s *NullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitNull(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringContext struct {
	ValueContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSTRING, 0)
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type DoubleContext struct {
	ValueContext
}

func NewDoubleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DoubleContext {
	var p = new(DoubleContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *DoubleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserDOUBLE, 0)
}

func (s *DoubleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitDouble(s)

	default:
		return t.VisitChildren(s)
	}
}

type VersionContext struct {
	ValueContext
}

func NewVersionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VersionContext {
	var p = new(VersionContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *VersionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionContext) VERSION() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserVERSION, 0)
}

func (s *VersionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitVersion(s)

	default:
		return t.VisitChildren(s)
	}
}

type LongContext struct {
	ValueContext
}

func NewLongContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LongContext {
	var p = new(LongContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *LongContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LongContext) INT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserINT, 0)
}

func (s *LongContext) EXP() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserEXP, 0)
}

func (s *LongContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitLong(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListOfIntsContext struct {
	ValueContext
}

func NewListOfIntsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListOfIntsContext {
	var p = new(ListOfIntsContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ListOfIntsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOfIntsContext) ListInts() IListIntsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListIntsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListIntsContext)
}

func (s *ListOfIntsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListOfInts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JsonQueryParserRULE_value)
	var _la int

	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Match(JsonQueryParserBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(87)
			p.Match(JsonQueryParserNULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewVersionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(88)
			p.Match(JsonQueryParserVERSION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(89)
			p.Match(JsonQueryParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewDoubleContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(90)
			p.Match(JsonQueryParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewLongContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserT__4 {
			{
				p.SetState(91)
				p.Match(JsonQueryParserT__4)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(94)
			p.Match(JsonQueryParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(95)
				p.Match(JsonQueryParserEXP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 7:
		localctx = NewListOfIntsContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(98)
			p.ListInts()
		}

	case 8:
		localctx = NewListOfDoublesContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(99)
			p.ListDoubles()
		}

	case 9:
		localctx = NewListOfStringsContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(100)
			p.ListStrings()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRegexValueContext is an interface to support dynamic dispatch.
type IRegexValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REGEX() antlr.TerminalNode

	// IsRegexValueContext differentiates from other interfaces.
	IsRegexValueContext()
}

type RegexValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegexValueContext() *RegexValueContext {
	var p = new(RegexValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_regexValue
	return p
}

func InitEmptyRegexValueContext(p *RegexValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_regexValue
}

func (*RegexValueContext) IsRegexValueContext() {}

func NewRegexValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegexValueContext {
	var p = new(RegexValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_regexValue

	return p
}

func (s *RegexValueContext) GetParser() antlr.Parser { return s.parser }

func (s *RegexValueContext) REGEX() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserREGEX, 0)
}

func (s *RegexValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegexValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitRegexValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) RegexValue() (localctx IRegexValueContext) {
	localctx = NewRegexValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JsonQueryParserRULE_regexValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(JsonQueryParserREGEX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIpValueContext is an interface to support dynamic dispatch.
type IIpValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IP_ADDRESS() antlr.TerminalNode
	IP_CIDR() antlr.TerminalNode

	// IsIpValueContext differentiates from other interfaces.
	IsIpValueContext()
}

type IpValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpValueContext() *IpValueContext {
	var p = new(IpValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_ipValue
	return p
}

func InitEmptyIpValueContext(p *IpValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_ipValue
}

func (*IpValueContext) IsIpValueContext() {}

func NewIpValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IpValueContext {
	var p = new(IpValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_ipValue

	return p
}

func (s *IpValueContext) GetParser() antlr.Parser { return s.parser }

func (s *IpValueContext) IP_ADDRESS() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserIP_ADDRESS, 0)
}

func (s *IpValueContext) IP_CIDR() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserIP_CIDR, 0)
}

func (s *IpValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IpValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IpValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitIpValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) IpValue() (localctx IIpValueContext) {
	localctx = NewIpValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JsonQueryParserRULE_ipValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		_la = p.GetTokenStream().LA(1)

		if !(_la == JsonQueryParserIP_ADDRESS || _la == JsonQueryParserIP_CIDR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListIPsContext is an interface to support dynamic dispatch.
type IListIPsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SubListOfIPs() ISubListOfIPsContext

	// IsListIPsContext differentiates from other interfaces.
	IsListIPsContext()
}

type ListIPsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListIPsContext() *ListIPsContext {
	var p = new(ListIPsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listIPs
	return p
}

func InitEmptyListIPsContext(p *ListIPsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listIPs
}

func (*ListIPsContext) IsListIPsContext() {}

func NewListIPsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListIPsContext {
	var p = new(ListIPsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_listIPs

	return p
}

func (s *ListIPsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListIPsContext) SubListOfIPs() ISubListOfIPsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfIPsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfIPsContext)
}

func (s *ListIPsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListIPsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListIPsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListIPs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) ListIPs() (localctx IListIPsContext) {
	localctx = NewListIPsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JsonQueryParserRULE_listIPs)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Match(JsonQueryParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)
		p.SubListOfIPs()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISubListOfIPsContext is an interface to support dynamic dispatch.
type ISubListOfIPsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IpValue() IIpValueContext
	COMMA() antlr.TerminalNode
	SubListOfIPs() ISubListOfIPsContext

	// IsSubListOfIPsContext differentiates from other interfaces.
	IsSubListOfIPsContext()
}

type SubListOfIPsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubListOfIPsContext() *SubListOfIPsContext {
	var p = new(SubListOfIPsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfIPs
	return p
}

func InitEmptySubListOfIPsContext(p *SubListOfIPsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfIPs
}

func (*SubListOfIPsContext) IsSubListOfIPsContext() {}

func NewSubListOfIPsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubListOfIPsContext {
	var p = new(SubListOfIPsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_subListOfIPs

	return p
}

func (s *SubListOfIPsContext) GetParser() antlr.Parser { return s.parser }

func (s *SubListOfIPsContext) IpValue() IIpValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIpValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIpValueContext)
}

func (s *SubListOfIPsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCOMMA, 0)
}

func (s *SubListOfIPsContext) SubListOfIPs() ISubListOfIPsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfIPsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfIPsContext)
}

func (s *SubListOfIPsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubListOfIPsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubListOfIPsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitSubListOfIPs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) SubListOfIPs() (localctx ISubListOfIPsContext) {
	localctx = NewSubListOfIPsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JsonQueryParserRULE_subListOfIPs)
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.IpValue()
		}
		{
			p.SetState(111)
			p.Match(JsonQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.SubListOfIPs()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(114)
			p.IpValue()
		}
		{
			p.SetState(115)
			p.Match(JsonQueryParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListStringsContext is an interface to support dynamic dispatch.
type IListStringsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SubListOfStrings() ISubListOfStringsContext

	// IsListStringsContext differentiates from other interfaces.
	IsListStringsContext()
}

type ListStringsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListStringsContext() *ListStringsContext {
	var p = new(ListStringsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listStrings
	return p
}

func InitEmptyListStringsContext(p *ListStringsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listStrings
}

func (*ListStringsContext) IsListStringsContext() {}

func NewListStringsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListStringsContext {
	var p = new(ListStringsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_listStrings

	return p
}

func (s *ListStringsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListStringsContext) SubListOfStrings() ISubListOfStringsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfStringsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfStringsContext)
}

func (s *ListStringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListStringsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListStringsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListStrings(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) ListStrings() (localctx IListStringsContext) {
	localctx = NewListStringsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JsonQueryParserRULE_listStrings)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(JsonQueryParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(120)
		p.SubListOfStrings()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISubListOfStringsContext is an interface to support dynamic dispatch.
type ISubListOfStringsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	SubListOfStrings() ISubListOfStringsContext

	// IsSubListOfStringsContext differentiates from other interfaces.
	IsSubListOfStringsContext()
}

type SubListOfStringsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubListOfStringsContext() *SubListOfStringsContext {
	var p = new(SubListOfStringsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfStrings
	return p
}

func InitEmptySubListOfStringsContext(p *SubListOfStringsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfStrings
}

func (*SubListOfStringsContext) IsSubListOfStringsContext() {}

func NewSubListOfStringsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubListOfStringsContext {
	var p = new(SubListOfStringsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_subListOfStrings

	return p
}

func (s *SubListOfStringsContext) GetParser() antlr.Parser { return s.parser }

func (s *SubListOfStringsContext) STRING() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSTRING, 0)
}

func (s *SubListOfStringsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCOMMA, 0)
}

func (s *SubListOfStringsContext) SubListOfStrings() ISubListOfStringsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfStringsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfStringsContext)
}

func (s *SubListOfStringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubListOfStringsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubListOfStringsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitSubListOfStrings(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) SubListOfStrings() (localctx ISubListOfStringsContext) {
	localctx = NewSubListOfStringsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, JsonQueryParserRULE_subListOfStrings)
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.Match(JsonQueryParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Match(JsonQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.SubListOfStrings()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.Match(JsonQueryParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.Match(JsonQueryParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListDoublesContext is an interface to support dynamic dispatch.
type IListDoublesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SubListOfDoubles() ISubListOfDoublesContext

	// IsListDoublesContext differentiates from other interfaces.
	IsListDoublesContext()
}

type ListDoublesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListDoublesContext() *ListDoublesContext {
	var p = new(ListDoublesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listDoubles
	return p
}

func InitEmptyListDoublesContext(p *ListDoublesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listDoubles
}

func (*ListDoublesContext) IsListDoublesContext() {}

func NewListDoublesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListDoublesContext {
	var p = new(ListDoublesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_listDoubles

	return p
}

func (s *ListDoublesContext) GetParser() antlr.Parser { return s.parser }

func (s *ListDoublesContext) SubListOfDoubles() ISubListOfDoublesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfDoublesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfDoublesContext)
}

func (s *ListDoublesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListDoublesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListDoublesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListDoubles(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) ListDoubles() (localctx IListDoublesContext) {
	localctx = NewListDoublesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, JsonQueryParserRULE_listDoubles)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Match(JsonQueryParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)
		p.SubListOfDoubles()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISubListOfDoublesContext is an interface to support dynamic dispatch.
type ISubListOfDoublesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOUBLE() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	SubListOfDoubles() ISubListOfDoublesContext

	// IsSubListOfDoublesContext differentiates from other interfaces.
	IsSubListOfDoublesContext()
}

type SubListOfDoublesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubListOfDoublesContext() *SubListOfDoublesContext {
	var p = new(SubListOfDoublesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfDoubles
	return p
}

func InitEmptySubListOfDoublesContext(p *SubListOfDoublesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfDoubles
}

func (*SubListOfDoublesContext) IsSubListOfDoublesContext() {}

func NewSubListOfDoublesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubListOfDoublesContext {
	var p = new(SubListOfDoublesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_subListOfDoubles

	return p
}

func (s *SubListOfDoublesContext) GetParser() antlr.Parser { return s.parser }

func (s *SubListOfDoublesContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserDOUBLE, 0)
}

func (s *SubListOfDoublesContext) COMMA() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCOMMA, 0)
}

func (s *SubListOfDoublesContext) SubListOfDoubles() ISubListOfDoublesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfDoublesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfDoublesContext)
}

func (s *SubListOfDoublesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubListOfDoublesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubListOfDoublesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitSubListOfDoubles(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) SubListOfDoubles() (localctx ISubListOfDoublesContext) {
	localctx = NewSubListOfDoublesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, JsonQueryParserRULE_subListOfDoubles)
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(132)
			p.Match(JsonQueryParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.Match(JsonQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.SubListOfDoubles()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
			p.Match(JsonQueryParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.Match(JsonQueryParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListIntsContext is an interface to support dynamic dispatch.
type IListIntsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SubListOfInts() ISubListOfIntsContext

	// IsListIntsContext differentiates from other interfaces.
	IsListIntsContext()
}

type ListIntsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListIntsContext() *ListIntsContext {
	var p = new(ListIntsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listInts
	return p
}

func InitEmptyListIntsContext(p *ListIntsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listInts
}

func (*ListIntsContext) IsListIntsContext() {}

func NewListIntsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListIntsContext {
	var p = new(ListIntsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_listInts

	return p
}

func (s *ListIntsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListIntsContext) SubListOfInts() ISubListOfIntsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfIntsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfIntsContext)
}

func (s *ListIntsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListIntsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListIntsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListInts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) ListInts() (localctx IListIntsContext) {
	localctx = NewListIntsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, JsonQueryParserRULE_listInts)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(JsonQueryParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.SubListOfInts()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISubListOfIntsContext is an interface to support dynamic dispatch.
type ISubListOfIntsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	SubListOfInts() ISubListOfIntsContext

	// IsSubListOfIntsContext differentiates from other interfaces.
	IsSubListOfIntsContext()
}

type SubListOfIntsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubListOfIntsContext() *SubListOfIntsContext {
	var p = new(SubListOfIntsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfInts
	return p
}

func InitEmptySubListOfIntsContext(p *SubListOfIntsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfInts
}

func (*SubListOfIntsContext) IsSubListOfIntsContext() {}

func NewSubListOfIntsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubListOfIntsContext {
	var p = new(SubListOfIntsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_subListOfInts

	return p
}

func (s *SubListOfIntsContext) GetParser() antlr.Parser { return s.parser }

func (s *SubListOfIntsContext) INT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserINT, 0)
}

func (s *SubListOfIntsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCOMMA, 0)
}

func (s *SubListOfIntsContext) SubListOfInts() ISubListOfIntsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfIntsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfIntsContext)
}

func (s *SubListOfIntsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubListOfIntsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubListOfIntsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitSubListOfInts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) SubListOfInts() (localctx ISubListOfIntsContext) {
	localctx = NewSubListOfIntsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, JsonQueryParserRULE_subListOfInts)
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(142)
			p.Match(JsonQueryParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.Match(JsonQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)
			p.SubListOfInts()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(145)
			p.Match(JsonQueryParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.Match(JsonQueryParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *JsonQueryParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *QueryContext = nil
		if localctx != nil {
			t = localctx.(*QueryContext)
		}
		return p.Query_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *JsonQueryParser) Query_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
