package calculator

import (
	"errors"
	"testing"
)

func TestCalc(t *testing.T) {
	var tests = []struct {
		exp     Expression
		want    float64
		wantErr error
	}{
		{Expression{op1: 3, op2: 2, operator: "+"}, 5, nil},
		{Expression{op1: -9, op2: -2, operator: "+"}, -11, nil},
		{Expression{op1: -2, op2: 6, operator: "*"}, -12, nil},
		{Expression{op1: 1.23, op2: 6.5, operator: "*"}, 7.995, nil},
		{Expression{op1: 5, op2: 5, operator: "+"}, 10, nil},
		{Expression{op1: -8, op2: 3, operator: "+"}, -5, nil},
		{Expression{op1: 4, op2: 9, operator: "*"}, 36, nil},
		{Expression{op1: -2.5, op2: 5, operator: "*"}, -12.5, nil},
		{Expression{op1: 8, op2: 0, operator: "/"}, 0, ErrDivByZero},
		{Expression{op1: 10, op2: 2, operator: "/"}, 5, nil},
		{Expression{op1: 15, op2: 3, operator: "/"}, 5, nil},
		{Expression{op1: 7.5, op2: 1.5, operator: "/"}, 5, nil},
		{Expression{op1: 666, op2: 13, operator: "+-"}, 0, ErrUnknownOperator},
	}

	for _, tt := range tests {
		testName := tt.exp.String()
		t.Run(testName, func(t *testing.T) {
			got, err := Calc(tt.exp)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got error %v, want error %v", err, tt.wantErr)
			}
		})
	}
}
