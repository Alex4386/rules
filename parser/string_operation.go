package parser

import (
	"fmt"
	"net"
	"strings"
)

type StringOperation struct {
	NullOperation
}

func (o *StringOperation) getString(operand Operand) (string, error) {
	switch opVal := operand.(type) {
	case string:
		return opVal, nil
	case fmt.Stringer:
		return opVal.String(), nil
	default:
		return "", newErrInvalidOperand(operand, opVal)
	}
}

func (o *StringOperation) get(left Operand, right Operand) (string, string, error) {
	if left == nil {
		return "", "", ErrEvalOperandMissing
	}
	leftVal, err := o.getString(left)
	if err != nil {
		return "", "", err
	}
	rightVal, err := o.getString(right)
	if err != nil {
		return "", "", err
	}
	return strings.ToLower(leftVal), strings.ToLower(rightVal), nil

}

func (o *StringOperation) EQ(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l == r, nil
}

func (o *StringOperation) NE(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l != r, nil
}

func (o *StringOperation) GT(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l > r, nil
}

func (o *StringOperation) LT(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l < r, nil
}

func (o *StringOperation) GE(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l >= r, nil
}

func (o *StringOperation) LE(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l <= r, nil
}

func (o *StringOperation) CO(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}

	return strings.Contains(l, r), nil
}

func (o *StringOperation) SW(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return strings.HasPrefix(l, r), nil
}

func (o *StringOperation) EW(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return strings.HasSuffix(l, r), nil
}

func (o *StringOperation) IN(left Operand, right Operand) (bool, error) {
	leftVal, err := o.getString(left)
	if err != nil {
		return false, err
	}

	rightVal, ok := right.([]string)
	if !ok {
		if rightValT, ok := right.(string); !ok {
			return ok, newErrInvalidOperand(right, rightVal)
		} else {
			rightVal = []string{rightValT}
		}
	}

	for _, val := range rightVal {
		// if val is cidr, then check if leftVal is in the cidr
		if strings.Contains(val, "/") {
			_, cidr, err := net.ParseCIDR(val)
			if err == nil {
				ip := net.ParseIP(leftVal)
				if cidr.Contains(ip) {
					return true, nil
				}
			}
		}

		if leftVal == val {
			return true, nil
		}
	}
	return false, nil
}
