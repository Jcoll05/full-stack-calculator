import { useState } from 'react'
import '../stylesheets/calculator.css'
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

    // Handles input from the calculator buttons.
    const handleInput = (value: string) => {
        if (value === 'clear') {
            setDisplayValue('0')
            return
        }

        // Handle backspace input
        if (value === 'backspace') {
            setDisplayValue((current) =>
                current.length > 1 ? current.slice(0, -1) : '0',
            )
            return
        }

        // Handle decimal point input
        if (value === '.') {
            if (!displayValue.includes('.')) {
                setDisplayValue((current) => `${current}.`)
            }
            return
        }

        // Handle number input
        if (/^\d$/.test(value)) {
            setDisplayValue((current) =>
                current === '0' ? value : `${current}${value}`,
            )
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