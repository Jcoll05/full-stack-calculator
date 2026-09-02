import { useState } from 'react'
import '../stylesheets/calculator.css'
import { calculate } from '../services/calculatorApi'
import { evaluateExpression } from '../utils/expressionEvaluator'
import type { CalculatorOperation } from '../types/calculator'
import Display from './Display'
import Keypad from './Keypad'

// Define the structure of a calculator button.
interface CalculatorButton {
    label: string
    value: string
    type: 'number' | 'operator' | 'action' | 'equals'
}


function Calculator() {

    // Stores the value currently shown on the calculator display.
    const [displayValue, setDisplayValue] = useState('0')
    // Stores the current number being entered.
    const [currentValue, setCurrentValue] = useState('0')
    // Stores the previous value and the current operation for calculations.
    const [previousValue, setPreviousValue] = useState<number | null>(null)
    // Stores the current operation selected by the user.
    const [operation, setOperation] = useState<CalculatorOperation | null>(null)
    // Stores whether the next number should start a new operand.
    const [waitingForOperand, setWaitingForOperand] = useState(false)
    // Stores the full expression being built for display purposes.
    const [expression, setExpression] = useState('')
    // Stores whether the user is currently entering a parenthesized expression.
    const [insideParentheses, setInsideParentheses] = useState(false)

    // Function to get the symbol for a given operator value.
    const getOperatorSymbol = (value: string) => {
        switch (value) {
            case 'add':
                return '+'
            case 'subtract':
                return '−'
            case 'multiply':
                return '×'
            case 'divide':
                return '÷'
            case 'power':
                return '^'
            default:
                return value
        }
    }

    // Handles input from the calculator buttons.
    const handleInput = async (value: string) => {

        // Handle clear input
        if (value === 'clear') {
            setDisplayValue('0')
            setCurrentValue('0')
            setPreviousValue(null)
            setOperation(null)
            setWaitingForOperand(false)
            setExpression('')
            setInsideParentheses(false)
            return
        }

        // Handle backspace input
        if (value === 'backspace') {
            // Nothing to delete.
            if (expression === '') {
                return
            }

            // Remove the last character from the expression.
            // Operators have a space after them, so remove both the space and operator.
            const trimmedExpression = expression.trimEnd()

            const nextExpression =
                ['+', '−', '×', '÷', '^'].includes(trimmedExpression.slice(-1))
                    ? trimmedExpression.slice(0, -1).trimEnd()
                    : expression.slice(0, -1)

            // Update the expression and display.
            setExpression(nextExpression)
            setDisplayValue(nextExpression || '0')

            // If the deleted character was a closing parenthesis,
            // reopen the parenthesized expression.
            if (expression.endsWith(')')) {
                setInsideParentheses(true)
            }

            // If the expression now ends with an operator,
            // the user is waiting for the next operand.
            const lastCharacter = nextExpression.trim().slice(-1)

            if (['+', '−', '×', '÷', '^'].includes(lastCharacter)) {
                setWaitingForOperand(true)
                setCurrentValue('0')
                return
            }

            // If the expression contains an operator, determine
            // the current number being entered from the expression.
            const parts = nextExpression.trim().split(' ')
            const lastPart = parts[parts.length - 1]

            if (!isNaN(Number(lastPart))) {
                setCurrentValue(lastPart)
            } else {
                setCurrentValue('0')
            }

            setWaitingForOperand(false)
            return
        }

        // Handle decimal point input
        if (value === '.') {
            // Prevent adding multiple decimal points to the current value.
            if (!currentValue.includes('.')) {
                const nextValue = `${currentValue}.`

                const nextExpression =
                    expression === '' ? nextValue : `${expression}.`

                setCurrentValue(nextValue)
                setExpression(nextExpression)
                setDisplayValue(nextExpression)
            }

            return
        }

        // Handle operator input
        if (
            value === 'add' ||
            value === 'subtract' ||
            value === 'multiply' ||
            value === 'divide' ||
            value === 'power'
        ) {
            const nextExpression =
                expression === ''
                    ? `${currentValue} ${getOperatorSymbol(value)}`
                    : `${expression} ${getOperatorSymbol(value)}`

            setPreviousValue(Number(currentValue))
            setOperation(value as CalculatorOperation)
            setExpression(nextExpression)
            setDisplayValue(nextExpression)
            setWaitingForOperand(true)
            return
        }

        // Handle equals input
        if (value === 'equals') {
            if (expression === '') return

            try {
                const result = evaluateExpression(expression)

                setDisplayValue(String(result))
                setCurrentValue(String(result))
                setPreviousValue(null)
                setOperation(null)
                setWaitingForOperand(false)
                setExpression(String(result))
            } catch (error) {
                const message =
                    error instanceof Error
                        ? error.message
                        : 'Invalid expression'

                setDisplayValue(message)
            }

            return
        }

        // Handle number input
        if (/^\d$/.test(value)) {
            // If waiting for the next operand, start a new current value.
            if (waitingForOperand) {
                const nextExpression = `${expression} ${value}`

                setCurrentValue(value)
                setExpression(nextExpression)
                setDisplayValue(nextExpression)
                setWaitingForOperand(false)
                return
            }

            // Replace the leading zero instead of creating values such as "09".
            const nextValue =
                currentValue === '0'
                    ? value
                    : `${currentValue}${value}`

            // Build the next expression.
            let nextExpression: string

            if (expression === '') {
                nextExpression = nextValue
            } else if (currentValue === '0' && expression.endsWith('0')) {
                // Replace the current operand's leading zero.
                nextExpression =
                    expression.slice(0, -1) + nextValue
            } else {
                // Append the digit normally.
                nextExpression = `${expression}${value}`
            }

            setCurrentValue(nextValue)
            setExpression(nextExpression)
            setDisplayValue(nextExpression)

            return
        }

        // Handle square root input
        if (value === 'sqrt') {
            const nextExpression = `${expression}√`

            setExpression(nextExpression)
            setDisplayValue(nextExpression)
            setCurrentValue('0')
            setWaitingForOperand(true)

            return
        }


        // Handle percentage input
        if (value === 'percentage') {
            const nextExpression = `${expression}%`

            setExpression(nextExpression)
            setDisplayValue(nextExpression)

            return
        }

        // Handle opening parenthesis
        if (value === '(') {
            setExpression(`${expression}(`)
            setDisplayValue(`${expression}(`)
            setInsideParentheses(true)
            setCurrentValue('0')
            setPreviousValue(null)
            setOperation(null)
            setWaitingForOperand(false)
            return
        }

        // Handle closing parenthesis
        if (value === ')') {
            if (
                !insideParentheses ||
                previousValue === null ||
                operation === null
            ) {
                return
            }

            const response = await calculate({
                operation,
                a: previousValue,
                b: Number(currentValue),
            })

            const nextExpression = `${expression})`

            setExpression(nextExpression)
            setDisplayValue(nextExpression)
            setCurrentValue(String(response.result))
            setPreviousValue(null)
            setOperation(null)
            setWaitingForOperand(false)
            setInsideParentheses(false)

            return
        }
    }


    // Define the buttons for the calculator keypad.
    const buttons: CalculatorButton[] = [
        { label: 'AC', value: 'clear', type: 'action' },
        { label: '(', value: '(', type: 'action' },
        { label: ')', value: ')', type: 'action' },
        { label: '⌫', value: 'backspace', type: 'action' },

        { label: '√', value: 'sqrt', type: 'operator' },
        { label: 'xʸ', value: 'power', type: 'operator' },
        { label: '%', value: 'percentage', type: 'operator' },
        { label: '÷', value: 'divide', type: 'operator' },

        { label: '7', value: '7', type: 'number' },
        { label: '8', value: '8', type: 'number' },
        { label: '9', value: '9', type: 'number' },
        { label: '×', value: 'multiply', type: 'operator' },

        { label: '4', value: '4', type: 'number' },
        { label: '5', value: '5', type: 'number' },
        { label: '6', value: '6', type: 'number' },
        { label: '−', value: 'subtract', type: 'operator' },

        { label: '1', value: '1', type: 'number' },
        { label: '2', value: '2', type: 'number' },
        { label: '3', value: '3', type: 'number' },
        { label: '+', value: 'add', type: 'operator' },

        { label: '0', value: '0', type: 'number' },
        { label: '.', value: '.', type: 'number' },
        { label: '=', value: 'equals', type: 'equals' },
    ]

    return (
        <div className="calculator-container">
            <section className="calculator">
                <Display value={displayValue} />
                <Keypad buttons={buttons} onInput={handleInput} />
            </section>
        </div>
    )
}

export default Calculator