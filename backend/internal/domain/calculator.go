package domain

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// Validation errors
var (
	ErrUnsupportedOperation = errors.New("unsupported operation")
	ErrDivisionByZero       = errors.New("cannot divide by zero")
	ErrNegativeSquareRoot   = errors.New("cannot calculate square root of a negative number")

	ErrInvalidExpression         = errors.New("invalid expression")
	ErrMissingClosingParenthesis = errors.New("missing closing parenthesis")
	ErrUnexpectedOperator        = errors.New("unexpected operator")
	ErrInvalidNumber             = errors.New("invalid number")
)

// parsedValue represents a parsed value and whether it originated from a percentage.
type parsedValue struct {
	value        float64
	isPercentage bool
}

// EvaluateExpression evaluates a mathematical expression represented as a string.
// It supports addition, subtraction, multiplication, division, exponentiation,
// square root, and percentage operations. The function returns the result of
// the evaluation or an error if the expression is invalid.
func EvaluateExpression(expression string) (float64, error) {
	tokens := tokenize(expression) // Tokenize the input expression into a slice of strings.

	// If there are no tokens, the expression is invalid.
	if len(tokens) == 0 {
		return 0, ErrInvalidExpression
	}

	result, err := parseExpression(&tokens)

	if err != nil {
		return 0, err
	}

	// If there are leftover tokens after parsing, the expression is invalid.
	if len(tokens) > 0 {
		return 0, ErrInvalidExpression
	}

	return result.value, nil
}

// tokenize converts the raw mathematical expression into a sequence of tokens.
// Each token represents a meaningful part of the expression, such as a number,
// operator, parenthesis, square root, or percentage symbol. Separating the
// expression into tokens allows the parser to process each element individually
// and apply operator precedence and nested-expression rules without having to
// interpret the entire string at once.
func tokenize(expression string) []string {
	expression = strings.ReplaceAll(expression, "×", "*")
	expression = strings.ReplaceAll(expression, "÷", "/")
	expression = strings.ReplaceAll(expression, "−", "-")
	expression = strings.ReplaceAll(expression, "√", " √ ")
	expression = strings.ReplaceAll(expression, "(", " ( ")
	expression = strings.ReplaceAll(expression, ")", " ) ")
	expression = strings.ReplaceAll(expression, "%", " % ")

	expression = strings.TrimSpace(expression)

	// If the expression is empty after trimming, return nil to indicate no tokens.
	if expression == "" {
		return nil
	}

	return strings.Fields(expression)
}

// parseExpression parses the tokens and evaluates the expression according to operator precedence.
func parseExpression(tokens *[]string) (parsedValue, error) {
	result, err := parseTerm(tokens)
	if err != nil {
		return parsedValue{}, err
	}

	// Continue parsing while there are tokens left and the next token is not a closing parenthesis.
	for len(*tokens) > 0 {
		if (*tokens)[0] == ")" {
			break
		}

		operator := consumeToken(tokens) // Consume the next token as the operator.

		// If the operator is not addition or subtraction, return an unexpected operator error.
		if operator != "+" && operator != "-" {
			return parsedValue{}, ErrUnexpectedOperator
		}

		nextValue, err := parseTerm(tokens) // Parse the next term in the expression.

		if err != nil {
			return parsedValue{}, err
		}

		// Handle percentage calculations based on the current operator and the next value.
		if nextValue.isPercentage {
			percentageValue := result.value * nextValue.value

			// If the next value is a percentage, calculate its effect on the current result based on the operator.
			if operator == "+" {
				result.value += percentageValue
			} else {
				result.value -= percentageValue
			}

			result.isPercentage = false
			continue
		}

		// Perform the addition or subtraction based on the operator.
		if operator == "+" {
			result.value += nextValue.value
		} else {
			result.value -= nextValue.value
		}

		result.isPercentage = false
	}

	return result, nil
}

// parseTerm parses a term in the expression, handling multiplication and division.
func parseTerm(tokens *[]string) (parsedValue, error) {
	result, err := parsePower(tokens) // Parse the next power in the expression.

	if err != nil {
		return parsedValue{}, err
	}

	// Continue parsing while there are tokens left and the next token is a multiplication or division operator.
	for len(*tokens) > 0 {
		operator := (*tokens)[0]

		// Implicit multiplication:
		// 2(2 + 2) means 2 × (2 + 2).
		if operator == "(" {
			nextValue, err := parsePower(tokens)
			if err != nil {
				return parsedValue{}, err
			}

			result.value *= nextValue.value
			result.isPercentage = false

			continue
		}

		// Stop when the next token is not a multiplication or division operator.
		if operator != "*" && operator != "/" {
			break
		}

		consumeToken(tokens)

		nextValue, err := parsePower(tokens)

		if err != nil {
			return parsedValue{}, err
		}

		// Perform multiplication or division based on the operator.
		if operator == "*" {
			result.value *= nextValue.value
		} else {
			if nextValue.value == 0 {
				return parsedValue{}, ErrDivisionByZero
			}

			result.value /= nextValue.value
		}

		result.isPercentage = false
	}

	return result, nil
}

// parsePower parses a power operation in the expression, handling exponentiation.
func parsePower(tokens *[]string) (parsedValue, error) {
	result, err := parseNumber(tokens) // Parse the next number in the expression.
	if err != nil {
		return parsedValue{}, err
	}

	for len(*tokens) > 0 && (*tokens)[0] == "^" {
		consumeToken(tokens)

		exponent, err := parseNumber(tokens)

		if err != nil {
			return parsedValue{}, err
		}

		result.value = math.Pow(result.value, exponent.value)
		result.isPercentage = false
	}

	return result, nil
}

// parseNumber parses a number or a sub-expression in the expression, handling square roots and parentheses.
func parseNumber(tokens *[]string) (parsedValue, error) {

	// If there are no tokens left, return an invalid expression error.
	if len(*tokens) == 0 {
		return parsedValue{}, ErrInvalidExpression
	}

	token := consumeToken(tokens)

	// If the token is an operator, return an unexpected operator error.
	if token == "+" || token == "-" || token == "*" || token == "/" || token == "^" {
		return parsedValue{}, ErrUnexpectedOperator
	}

	// Handle square root operation
	if token == "√" {
		operand, err := parseNumber(tokens)
		if err != nil {
			return parsedValue{}, err
		}

		// If the operand is negative, return an error for negative square root.
		if operand.value < 0 {
			return parsedValue{}, ErrNegativeSquareRoot
		}

		return parsedValue{
			value:        math.Sqrt(operand.value),
			isPercentage: false,
		}, nil
	}

	// Handle parentheses
	if token == "(" {
		result, err := parseExpression(tokens)

		if err != nil {
			return parsedValue{}, err
		}

		// If there are no tokens left after parsing, return a missing closing parenthesis error.
		if len(*tokens) == 0 {
			return parsedValue{}, ErrMissingClosingParenthesis
		}

		consumeToken(tokens)

		// If the next token is a percentage sign, consume it and return the result as a percentage.
		if len(*tokens) > 0 && (*tokens)[0] == "%" {
			consumeToken(tokens)

			return parsedValue{
				value:        result.value / 100,
				isPercentage: true,
			}, nil
		}

		return result, nil
	}

	value, err := strconv.ParseFloat(token, 64) // Parse the token as a float64 value.

	// If there was an error during parsing, return an invalid number error.
	if err != nil {
		return parsedValue{}, ErrInvalidNumber
	}

	// Handle percentage
	if len(*tokens) > 0 && (*tokens)[0] == "%" {
		consumeToken(tokens)

		return parsedValue{
			value:        value / 100,
			isPercentage: true,
		}, nil
	}

	return parsedValue{
		value:        value,
		isPercentage: false,
	}, nil
}

// consumeToken removes and returns the first token from the slice of tokens.
func consumeToken(tokens *[]string) string {
	token := (*tokens)[0]
	*tokens = (*tokens)[1:]
	return token
}
