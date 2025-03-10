package calculator

import "fmt"

// Expression defines a mathematical expression consisting of two operands and an operator.
type Expression struct {
	op1      float64
	op2      float64
	operator string
}

func (e *Expression) String() string {
	return fmt.Sprintf("%v%s%v", e.op1, e.operator, e.op2)
}

func NewExpression(op1, op2 float64, operator string) *Expression {
	return &Expression{op1: op1, op2: op2, operator: operator}
}
