package calculator

import (
	"fmt"
	"github.com/Knetic/govaluate"
)

func Evaluate(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return "", err
	}

	result, err := expr.Evaluate(nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", result), nil
}
