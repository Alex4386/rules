package parser

import (
	"net"
)

type IPCompareOperation struct {
	NullOperation
}

func tryConversion(op Operand) any {
	if op == nil {
		return nil
	}

	if val, ok := op.(string); ok {
		return val
	} else if val, ok := op.(net.IP); ok {
		return val
	} else if val, ok := op.(*net.IPNet); ok {
		return val
	}

	return nil
}

func (o *IPCompareOperation) get(left Operand, right Operand) (net.IP, any, error) {
	if left == nil {
		return nil, nil, ErrEvalOperandMissing
	}
	leftVal, ok := left.(net.IP)
	if !ok {
		var leftStr string
		if leftStr, ok = left.(string); !ok {
			return nil, nil, ErrEvalOperandMissing
		}

		leftVal = net.ParseIP(leftStr)
		if leftVal == nil {
			return nil, nil, newErrInvalidOperand(left, leftVal)
		}
	}

	if right == nil {
		return leftVal, nil, nil
	}

	rightVal, ok := right.(net.IP)
	if !ok {
		if rightNet, ok := right.(*net.IPNet); ok {
			return leftVal, rightNet, nil
		} else {
			var rightStr string
			if rightStr, ok = right.(string); !ok {
				return nil, nil, ErrEvalOperandMissing
			}

			rightVal = net.ParseIP(rightStr)
			if rightVal == nil {
				return nil, nil, newErrInvalidOperand(right, rightVal)
			}
		}
	}

	return leftVal, rightVal, nil
}

func (o *IPCompareOperation) EQ(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}

	if rIP, ok := r.(net.IP); ok {
		return l.Equal(rIP), nil
	} else if rNet, ok := r.(*net.IPNet); ok {
		return rNet.Contains(l), nil
	}

	return false, nil
}

func (o *IPCompareOperation) NE(left Operand, right Operand) (bool, error) {
	value, err := o.EQ(left, right)
	if err != nil {
		return false, err
	}

	return !value, nil
}

func (o *IPCompareOperation) IN(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}

	if rCIDR, ok := r.(*net.IPNet); ok {
		return rCIDR.Contains(l), nil
	} else if rArr, ok := r.([]any); ok {
		for _, data := range rArr {
			if rIP, ok := data.(net.IP); ok {
				if l.Equal(rIP) {
					return true, nil
				}
			} else if rNet, ok := data.(*net.IPNet); ok {
				if rNet.Contains(l) {
					return true, nil
				}
			}
		}
	}

	return false, nil
}
