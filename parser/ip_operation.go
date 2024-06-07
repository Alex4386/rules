package parser

import (
	"net"
)

type IPCompareOperation struct {
	NullOperation
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

	var rightVal any
	if v, ok := right.(string); ok {
		_, ipNet, err := net.ParseCIDR(v)
		if err != nil {
			rightVal = net.ParseIP(v)
		} else {
			rightVal = ipNet
		}
	} else if v, ok := right.(net.IP); ok {
		rightVal = v
	} else if v, ok := right.(*net.IPNet); ok {
		rightVal = v
	} else if v, ok := right.([]net.IP); ok {
		rightVal = v
	} else if v, ok := right.([]*net.IPNet); ok {
		rightVal = v
	} else {
		return nil, nil, newErrInvalidOperand(right, rightVal)
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

	if v, ok := r.(net.IP); ok {
		return l.Equal(v), nil
	} else if v, ok := r.(*net.IPNet); ok {
		return v.Contains(l), nil
	} else if v, ok := r.([]net.IP); ok {
		for _, data := range v {
			if l.Equal(data) {
				return true, nil
			}
		}
	} else if v, ok := r.([]*net.IPNet); ok {
		for _, data := range v {
			if data.Contains(l) {
				return true, nil
			}
		}
	}

	return false, nil
}
