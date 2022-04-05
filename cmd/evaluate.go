package main

import "errors"

type evaluator struct{}

func Evaluate(expr Expr) (float64, error) {
	result, err := expr.Accept(evaluator{})
	return result.(float64), err
}

func (e evaluator) VisitBinary(expr Binary) (any, error) {
	left, err := expr.Left.Accept(e)
	if err != nil {
		return float64(0), err
	}

	right, err := expr.Right.Accept(e)
	if err != nil {
		return float64(0), err
	}

	switch expr.Operator.Type {
	case PLUS:
		return left.(float64) + right.(float64), nil
	case MINUS:
		return left.(float64) - right.(float64), nil
	case STAR:
		return left.(float64) * right.(float64), nil
	case SLASH:
		return left.(float64) / right.(float64), nil
	}

	return float64(0), errors.New("invalid-binary-operator")
}

func (e evaluator) VisitGrouping(expr Grouping) (any, error) {
	return expr.Expr.Accept(e)
}

func (e evaluator) VisitLiteral(expr Literal) (any, error) {
	return expr.Value, nil
}

func (e evaluator) VisitUnary(expr Unary) (any, error) {
	if expr.Operator.Type == MINUS {
		value, err := expr.Right.Accept(e)
		if err != nil {
			return float64(0), err
		}
		return -value.(float64), nil
	}

	return float64(0), errors.New("invalid-unary-operator")
}
