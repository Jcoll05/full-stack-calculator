import { describe, expect, it, vi } from 'vitest'
import { calculate } from './calculatorApi'

describe('calculatorApi', () => {
    it('sends an expression to the backend and returns the result', async () => {
        const fetchMock = vi.fn().mockResolvedValue({
            ok: true,
            json: async () => ({ result: 16 }),
        })

        vi.stubGlobal('fetch', fetchMock)

        const result = await calculate({
            expression: '8 × 2',
        })

        expect(fetchMock).toHaveBeenCalledWith(
            'http://localhost:8080/api/v1/calculate',
            {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    expression: '8 × 2',
                }),
            },
        )

        expect(result).toEqual({ result: 16 })
    })

    it('throws the backend error when the response is not ok', async () => {
        vi.stubGlobal(
            'fetch',
            vi.fn().mockResolvedValue(
                new Response(
                    JSON.stringify({
                        error: 'cannot divide by zero',
                    }),
                    { status: 400 },
                ),
            ),
        )

        await expect(
            calculate({ expression: '8 ÷ 0' }),
        ).rejects.toThrow('cannot divide by zero')

        vi.unstubAllGlobals()
    })

    it('uses a fallback error when the backend provides no error message', async () => {
        vi.stubGlobal(
            'fetch',
            vi.fn().mockResolvedValue(
                new Response(
                    JSON.stringify({}),
                    { status: 400 },
                ),
            ),
        )

        await expect(
            calculate({ expression: 'invalid' }),
        ).rejects.toThrow('Calculation failed')

        vi.unstubAllGlobals()
    })
})