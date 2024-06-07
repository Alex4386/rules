package parser

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func CIDRBuilder(cidr ...string) []*net.IPNet {
	var result []*net.IPNet
	for _, c := range cidr {
		_, ipNet, err := net.ParseCIDR(c)
		if err != nil {
			panic(err)
		}
		result = append(result, ipNet)
	}
	return result
}

func TestVariableMatchedRule(t *testing.T) {
	tests := []testCase{
		{
			"x eq y",
			obj{
				"x": 1,
				"y": 1,
			},
			true,
			false,
		},
		{
			"x eq z",
			obj{
				"x": net.ParseIP("1.1.1.1"),
				"z": net.ParseIP("1.1.1.1"),
			},
			true,
			false,
		},
		{
			"x in arr",
			obj{
				"x": net.ParseIP("1.1.1.1"),
				"arr": []net.IP{
					net.ParseIP("1.1.1.1"),
					net.ParseIP("1.1.1.2"),
				},
			},
			true,
			false,
		},
		{
			"x in arr",
			obj{
				"x":   net.ParseIP("1.1.1.1"),
				"arr": CIDRBuilder("1.1.1.0/24", "1.1.0.0/24"),
			},
			true,
			false,
		},
	}

	for _, tt := range tests {
		result, err := evalPanic(t, tt.rule, tt.input)
		if tt.hasError {
			assert.Error(t, err)
			continue
		} else {
			assert.NoError(t, err)
			assert.Equal(t, result, tt.result, fmt.Sprintf("rule :%s, input :%v", tt.rule, tt.input))
		}
		assert.Equal(t, false, Evaluate(tt.rule, obj{"x": 4.5}), tt.rule)
		assert.Equal(t, tt.result, Evaluate(fmt.Sprintf("(%s)", tt.rule), tt.input), tt.rule)
	}
}

func TestIPMatchedRule(t *testing.T) {
	x, err := NewEvaluatorWithPanicOnParseError("x IN 1.0.0.0/8")
	if err != nil {
		t.Fatal(err)
	}

	vaaa, err := x.Process(obj{"x": net.ParseIP("1.1.1.1")})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("returned value: ", vaaa)

	tests := []testCase{
		{
			`x IN 1.0.0.0/8`,
			obj{
				"x": net.ParseIP("1.1.1.1"),
			},
			true,
			false,
		},
		{
			`x in 1.0.0.1/32`,
			obj{
				"x": net.ParseIP("1.1.1.1"),
			},
			false,
			false,
		},
		{
			`x eq 1.1.1.1`,
			obj{
				"x": net.ParseIP("1.1.1.1"),
			},
			true,
			false,
		},
		{
			`src.ip eq 1.1.1.1`,
			obj{
				"src": obj{
					"ip": net.ParseIP("1.1.1.1"),
				},
			},
			true,
			false,
		},
		{
			`x in 1.0.0.0/8`,
			obj{
				"x": net.ParseIP("1.1.1.1"),
			},
			true,
			false,
		},
		{
			`x in 1.0.0.1/32`,
			obj{
				"x": net.ParseIP("1.1.1.1"),
			},
			false,
			false,
		},
		{
			`x in 2001::8a2e:1/8`,
			obj{
				"x": net.ParseIP("2001:0db8:85a3:0000:0000:8a2e:0370:7334"),
			},
			false,
			false,
		},
	}

	for _, tt := range tests {
		result, err := eval(t, tt.rule, tt.input)
		if tt.hasError {
			assert.Error(t, err)
			continue
		} else {
			assert.NoError(t, err)
			assert.Equal(t, result, tt.result, fmt.Sprintf("rule :%s, input :%v", tt.rule, tt.input))
		}
		assert.Equal(t, false, Evaluate(tt.rule, obj{"x": 4.5}), tt.rule)
		assert.Equal(t, tt.result, Evaluate(fmt.Sprintf("(%s)", tt.rule), tt.input), tt.rule)
	}
}
