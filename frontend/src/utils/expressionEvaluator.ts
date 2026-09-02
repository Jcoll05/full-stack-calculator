// This file contains a simple expression evaluator that can handle basic arithmetic operations, parentheses, and percentages.
export function evaluateExpression(expression: string): number {
    const tokens = tokenize(expression)

    return parseExpression(tokens).value
}

// Define a structure to hold the parsed value and whether it is a percentage.
interface ParsedValue {
    value: number
    isPercentage: boolean
}

// Tokenize the input expression into an array of tokens for easier parsing.
function tokenize(expression: string): string[] {
    return expression
        .replace(/×/g, '*')
        .replace(/÷/g, '/')
        .replace(/−/g, '-')
        .replace(/√/g, ' √ ')
        .replace(/\(/g, ' ( ')
        .replace(/\)/g, ' ) ')
        .replace(/%/g, ' % ')
        .trim()
        .split(/\s+/)
}

// Parse the expression recursively, respecting operator precedence and handling parentheses.
function parseExpression(tokens: string[]): ParsedValue {
    let result = parseTerm(tokens)

    // Handle addition and subtraction, including calculator-style percentage operations.
    while (tokens.length > 0) {
        if (tokens[0] === ')') {
            break
        }

        // Handle calculator-style percentage operations: e.g., "100 + 10%" means "100 + (100 * 0.10)"
        const operator = tokens.shift()

        // Ensure that the operator is either '+' or '-'.
        if (operator !== '+' && operator !== '-') {
            throw new Error(`Unexpected operator: ${operator}`)
        }
        
        // Handle the next term, which could be a number, a percentage, or an expression in parentheses.
        const nextValue = parseTerm(tokens)

        // If the next value is a percentage, calculate it based on the current result.
        if (nextValue.isPercentage) {
            const percentageValue = result.value * nextValue.value

            // Perform the addition or subtraction with the calculated percentage value.
            if (operator === '+') {
                result = {
                    value: result.value + percentageValue,
                    isPercentage: false,
                }
            } else {
                result = {
                    value: result.value - percentageValue,
                    isPercentage: false,
                }
            }

            continue
        }

        // If the next value is not a percentage, perform the addition or subtraction directly.
        if (operator === '+') {
            result = {
                value: result.value + nextValue.value,
                isPercentage: false,
            }
        } else {
            result = {
                value: result.value - nextValue.value,
                isPercentage: false,
            }
        }
    }

    return result
}

// Parse terms, handling multiplication and division, including implicit multiplication before parentheses.
function parseTerm(tokens: string[]): ParsedValue {
    let result = parsePower(tokens) // Start by parsing the first power term.

    // Handle multiplication and division, including implicit multiplication before parentheses.
    while (tokens.length > 0) {
        const operator = tokens[0]

        // Implicit multiplication:
        // 2(2 + 2) means 2 × (2 + 2)
        if (operator === '(') {
            const nextValue = parsePower(tokens)

            result = {
                value: result.value * nextValue.value,
                isPercentage: false,
            }

            continue
        }

        // Ensure that the operator is either '*' or '/'.
        if (operator !== '*' && operator !== '/') {
            break
        }

        tokens.shift() // Remove the operator from the tokens.

        // Parse the next power term after the operator.
        const nextValue = parsePower(tokens)

        // Perform multiplication or division based on the operator.
        if (operator === '*') {
            result = {
                value: result.value * nextValue.value,
                isPercentage: false,
            }
        } else {
            if (nextValue.value === 0) {
                throw new Error('Cannot divide by zero')
            }

            result = {
                value: result.value / nextValue.value,
                isPercentage: false,
            }
        }
    }

    return result
}

// Parse power expressions, handling exponentiation.
function parsePower(tokens: string[]): ParsedValue {
    let result = parseNumber(tokens)

    // Handle exponentiation, allowing for multiple consecutive exponentiation operations.
    while (tokens.length > 0 && tokens[0] === '^') {
        tokens.shift()

        // Parse the exponent after the '^' operator.
        const exponent = parseNumber(tokens)

        result = {
            value: Math.pow(result.value, exponent.value),
            isPercentage: false,
        }
    }

    return result
}

// Parse numbers, handling square roots, parentheses, and percentages.
function parseNumber(tokens: string[]): ParsedValue {
    const token = tokens.shift() // Get the next token from the tokens array.

    // Handle invalid tokens, square roots, parentheses, and percentages.
    if (token === undefined || token === '') {
        throw new Error('Invalid expression')
    }

    // Handle square root operations, ensuring that the operand is not negative.
    if (token === '√') {
        const operand = parseNumber(tokens)

        // Ensure that the operand for the square root is not negative.
        if (operand.value < 0) {
            throw new Error('Cannot calculate square root of a negative number')
        }

        return {
            value: Math.sqrt(operand.value),
            isPercentage: false,
        }
    }

    // Handle expressions within parentheses, ensuring that they are properly closed and can also be percentages.
    if (token === '(') {
        const result = parseExpression(tokens) // Recursively parse the expression within the parentheses.

        const closingToken = tokens.shift() // Get the next token, which should be the closing parenthesis.

        // Ensure that the closing token is indeed a closing parenthesis.
        if (closingToken !== ')') {
            throw new Error('Missing closing parenthesis')
        }

        // Handle percentages after parentheses, e.g., "(2 + 2)%"
        if (tokens[0] === '%') {
            tokens.shift()

            return {
                value: result.value / 100,
                isPercentage: true,
            }
        }

        return result
    }

    // Parse the token as a number, throwing an error if it is not a valid number.
    const value = Number(token)

    // Ensure that the parsed value is a valid number, throwing an error if it is NaN.
    if (Number.isNaN(value)) {
        throw new Error(`Invalid number: ${token}`)
    }

    // Handle percentages, e.g., "2%" means "0.02"
    if (tokens[0] === '%') {
        tokens.shift()

        return {
            value: value / 100,
            isPercentage: true,
        }
    }

    return {
        value,
        isPercentage: false,
    }
}