import '../stylesheets/display.css'

// Define the props for the Display component.
interface DisplayProps {
    value: string
}

// Display component to show the current value on the calculator.
function Display({ value }: DisplayProps) {
    return (
        <div className="display">
            <span className="display-value">{value}</span>
        </div>
    )
}

export default Display