package parser

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluationFixes(t *testing.T) {
	in := obj{
		"name": "Alice", "other": "^ali", "pat": "^Ali", "bad": "(", "x": 1.0, "y": 1.0, "f": 1.9, "n": 3, "big": int64(7),
		"ip": "10.1.2.3", "ip6": net.ParseIP("2001:db8::1"), "mapped": "::ffff:1.2.3.4",
		"a": obj{"b": "deep"},
	}
	tests := []struct {
		rule    string
		want    bool
		wantErr bool
	}{
		// the operation type of one comparison must not leak into the next
		{`name == "Alice" and x == y`, true, false},
		// string comparisons are case-sensitive
		{`name == "alice"`, false, false},
		{`name != "alice"`, true, false},
		{`name co "LIC"`, false, false},
		{`name sw "Ali"`, true, false},
		{`x == 1 and name == other`, false, false},
		{`ip in 10.0.0.0/8 and n > 2`, true, false},
		// a missing nested path must not disturb the comparisons after it
		{`nope.deep.er == 1 or n == 3`, true, false},
		{`n == 3 and nope.deep pr`, false, false},
		{`nope.deep == "x" or a.b == "deep"`, true, false},
		// missing left operand compared with a variable
		{`nope == name`, false, false},
		{`nope == alsonope`, true, false},
		// regex pattern taken from the input
		{`name mt other`, false, false},
		{`name mt pat`, true, false},
		{`name mt bad`, false, false},
		{`name mt nope`, false, false},
		// JSON numbers are float64
		{`x in [1, 2]`, true, false},
		{`f in [1, 2]`, false, false},
		{`n in [1.5, 3.0]`, true, false},
		{`big in [7, 8]`, true, false},
		{`nope in [1, 2]`, false, false},
		{`f == 1`, false, false},
		{`f > 1`, true, false},
		{`f < 2`, true, false},
		{`x == 1`, true, false},
		// string IN is case-sensitive
		{`name in ["bob", "ALICE"]`, false, false},
		{`name in ["bob", "Alice"]`, true, false},
		{`nope in ["bob"]`, false, false},
		// IP lists
		{`ip in [192.168.0.0/16, 10.0.0.0/8]`, true, false},
		{`ip in [192.168.0.0/16, 10.1.2.3]`, true, false},
		{`ip in [192.168.0.0/16, 10.1.2.4]`, false, false},
		{`ip6 in [10.0.0.0/8, 2001:db8::/32]`, true, false},
		{`ip6 in [2001:db8::1]`, true, false},
		{`mapped eq ::ffff:1.2.3.4`, true, false},
		{`ip in [1:2:3:4:5:6:7::8:9]`, false, true},
	}
	for _, tt := range tests {
		ev, err := NewEvaluator(tt.rule)
		if !assert.NoError(t, err, tt.rule) {
			continue
		}
		got, err := ev.Process(in)
		assert.Equal(t, tt.want, got, tt.rule)
		assert.Equal(t, tt.wantErr, err != nil, "%s: err=%v", tt.rule, err)
		if dbg := ev.LastDebugErr(); dbg != nil {
			_ = dbg.Error() // must not panic on nil operands
		}
	}

	// a list is only valid with `in`
	_, err := NewEvaluator(`ip eq [10.0.0.0/8]`)
	assert.Error(t, err)
}
