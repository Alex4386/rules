package parser

import "testing"

func TestRegexCache(t *testing.T) {
	in := map[string]interface{}{"name": "Hello Example.COM"}
	for _, opts := range []EvaluatorOptions{{}, {DisableASTRegexCache: true}} {
		for rule, want := range map[string]bool{
			`name mt /exa.*com/i`:                     true,
			`name mt /exa.*com/`:                      false,
			`name mt /^hello/i and name mt /nomatch/`: false,
		} {
			ev, err := NewEvaluatorWithOptions(rule, opts)
			if err != nil {
				t.Fatal(err)
			}
			if cached := len(ev.regexCache) > 0; cached == opts.DisableASTRegexCache {
				t.Fatalf("%s: cached=%v with opts=%+v", rule, cached, opts)
			}
			for i := 0; i < 2; i++ {
				if got, err := ev.Process(in); err != nil || got != want {
					t.Fatalf("%s: got %v, %v; want %v", rule, got, err, want)
				}
			}
		}
		// an invalid pattern still surfaces at Process, cached or not
		ev, _ := NewEvaluatorWithOptions(`name mt /(unclosed/`, opts)
		if _, err := ev.Process(in); err == nil {
			t.Fatal("expected compile error")
		}
	}
}

func benchRegex(b *testing.B, opts EvaluatorOptions) {
	ev, _ := NewEvaluatorWithOptions(`name mt /exa.*com/i`, opts)
	in := map[string]interface{}{"name": "hello world example.com"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ev.Process(in)
	}
}

func BenchmarkRegexCached(b *testing.B) { benchRegex(b, EvaluatorOptions{}) }
func BenchmarkRegexUncached(b *testing.B) {
	benchRegex(b, EvaluatorOptions{DisableASTRegexCache: true})
}
