import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import Calculator from './Calculator'
import { vi } from 'vitest'
import { calculate } from '../services/calculatorApi'

vi.mock('../services/calculatorApi', () => ({
    calculate: vi.fn(),
}))

describe('Calculator', () => {
    it('builds an expression when numbers are entered', async () => {
        const user = userEvent.setup()

        render(<Calculator />)

        await user.click(screen.getByRole('button', { name: '8' }))
        await user.click(screen.getByRole('button', { name: '5' }))

        expect(screen.getByText('85')).toBeInTheDocument()
    })

    it('builds an expression when an operator is entered', async () => {
        const user = userEvent.setup()

        render(<Calculator />)

        await user.click(screen.getByRole('button', { name: '8' }))
        await user.click(screen.getByRole('button', { name: '+' }))
        await user.click(screen.getByRole('button', { name: '5' }))

        expect(screen.getByText('8 + 5')).toBeInTheDocument()
    })

    it('clears the calculator state', async () => {
        const user = userEvent.setup()

        render(<Calculator />)

        await user.click(screen.getByRole('button', { name: '8' }))
        await user.click(screen.getByRole('button', { name: '+' }))
        await user.click(screen.getByRole('button', { name: '5' }))
        await user.click(screen.getByRole('button', { name: 'AC' }))

        expect(
            screen.getByText('0', { selector: '.display-value' }),
        ).toBeInTheDocument()
    })

    it('removes the last entered digit', async () => {
        const user = userEvent.setup()

        render(<Calculator />)

        await user.click(screen.getByRole('button', { name: '8' }))
        await user.click(screen.getByRole('button', { name: '5' }))
        await user.click(screen.getByRole('button', { name: '⌫' }))

        expect(
            screen.getByText('8', { selector: '.display-value' }),
        ).toBeInTheDocument()
    })

    it('sends the expression to the backend and displays the result', async () => {
        const user = userEvent.setup()

        vi.mocked(calculate).mockResolvedValue({
            result: 13,
        })

        render(<Calculator />)

        await user.click(screen.getByRole('button', { name: '8' }))
        await user.click(screen.getByRole('button', { name: '+' }))
        await user.click(screen.getByRole('button', { name: '5' }))
        await user.click(screen.getByRole('button', { name: '=' }))

        expect(screen.getByText('13', { selector: '.display-value' }))
            .toBeInTheDocument()
    })

    it('displays an error when the backend rejects the expression', async () => {
        const user = userEvent.setup()

        vi.mocked(calculate).mockRejectedValue(
            new Error('cannot divide by zero'),
        )

        render(<Calculator />)

        await user.click(screen.getByRole('button', { name: '8' }))
        await user.click(screen.getByRole('button', { name: '÷' }))
        await user.click(screen.getByRole('button', { name: '0' }))
        await user.click(screen.getByRole('button', { name: '=' }))

        expect(
            screen.getByText('cannot divide by zero', {
                selector: '.display-value',
            }),
        ).toBeInTheDocument()
    })

    it('builds decimal values correctly', async () => {
        const user = userEvent.setup()

        render(<Calculator />)

        await user.click(screen.getByRole('button', { name: '8' }))
        await user.click(screen.getByRole('button', { name: '.' }))
        await user.click(screen.getByRole('button', { name: '5' }))

        expect(
            screen.getByText('8.5', { selector: '.display-value' }),
        ).toBeInTheDocument()
    })
})