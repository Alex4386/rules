package parser

import (
	"errors"
	"fmt"
	regexp "github.com/wasilibs/go-re2"
	"net"
	"reflect"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type JsonQueryVisitorImpl struct {
	antlr.ParseTreeVisitor

	item map[string]interface{}

	currentOperation Operation
	leftOp           Operand
	rightOp          Operand

	err      error
	debugErr error

	doPanic bool

	// regexCache holds regex literals precompiled at parse time; nil when disabled.
	regexCache map[*RegexValueContext]*regexp.Regexp
}

func NewJsonQueryVisitorImpl(item map[string]interface{}) *JsonQueryVisitorImpl {
	return &JsonQueryVisitorImpl{
		item: item,
	}
}

func (j *JsonQueryVisitorImpl) setDebugErr(err error) {
	if j.doPanic {
		panic(err)
	}
	j.debugErr = err
}

func (j *JsonQueryVisitorImpl) setErr(err error) {
	if j.doPanic {
		panic(err)
	}
	j.err = err
}

func (j *JsonQueryVisitorImpl) hasErr() bool {
	return j.err != nil
}

func (j *JsonQueryVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	if j.hasErr() {
		return false
	}

	switch val := tree.(type) {
	case *LogicalExpContext:
		return val.Accept(j).(bool)
	case *CompareExpContext:
		return val.Accept(j).(bool)
	case *ParenExpContext:
		return val.Accept(j).(bool)
	case *PresentExpContext:
		return val.Accept(j).(bool)
	case *RegexExpContext:
		return val.Accept(j).(bool)
	case *IpCompareExpContext:
		return val.Accept(j).(bool)
	default:
		j.setErr(errors.New("Invalid rule"))
		return false
	}
}

func (j *JsonQueryVisitorImpl) VisitParenExp(ctx *ParenExpContext) interface{} {
	result := ctx.Query().Accept(j).(bool)
	if ctx.NOT() != nil {
		return !result
	}
	return result
}

func (j *JsonQueryVisitorImpl) VisitLogicalExp(ctx *LogicalExpContext) interface{} {
	left := ctx.Query(0).Accept(j).(bool)
	op := ctx.LOGICAL_OPERATOR().GetText()
	if op == "or" {
		if left {
			return left
		}
		return ctx.Query(1).Accept(j).(bool)
	}
	// means it is and
	if !left {
		return left
	}
	return ctx.Query(1).Accept(j).(bool)
}

func (j *JsonQueryVisitorImpl) VisitPresentExp(ctx *PresentExpContext) interface{} {
	ctx.AttrPath().Accept(j)
	return j.leftOp != nil
}

func (j *JsonQueryVisitorImpl) VisitCompareExp(ctx *CompareExpContext) interface{} {
	ctx.AttrPath().Accept(j)
	ctx.Value().Accept(j)
	if j.hasErr() {
		return false
	}
	var apply func(Operand, Operand) (bool, error)
	currentOp := j.currentOperation
	if currentOp == nil {
		j.setErr(ErrInvalidOperation)
		return false
	}
	if _, ok := currentOp.(*IntOperation); ok {
		switch j.leftOp.(type) {
		case float64, float32:
			// JSON numbers are float64: don't truncate 1.9 to compare with 1
			currentOp = &FloatOperation{}
		}
	}

	switch ctx.op.GetTokenType() {
	case JsonQueryParserEQ:
		apply = currentOp.EQ
	case JsonQueryParserNE:
		apply = currentOp.NE
	case JsonQueryParserGT:
		apply = currentOp.GT
	case JsonQueryParserLT:
		apply = currentOp.LT
	case JsonQueryParserLE:
		apply = currentOp.LE
	case JsonQueryParserGE:
		apply = currentOp.GE
	case JsonQueryParserCO:
		apply = currentOp.CO
	case JsonQueryParserSW:
		apply = currentOp.SW
	case JsonQueryParserEW:
		apply = currentOp.EW
	case JsonQueryParserIN:
		apply = currentOp.IN
	default:
		j.setErr(fmt.Errorf("Unknown operation %s", ctx.op.GetText()))
		return false
	}
	defer func() { j.rightOp = nil }()
	ret, err := apply(j.leftOp, j.rightOp)
	if err != nil {
		switch err {
		case ErrInvalidOperation:
			// in case of invalid operation lets rather
			// be conservative and return false because the rule doesn't even make
			// sense. It can be argued that it would be false positive if we were
			// to return true
			j.setErr(err)
			j.setDebugErr(
				newNestedError(err, "Not a valid operation for datatypes").Set(ErrVals{
					"operation":           ctx.op.GetTokenType(),
					"object_path_operand": j.leftOp,
					"rule_operand":        j.rightOp,
				}),
			)
		case ErrEvalOperandMissing:
			j.setDebugErr(
				newNestedError(err, "Eval operand missing in input object").Set(ErrVals{
					"attr_path":           ctx.AttrPath().GetText(),
					"value":               ctx.Value().GetText(),
					"object_path_operand": j.leftOp,
					"rule_operand":        j.rightOp,
				}),
			)
		default:
			switch err.(type) {
			case *ErrInvalidOperand:
				j.setDebugErr(
					newNestedError(err, "operands are not the right value type").Set(ErrVals{
						"attr_path":           ctx.AttrPath().GetText(),
						"object_path_operand": j.leftOp,
						"rule_operand":        j.rightOp,
					}),
				)
			default:
				j.setDebugErr(
					newNestedError(err, "unknown error").Set(ErrVals{
						"attr_path":           ctx.AttrPath().GetText(),
						"object_path_operand": j.leftOp,
						"rule_operand":        j.rightOp,
					}),
				)

			}
		}

		return false
	}
	return ret
}

func GetSubAttr(source interface{}, index string) interface{} {
	if source == nil {
		return nil
	}

	if m, ok := source.(map[string]interface{}); ok {
		return m[index]
	}

	if reflect.TypeOf(source).Kind() == reflect.Struct {
		field := reflect.ValueOf(source).FieldByName(index)
		if !field.IsValid() {
			return nil
		}
		return field.Interface()
	} else if reflect.TypeOf(source).Kind() == reflect.Map {
		field := reflect.ValueOf(source).MapIndex(reflect.ValueOf(index))
		if !field.IsValid() {
			return nil
		}
		return field.Interface()
	}

	return nil
}

// VisitAttrPath resolves the left operand and starts a fresh comparison, so
// nothing leaks in from the previous one.
func (j *JsonQueryVisitorImpl) VisitAttrPath(ctx *AttrPathContext) interface{} {
	var item interface{} = j.item
	for cur := IAttrPathContext(ctx); ; {
		item = GetSubAttr(item, cur.ATTRNAME().GetText())
		sub := cur.SubAttr()
		if sub == nil || sub.AttrPath() == nil {
			break
		}
		cur = sub.AttrPath()
	}
	j.leftOp = item
	j.rightOp = nil
	j.currentOperation = GetCurrentOperationByRight(item)
	return nil
}

func (j *JsonQueryVisitorImpl) VisitSubAttr(ctx *SubAttrContext) interface{} {
	return ctx.AttrPath().Accept(j)
}

func GetCurrentOperationByRight(right interface{}) Operation {
	if right == nil {
		return &NullOperation{}
	} else if _, ok := right.(net.IP); ok {
		return &IPCompareOperation{}
	} else if _, ok := right.(*net.IPNet); ok {
		return &IPCompareOperation{}
	} else if _, ok := right.([]net.IP); ok {
		return &IPCompareOperation{}
	} else if _, ok := right.([]*net.IPNet); ok {
		return &IPCompareOperation{}
	} else if _, ok := right.(*regexp.Regexp); ok {
		return &RegexOperation{}
	}

	reflectType := reflect.TypeOf(right)
	//fmt.Println("reflectType", reflectType.String())

	if reflectType.Kind() == reflect.Slice || reflectType.Kind() == reflect.Array {
		if reflectType.Elem().Kind() == reflect.String {
			return &StringOperation{}
		}
		if reflectType.Elem().Kind() == reflect.Float64 || reflectType.Elem().Kind() == reflect.Float32 {
			return &FloatOperation{}
		}
		if reflectType.Elem().Kind() == reflect.Int || reflectType.Elem().Kind() == reflect.Int64 {
			return &IntOperation{}
		}
		if reflectType.Elem().Kind() == reflect.Bool {
			return &BoolOperation{}
		}

	} else {
		if reflectType.Kind() == reflect.String {
			return &StringOperation{}
		} else if reflectType.Kind() == reflect.Float64 || reflectType.Kind() == reflect.Float32 {
			return &FloatOperation{}
		} else if reflectType.Kind() == reflect.Int || reflectType.Kind() == reflect.Int64 {
			return &IntOperation{}
		} else if reflectType.Kind() == reflect.Bool {
			return &BoolOperation{}
		}
	}

	return nil
}

func (j *JsonQueryVisitorImpl) VisitValueAttrPath(ctx *ValueAttrPathContext) interface{} {
	var item interface{} = j.item
	for cur := IValueAttrPathContext(ctx); ; {
		item = GetSubAttr(item, cur.ATTRNAME().GetText())
		sub := cur.ValueSubAttr()
		if sub == nil || sub.ValueAttrPath() == nil {
			break
		}
		cur = sub.ValueAttrPath()
	}
	j.rightOp = item

	// the left operand picks the operation; fall back to the right one when
	// the left is missing or of an unsupported type
	if j.leftOp == nil || j.currentOperation == nil {
		if op := GetCurrentOperationByRight(item); op != nil {
			j.currentOperation = op
		}
	}
	return nil
}

func (j *JsonQueryVisitorImpl) VisitValueSubAttr(ctx *ValueSubAttrContext) interface{} {
	return ctx.ValueAttrPath().Accept(j)
}

func (j *JsonQueryVisitorImpl) VisitBoolean(ctx *BooleanContext) interface{} {
	j.currentOperation = &BoolOperation{}

	val, err := strconv.ParseBool(ctx.GetText())
	if err != nil {
		j.rightOp = nil
		j.setErr(newNestedError(err, "Error converting boolean"))
		return nil
	}
	j.rightOp = val
	return nil
}

func (j *JsonQueryVisitorImpl) VisitNull(ctx *NullContext) interface{} {
	j.currentOperation = &NullOperation{}
	j.rightOp = nil
	return nil
}

func getString(s string) string {
	if len(s) > 2 {
		return s[1 : len(s)-1]
	}
	return ""
}

func (j *JsonQueryVisitorImpl) VisitString(ctx *StringContext) interface{} {
	j.currentOperation = &StringOperation{}
	j.rightOp = getString(ctx.GetText())
	return nil
}

func (j *JsonQueryVisitorImpl) VisitDouble(ctx *DoubleContext) interface{} {
	j.currentOperation = &FloatOperation{}
	val, err := strconv.ParseFloat(ctx.GetText(), 10)
	if err != nil {
		j.rightOp = nil
		j.setErr(err)
		return nil
	}
	j.rightOp = val
	return nil
}

func (j *JsonQueryVisitorImpl) VisitVersion(ctx *VersionContext) interface{} {
	j.currentOperation = &VersionOperation{}
	j.rightOp = ctx.VERSION().GetText()
	return nil
}

func (j *JsonQueryVisitorImpl) VisitLong(ctx *LongContext) interface{} {
	j.currentOperation = &IntOperation{}
	val, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err != nil {
		j.rightOp = nil
		j.setErr(err)
		return nil
	}
	j.rightOp = int(val)
	return nil
}

func (j *JsonQueryVisitorImpl) VisitListOfInts(ctx *ListOfIntsContext) interface{} {
	j.currentOperation = &IntOperation{}
	return ctx.ListInts().Accept(j)
}

func (j *JsonQueryVisitorImpl) VisitListInts(ctx *ListIntsContext) interface{} {
	return ctx.SubListOfInts().Accept(j)
}

func (j *JsonQueryVisitorImpl) VisitSubListOfInts(ctx *SubListOfIntsContext) interface{} {
	if j.rightOp == nil {
		j.rightOp = make([]int, 0)
	}
	list := j.rightOp.([]int)
	val, err := strconv.ParseInt(ctx.INT().GetText(), 10, 64)
	if err != nil {
		j.setErr(err)
		return nil
	}
	j.rightOp = append(list, int(val))
	if ctx.SubListOfInts() == nil || ctx.SubListOfInts().IsEmpty() {
		return nil
	}
	return ctx.SubListOfInts().Accept(j)
}

func (j *JsonQueryVisitorImpl) VisitListOfDoubles(ctx *ListOfDoublesContext) interface{} {
	j.currentOperation = &FloatOperation{}
	return ctx.ListDoubles().Accept(j)
}

func (j *JsonQueryVisitorImpl) VisitListDoubles(ctx *ListDoublesContext) interface{} {
	return ctx.SubListOfDoubles().Accept(j)
}

func (j *JsonQueryVisitorImpl) VisitSubListOfDoubles(ctx *SubListOfDoublesContext) interface{} {
	if j.rightOp == nil {
		j.rightOp = make([]float64, 0)
	}
	list := j.rightOp.([]float64)
	val, err := strconv.ParseFloat(ctx.DOUBLE().GetText(), 10)
	if err != nil {
		j.setErr(err)
		return nil
	}
	j.rightOp = append(list, val)
	if ctx.SubListOfDoubles() == nil || ctx.SubListOfDoubles().IsEmpty() {
		return nil
	}
	return ctx.SubListOfDoubles().Accept(j)
}

func (j *JsonQueryVisitorImpl) VisitListOfStrings(ctx *ListOfStringsContext) interface{} {
	j.currentOperation = &StringOperation{}
	return ctx.ListStrings().Accept(j)
}

func (j *JsonQueryVisitorImpl) VisitListStrings(ctx *ListStringsContext) interface{} {
	return ctx.SubListOfStrings().Accept(j)
}

func (j *JsonQueryVisitorImpl) VisitSubListOfStrings(ctx *SubListOfStringsContext) interface{} {
	if j.rightOp == nil {
		j.rightOp = make([]string, 0)
	}
	val := getString(ctx.STRING().GetText())
	list := j.rightOp.([]string)
	j.rightOp = append(list, val)
	if ctx.SubListOfStrings() == nil || ctx.SubListOfStrings().IsEmpty() {
		return nil
	}
	return ctx.SubListOfStrings().Accept(j)
}

/* NEEEEEEEEEEEEEEEEEEEEEW STUFFFFFFFFFFFFFFFFFFFF */

// parseIPNet reads an address or CIDR literal; a bare address becomes a
// single-host network.
func parseIPNet(txt string) (*net.IPNet, error) {
	if strings.Contains(txt, "/") {
		_, ipNet, err := net.ParseCIDR(txt)
		return ipNet, err
	}
	ip := net.ParseIP(txt)
	if ip == nil {
		return nil, errors.New("Invalid IP value")
	}
	if v4 := ip.To4(); v4 != nil {
		ip = v4
	}
	return &net.IPNet{IP: ip, Mask: net.CIDRMask(len(ip)*8, len(ip)*8)}, nil
}

func (j *JsonQueryVisitorImpl) VisitListIPs(ctx *ListIPsContext) interface{} {
	j.currentOperation = &IPCompareOperation{}
	nets := []*net.IPNet{}
	for sub := ctx.SubListOfIPs(); sub != nil && sub.IpValue() != nil; sub = sub.SubListOfIPs() {
		ipNet, err := parseIPNet(sub.IpValue().GetText())
		if err != nil {
			j.setErr(err)
			return nil
		}
		nets = append(nets, ipNet)
	}
	j.rightOp = nets
	return nil
}

// VisitSubListOfIPs is only here to satisfy JsonQueryVisitor; VisitListIPs walks the list.
func (j *JsonQueryVisitorImpl) VisitSubListOfIPs(ctx *SubListOfIPsContext) interface{} {
	return nil
}

func (j *JsonQueryVisitorImpl) VisitIpValue(ctx *IpValueContext) interface{} {
	j.currentOperation = &IPCompareOperation{}

	isCIDR := false
	txt := ""
	var err error

	if ctx.IP_ADDRESS() != nil {
		txt = ctx.IP_ADDRESS().GetText()
	} else if ctx.IP_CIDR() != nil {
		txt = ctx.IP_CIDR().GetText()
		isCIDR = true
	} else {
		j.setErr(errors.New("Invalid IP value"))
		return nil
	}

	if j.rightOp == "" {
		j.setErr(errors.New("Invalid IP value"))
		return nil
	}

	if isCIDR {
		var val *net.IPNet
		_, val, err = net.ParseCIDR(txt)
		if err != nil {
			j.setErr(err)
			return nil
		}

		j.rightOp = val
	} else {
		ip := net.ParseIP(txt)
		if ip == nil {
			// the lexer accepts more groups than a real address can have
			j.setErr(errors.New("Invalid IP value"))
			return nil
		}
		j.rightOp = ip
	}

	return nil
}

// compileRegexLiteral compiles a /pattern/flags literal from the rule text.
func compileRegexLiteral(rawRegex string) (*regexp.Regexp, error) {
	lastIndex := strings.LastIndex(rawRegex, "/")
	if !strings.HasPrefix(rawRegex, "/") || lastIndex < 1 {
		return nil, errors.New("Invalid regex")
	}

	regex := rawRegex[1:lastIndex]
	for _, flag := range rawRegex[lastIndex+1:] {
		switch flag {
		case 'i':
			regex = "(?i)" + regex
		case 'm':
			regex = "(?m)" + regex
		}
	}
	return regexp.Compile(regex)
}

// collectRegexes precompiles every regex literal in the tree. Literals that
// fail to compile are left out, so Process reports the error as before.
func collectRegexes(tree antlr.Tree, cache map[*RegexValueContext]*regexp.Regexp) {
	if ctx, ok := tree.(*RegexValueContext); ok && ctx.REGEX() != nil {
		if re, err := compileRegexLiteral(ctx.REGEX().GetText()); err == nil {
			cache[ctx] = re
		}
		return
	}
	for _, child := range tree.GetChildren() {
		collectRegexes(child, cache)
	}
}

func (j *JsonQueryVisitorImpl) VisitRegexValue(ctx *RegexValueContext) interface{} {
	if ctx.REGEX() == nil {
		// `a mt b`: the pattern comes from the input object
		ctx.ValueAttrPath().Accept(j)
		j.currentOperation = &RegexOperation{}
		return nil
	}
	j.currentOperation = &RegexOperation{}
	if re, ok := j.regexCache[ctx]; ok {
		j.rightOp = re
		return nil
	}

	compiledRegex, err := compileRegexLiteral(ctx.REGEX().GetText())
	if err != nil {
		j.setErr(err)
		return nil
	}

	j.rightOp = compiledRegex
	return nil
}

func (j *JsonQueryVisitorImpl) VisitRegexExp(ctx *RegexExpContext) interface{} {
	ctx.AttrPath().Accept(j)
	ctx.RegexValue().Accept(j)
	if j.hasErr() {
		return false
	}
	var apply func(Operand, Operand) (bool, error)
	currentOp := j.currentOperation
	if currentOp == nil {
		j.setErr(ErrInvalidOperation)
		return false
	}
	switch ctx.op.GetTokenType() {
	case JsonQueryParserMT:
		apply = currentOp.MT
	default:
		j.setErr(fmt.Errorf("Unknown operation %s", ctx.op.GetText()))
		return false
	}
	defer func() { j.rightOp = nil }()
	ret, err := apply(j.leftOp, j.rightOp)
	if err != nil {
		switch err {
		case ErrInvalidOperation:
			// in case of invalid operation lets rather
			// be conservative and return false because the rule doesn't even make
			// sense. It can be argued that it would be false positive if we were
			// to return true
			j.setErr(err)
			j.setDebugErr(
				newNestedError(err, "Not a valid operation for datatypes").Set(ErrVals{
					"operation":           ctx.op.GetTokenType(),
					"object_path_operand": j.leftOp,
					"rule_operand":        j.rightOp,
				}),
			)
		case ErrEvalOperandMissing:
			j.setDebugErr(
				newNestedError(err, "Eval operand missing in input object").Set(ErrVals{
					"attr_path": ctx.AttrPath().GetText(),
				}),
			)
		default:
			switch err.(type) {
			case *ErrInvalidOperand:
				j.setDebugErr(
					newNestedError(err, "operands are not the right value type").Set(ErrVals{
						"attr_path":           ctx.AttrPath().GetText(),
						"object_path_operand": j.leftOp,
						"rule_operand":        j.rightOp,
					}),
				)
			default:
				j.setDebugErr(
					newNestedError(err, "unknown error").Set(ErrVals{
						"attr_path":           ctx.AttrPath().GetText(),
						"object_path_operand": j.leftOp,
						"rule_operand":        j.rightOp,
					}),
				)

			}
		}

		return false
	}
	return ret
}

func (j *JsonQueryVisitorImpl) VisitIpCompareExp(ctx *IpCompareExpContext) interface{} {
	ctx.AttrPath().Accept(j)
	if list := ctx.ListIPs(); list != nil {
		list.Accept(j)
	} else {
		ctx.IpValue().Accept(j)
	}
	if j.hasErr() {
		return false
	}
	var apply func(Operand, Operand) (bool, error)
	currentOp := j.currentOperation
	if currentOp == nil {
		j.setErr(ErrInvalidOperation)
		return false
	}
	switch ctx.op.GetTokenType() {
	case JsonQueryParserEQ:
		apply = currentOp.EQ
	case JsonQueryParserNE:
		apply = currentOp.NE
	case JsonQueryParserIN:
		apply = currentOp.IN
	default:
		j.setErr(fmt.Errorf("Unknown operation %s", ctx.op.GetText()))
		return false
	}
	defer func() { j.rightOp = nil }()
	ret, err := apply(j.leftOp, j.rightOp)
	if err != nil {
		switch err {
		case ErrInvalidOperation:
			// in case of invalid operation lets rather
			// be conservative and return false because the rule doesn't even make
			// sense. It can be argued that it would be false positive if we were
			// to return true
			j.setErr(err)
			j.setDebugErr(
				newNestedError(err, "Not a valid operation for datatypes").Set(ErrVals{
					"operation":           ctx.op.GetTokenType(),
					"object_path_operand": j.leftOp,
					"rule_operand":        j.rightOp,
				}),
			)
		case ErrEvalOperandMissing:
			j.setDebugErr(
				newNestedError(err, "Eval operand missing in input object").Set(ErrVals{
					"attr_path": ctx.AttrPath().GetText(),
				}),
			)
		default:
			switch err.(type) {
			case *ErrInvalidOperand:
				j.setDebugErr(
					newNestedError(err, "operands are not the right value type").Set(ErrVals{
						"attr_path":           ctx.AttrPath().GetText(),
						"object_path_operand": j.leftOp,
						"rule_operand":        j.rightOp,
					}),
				)
			default:
				j.setDebugErr(
					newNestedError(err, "unknown error").Set(ErrVals{
						"attr_path":           ctx.AttrPath().GetText(),
						"object_path_operand": j.leftOp,
						"rule_operand":        j.rightOp,
					}),
				)

			}
		}

		return false
	}
	return ret
}

func (j *JsonQueryVisitorImpl) VisitVariable(ctx *VariableContext) interface{} {
	return ctx.ValueAttrPath().Accept(j)
}
