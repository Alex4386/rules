package parser

import (
	"fmt"
	regexp "github.com/wasilibs/go-re2"
)

type RegexOperation struct {
	NullOperation
}

func (o *RegexOperation) getString(operand Operand) (string, error) {
	switch opVal := operand.(type) {
	case string:
		return opVal, nil
	case fmt.Stringer:
		return opVal.String(), nil
	default:
		return "", newErrInvalidOperand(operand, "")
	}
}

func (o *RegexOperation) get(left Operand, right Operand) (string, *regexp.Regexp, error) {
	if left == nil {
		return "", nil, ErrEvalOperandMissing
	}
	leftVal, err := o.getString(left)
	if err != nil {
		return "", nil, err
	}

	rightVal, ok := right.(*regexp.Regexp)
	if !ok {
		// pattern taken from the input object (`a mt b`)
		rightStr, ok := right.(string)
		if !ok {
			return "", nil, ErrEvalOperandMissing
		}
		if rightVal, err = regexp.Compile(rightStr); err != nil {
			return "", nil, err
		}
	}
	if rightVal == nil {
		return "", nil, ErrEvalOperandMissing
	}

	return leftVal, rightVal, nil
}

func (o *RegexOperation) MT(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}

	return r.MatchString(l), nil
}
