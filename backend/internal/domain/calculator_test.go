package domain

import (
	"errors"
	"testing"
)

// TestEvaluateExpression tests the EvaluateExpression function with various mathematical expressions.
func TestEvaluateExpression(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   float64
	}{
		{
			name:       "respects multiplication precedence",
			expression: "2 + 5 × 2",
			expected:   12,
		},
		{
			name:       "respects division precedence",
			expression: "10 + 6 ÷ 2",
			expected:   13,
		},
		{
			name:       "evaluates addition and subtraction",
			expression: "10 − 3 + 2",
			expected:   9,
		},
		{
			name:       "evaluates multiplication and division",
			expression: "20 ÷ 2 × 3",
			expected:   30,
		},
		{
			name:       "supports implicit multiplication before parentheses",
			expression: "2(2 + 2)",
			expected:   8,
		},
		{
			name:       "supports implicit multiplication after parentheses",
			expression: "(2 + 2)2",
			expected:   8,
		},
		{
			name:       "supports implicit multiplication after and before parentheses",
			expression: "2(2 + 2)2",
			expected:   16,
		},
		{
			name:       "supports zero after parentheses",
			expression: "(5 + 5)0",
			expected:   0,
		},
		{
			name:       "supports implicit multiplication between parentheses",
			expression: "(5 + 5)(2 + 2)",
			expected:   40,
		},
		{
			name:       "supports implicit multiplication after square root",
			expression: "√(25)2",
			expected:   10,
		},
		{
			name:       "evaluates standalone percentages",
			expression: "2%",
			expected:   0.02,
		},
		{
			name:       "evaluates calculator-style percentage addition",
			expression: "2% + 2%",
			expected:   0.0204,
		},
		{
			name:       "evaluates calculator-style percentage addition with a base value",
			expression: "100 + 10%",
			expected:   110,
		},
		{
			name:       "evaluates calculator-style percentage subtraction",
			expression: "100 − 10%",
			expected:   90,
		},
		{
			name:       "evaluates square roots",
			expression: "√9",
			expected:   3,
		},
		{
			name:       "evaluates square roots with parentheses",
			expression: "√(9 + 7)",
			expected:   4,
		},
		{
			name:       "evaluates square roots inside an expression",
			expression: "2 × √9",
			expected:   6,
		},
		{
			name:       "evaluates powers",
			expression: "2 ^ 3",
			expected:   8,
		},
		{
			name:       "evaluates powers before multiplication",
			expression: "2 × 3 ^ 2",
			expected:   18,
		},
		{
			name:       "evaluates powers before addition",
			expression: "2 + 3 ^ 2",
			expected:   11,
		},
		{
			name:       "evaluates powers inside parentheses",
			expression: "(2 + 3) ^ 2",
			expected:   25,
		},
		{
			name:       "evaluates powers with multiplication inside parentheses",
			expression: "(2 × 3) ^ 2",
			expected:   36,
		},
		{
			name:       "evaluates a complex expression with square root and powers",
			expression: "√(25 + 75) × 2 + 10% − 5 ÷ 2 ^ 2",
			expected:   20.75,
		},
		{
			name:       "evaluates powers inside a larger expression",
			expression: "2 + 3 ^ 2 × 4",
			expected:   38,
		},
		{
			name:       "evaluates parentheses with multiple operators and powers",
			expression: "(2 + 3) ^ 2 × 4 − 10",
			expected:   90,
		},
		{
			name:       "evaluates nested parentheses with powers",
			expression: "2 × ((3 + 2) ^ 2)",
			expected:   50,
		},
		{
			name:       "evaluates percentage multiplication",
			expression: "100 × 10%",
			expected:   10,
		},
		{
			name:       "evaluates percentage division",
			expression: "100 ÷ 10%",
			expected:   1000,
		},
		{
			name:       "evaluates percentage calculations with powers",
			expression: "100 + 10% × 2 ^ 2",
			expected:   100.4,
		},
		{
			name:       "evaluates a square root with a power",
			expression: "√9 ^ 2",
			expected:   9,
		},
		{
			name:       "evaluates a square root inside a powered expression",
			expression: "(√9 + 1) ^ 2",
			expected:   16,
		},
		{
			name:       "evaluates multiple square roots in an expression",
			expression: "√9 + √16",
			expected:   7,
		},
		{
			name:       "evaluates percentage of parenthesized expression",
			expression: "(2 + 2)%",
			expected:   0.04,
		},
		{
			name:       "boss fight expression",
			expression: "√((16 + 9) × 4) + 50% × (20 − 5) − 2 ^ 3 ÷ 4",
			expected:   15.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := EvaluateExpression(tt.expression)

			if err != nil {
				t.Fatalf("EvaluateExpression() returned unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf(
					"EvaluateExpression(%q) = %v, expected %v",
					tt.expression,
					result,
					tt.expected,
				)
			}
		})
	}
}

// TestEvaluateExpressionErrors tests the EvaluateExpression function for expected error cases.
func TestEvaluateExpressionErrors(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   error
	}{
		{
			name:       "division by zero",
			expression: "10 ÷ 0",
			expected:   ErrDivisionByZero,
		},
		{
			name:       "negative square root",
			expression: "√−9",
			expected:   ErrNegativeSquareRoot,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := EvaluateExpression(tt.expression)

			if !errors.Is(err, tt.expected) {
				t.Errorf(
					"EvaluateExpression(%q) error = %v, expected %v",
					tt.expression,
					err,
					tt.expected,
				)
			}
		})
	}
}

// TestEvaluateExpressionInvalidExpressions tests the EvaluateExpression function for various invalid expressions.
func TestEvaluateExpressionInvalidExpressions(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   error
	}{
		{
			name:       "empty expression",
			expression: "",
			expected:   ErrInvalidExpression,
		},
		{
			name:       "unexpected operator",
			expression: "2 3",
			expected:   ErrUnexpectedOperator,
		},
		{
			name:       "unexpected operator where operand is expected",
			expression: "2 × + 3",
			expected:   ErrUnexpectedOperator,
		},
		{
			name:       "missing operand after addition",
			expression: "2 +",
			expected:   ErrInvalidExpression,
		},
		{
			name:       "missing operand after multiplication",
			expression: "2 ×",
			expected:   ErrInvalidExpression,
		},
		{
			name:       "missing exponent",
			expression: "2 ^",
			expected:   ErrInvalidExpression,
		},
		{
			name:       "invalid number",
			expression: "abc",
			expected:   ErrInvalidNumber,
		},
		{
			name:       "missing operand after square root",
			expression: "√",
			expected:   ErrInvalidExpression,
		},
		{
			name:       "missing closing parenthesis",
			expression: "(2 + 3",
			expected:   ErrMissingClosingParenthesis,
		},
		{
			name:       "unexpected closing parenthesis",
			expression: "2 + 3)",
			expected:   ErrInvalidExpression,
		},
		{
			name:       "invalid operator after expression",
			expression: "2 * 3 *",
			expected:   ErrInvalidExpression,
		},
		{
			name:       "invalid expression inside parentheses",
			expression: "(2 + )",
			expected:   ErrInvalidNumber,
		},
		{
			name:       "invalid expression after implicit multiplication",
			expression: "2(abc)",
			expected:   ErrInvalidNumber,
		},
		{
			name:       "invalid expression after square root",
			expression: "√abc",
			expected:   ErrInvalidNumber,
		},
		{
			name:       "unexpected closing parenthesis between operands",
			expression: "2)3",
			expected:   ErrInvalidExpression,
		},
		{
			name:       "unexpected whitespace between operands inside parentheses",
			expression: "(2 3)",
			expected:   ErrUnexpectedOperator,
		},
		{
			name:       "missing closing parenthesis after implicit multiplication",
			expression: "2(3",
			expected:   ErrMissingClosingParenthesis,
		},
		{
			name:       "unexpected closing parenthesis before expression",
			expression: ")(2",
			expected:   ErrInvalidNumber,
		},
		{
			name:       "consecutive multiplication operators",
			expression: "2**3",
			expected:   ErrInvalidNumber,
		},
		{
			name:       "consecutive multiplication and division operators",
			expression: "2*/3",
			expected:   ErrInvalidNumber,
		},
		{
			name:       "empty parentheses",
			expression: "()",
			expected:   ErrInvalidNumber,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := EvaluateExpression(tt.expression)

			if !errors.Is(err, tt.expected) {
				t.Errorf(
					"EvaluateExpression(%q) error = %v, expected %v",
					tt.expression,
					err,
					tt.expected,
				)
			}
		})
	}
}
