package parser

import (
	"errors"
	"fmt"
)

// numIn reports whether left is in a rule list of ints or doubles.
// shortcut: compares as float64, so ints beyond 2^53 lose precision; compare as int64 if that matters.
func numIn(left Operand, right Operand) (bool, error) {
	if left == nil {
		return false, ErrEvalOperandMissing
	}
	l, err := toFloat(left)
	if err != nil {
		return false, err
	}
	switch list := right.(type) {
	case []int:
		for _, n := range list {
			if float64(n) == l {
				return true, nil
			}
		}
	case []float64:
		for _, n := range list {
			if n == l {
				return true, nil
			}
		}
	default:
		return false, newErrInvalidOperand(right, []float64{})
	}
	return false, nil
}

func toFloat(op Operand) (float64, error) {
	switch val := op.(type) {
	case int:
		return float64(val), nil
	case float64:
		return val, nil
	case float32:
		return float64(val), nil
	case int32:
		return float64(val), nil
	case int64:
		return float64(val), nil
	}
	var exp float64
	return 0, newErrInvalidOperand(op, exp)
}

func toInt(op Operand) (int, error) {
	switch val := op.(type) {
	case int:
		return val, nil
	case float64:
		return int(val), nil
	case int32:
		return int(val), nil
	case int64:
		return int(val), nil
	}
	var exp int
	return 0, newErrInvalidOperand(op, exp)
}

type Operand interface{}

var (
	ErrInvalidOperation   = errors.New("Invalid operation on the type")
	ErrEvalOperandMissing = errors.New("Operand not present")
)

type ErrInvalidOperand struct {
	Val     interface{}
	typeObj interface{}
}

func newErrInvalidOperand(val Operand, typeObj interface{}) *ErrInvalidOperand {
	return &ErrInvalidOperand{
		Val:     val,
		typeObj: typeObj,
	}
}

func (e *ErrInvalidOperand) Error() string {
	// %T rather than reflect.TypeOf(x).String(): either side may be nil
	return fmt.Sprintf("Operand %v is not the correct type. Expected: %T, Actual: %T",
		e.Val, e.typeObj, e.Val)
}

type Operation interface {
	EQ(left Operand, right Operand) (bool, error)
	NE(left Operand, right Operand) (bool, error)
	GT(left Operand, right Operand) (bool, error)
	LT(left Operand, right Operand) (bool, error)
	GE(left Operand, right Operand) (bool, error)
	LE(left Operand, right Operand) (bool, error)
	CO(left Operand, right Operand) (bool, error)
	SW(left Operand, right Operand) (bool, error)
	EW(left Operand, right Operand) (bool, error)
	IN(left Operand, right Operand) (bool, error)
	MT(left Operand, right Operand) (bool, error)
}
