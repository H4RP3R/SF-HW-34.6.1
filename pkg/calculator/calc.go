package calculator

import (
	"fmt"
)

var ErrDivByZero = fmt.Errorf("division by zero")
var ErrUnknownOperator = fmt.Errorf("unknown operator")

// Calc evaluates a mathematical expression represented by an Expression struct and returns the result as a float64.
// It supports basic mathematical operations: addition (+), subtraction (-), multiplication (*), and division (/).
// The function returns an error for unsupported or unknown operators.
//
// Parameters:
// - exp: Expression struct containing the operands (op1 and op2) and the operator.
//
// Returns:
// - float64: The result of evaluating the expression.
// - error: An error if the expression involves division by zero or an unknown operator.
func Calc(exp Expression) (float64, error) {
	switch exp.operator {
	case "+":
		return exp.op1 + exp.op2, nil
	case "-":
		return exp.op1 - exp.op2, nil
	case "*":
		return exp.op1 * exp.op2, nil
	case "/":
		if exp.op2 == 0 {
			return 0, ErrDivByZero
		}
		return exp.op1 / exp.op2, nil
	}

	return 0, ErrUnknownOperator
}
