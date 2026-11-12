package parser

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyntaxErrors(t *testing.T) {
	for _, rule := range []string{
		`(x eq 2)(x eq 1)`,
		`x eq 2 garbage !!`,
		`x eq`,
		`x eq 1 or`,
		`x === 1`,
		``,
	} {
		_, err := NewEvaluator(rule)
		assert.Error(t, err, rule)
	}
	for _, rule := range []string{"x eq 1", "x eq 1\n", "x eq 1 ", "(x eq 1) and name mt /a/i"} {
		_, err := NewEvaluator(rule)
		assert.NoError(t, err, rule)
	}
}

// run with -race
func TestConcurrentProcess(t *testing.T) {
	ev, err := NewEvaluator(`name mt /^hel+o/i and x.y eq 1 and z gt "a"`)
	assert.NoError(t, err)
	var wg sync.WaitGroup
	for g := 0; g < 8; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			in := map[string]interface{}{"name": "Hello", "x": map[string]interface{}{"y": g % 2}}
			for i := 0; i < 200; i++ {
				got, err := ev.Process(in)
				ev.LastDebugErr()
				if err != nil || got {
					t.Errorf("got %v, %v", got, err)
					return
				}
			}
		}(g)
	}
	wg.Wait()
}
