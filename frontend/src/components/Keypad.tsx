import '../stylesheets/keypad.css'

// Define the structure of a calculator button.
interface CalculatorButton {
    label: string
    value: string
    type: 'number' | 'operator' | 'action' | 'equals'
}

// Define the props for the Keypad component.
interface KeypadProps {
    buttons: CalculatorButton[]
    onInput: (value: string) => void
}

// Keypad component to render the calculator buttons.
function Keypad({ buttons, onInput }: KeypadProps) {
    return (
        <div className="keypad">
            {buttons.map((button) => (
                <button
                    key={button.value}
                    type="button"
                    className={`keypad-button keypad-button--${button.type} keypad-button--${button.value}`}
                    onClick={() => onInput(button.value)}
                >
                    {button.label}
                </button>
            ))}
        </div>
    )
}

export default Keypad