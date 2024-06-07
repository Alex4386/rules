package rules

import "github.com/alex4386/rules/parser"

func Evaluate(rule string, items map[string]interface{}) (bool, error) {
	ev, err := parser.NewEvaluator(rule)
	if err != nil {
		return false, err
	}
	return ev.Process(items)
}
