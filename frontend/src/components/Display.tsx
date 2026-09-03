import { useEffect, useRef } from 'react'

import '../stylesheets/display.css'

// Define the props for the Display component.
interface DisplayProps {
    value: string
}

// Display component to show the current value on the calculator.
function Display({ value }: DisplayProps) {
    const displayRef = useRef<HTMLSpanElement>(null)

    // Use an effect to scroll the display to the right whenever the value changes.
    useEffect(() => {
        const display = displayRef.current

        if (!display) {
            return
        }

        display.scrollLeft = display.scrollWidth
    }, [value])

    return (
        <div className="display">
            <span ref={displayRef} className="display-value">
                {value}
            </span>
        </div>
    )
}

export default Display