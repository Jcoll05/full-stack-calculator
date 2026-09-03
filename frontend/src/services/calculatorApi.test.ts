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
})