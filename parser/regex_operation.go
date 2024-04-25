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
		return "", newErrInvalidOperand(operand, opVal)
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

	if right == nil {
		return leftVal, nil, nil
	}

	rightVal, ok := right.(*regexp.Regexp)
	if !ok {
		if rightStr, ok := right.(string); !ok {
			return "", nil, ErrEvalOperandMissing
		} else {
			rightVal = regexp.MustCompile(rightStr)
		}
	}

	if rightVal == nil {
		return leftVal, nil, newErrInvalidOperand(right, rightVal)
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
