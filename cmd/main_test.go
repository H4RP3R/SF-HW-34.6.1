package main

import (
	"errors"
	"reflect"
	"regexp"
	"testing"

	calc "SF-HW-34.6.1/pkg/calculator"
)

func Test_parseExpression(t *testing.T) {
	type args struct {
		re  *regexp.Regexp
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    *calc.Expression
		wantErr error
	}{
		{
			name: "valid expression",
			args: args{
				re:  expRe,
				str: "2+3=?",
			},
			want:    calc.NewExpression(2, 3, "+"),
			wantErr: nil,
		},
		{
			name: "valid expression with spaces between",
			args: args{
				re:  expRe,
				str: "2 + 3 = ?",
			},
			want:    calc.NewExpression(2, 3, "+"),
			wantErr: nil,
		},
		{
			name: "valid expression with leading and trailing spaces",
			args: args{
				re:  expRe,
				str: "   2+3=?  ",
			},
			want:    calc.NewExpression(2, 3, "+"),
			wantErr: nil,
		},
		{
			name: "valid expression with tabs",
			args: args{
				re:  expRe,
				str: "	2	+	3	=	?		",
			},
			want:    calc.NewExpression(2, 3, "+"),
			wantErr: nil,
		},
		{
			name: "invalid expression without '=?'",
			args: args{
				re:  expRe,
				str: "2+3",
			},
			want:    nil,
			wantErr: ErrNoExpressionFound,
		},
		{
			name: "invalid expression with invalid math operator",
			args: args{
				re:  expRe,
				str: "2$3",
			},
			want:    nil,
			wantErr: ErrNoExpressionFound,
		},
		{
			name: "invalid expression with doubled valid math operator",
			args: args{
				re:  expRe,
				str: "2--3",
			},
			want:    nil,
			wantErr: ErrNoExpressionFound,
		},
		{
			name: "invalid expression without one operand",
			args: args{
				re:  expRe,
				str: "2+ =?",
			},
			want:    nil,
			wantErr: ErrNoExpressionFound,
		},
		{
			name: "invalid expression with letters instead of numbers",
			args: args{
				re:  expRe,
				str: "a+b=?",
			},
			want:    nil,
			wantErr: ErrNoExpressionFound,
		},
		{
			name: "valid expression with big operands",
			args: args{
				re:  expRe,
				str: "87123312+1237812563=?",
			},
			want:    calc.NewExpression(87123312, 1237812563, "+"),
			wantErr: nil,
		},
		{
			name: "valid expression with one negative operand",
			args: args{
				re:  expRe,
				str: "113+-67=?",
			},
			want:    calc.NewExpression(113, -67, "+"),
			wantErr: nil,
		},
		{
			name: "valid expression with both negative operands",
			args: args{
				re:  expRe,
				str: "-113+-67=?",
			},
			want:    calc.NewExpression(-113, -67, "+"),
			wantErr: nil,
		},
		{
			name: "valid expression with subtraction and negative operands",
			args: args{
				re:  expRe,
				str: "-3 - -5 =?",
			},
			want:    calc.NewExpression(-3, -5, "-"),
			wantErr: nil,
		},
		{
			name: "valid expression with floating point numbers",
			args: args{
				re:  expRe,
				str: "3.5+2.1=?",
			},
			want:    calc.NewExpression(3.5, 2.1, "+"),
			wantErr: nil,
		},
		{
			name: "invalid expression with more than two operands",
			args: args{
				re:  expRe,
				str: "2+3+4=?",
			},
			want:    nil,
			wantErr: ErrNoExpressionFound,
		},
		{
			name: "invalid expression with extra characters",
			args: args{
				re:  expRe,
				str: "2+3a=?",
			},
			want:    nil,
			wantErr: ErrNoExpressionFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseExpression(tt.args.re, tt.args.str)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("parseExpression() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}
