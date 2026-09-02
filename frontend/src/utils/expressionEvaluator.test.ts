import { describe, expect, it } from 'vitest'
import { evaluateExpression } from './expressionEvaluator'

// Test suite for the evaluateExpression function, which evaluates mathematical expressions.
describe('evaluateExpression', () => {
    it('respects multiplication precedence', () => {
        expect(evaluateExpression('2 + 5 × 2')).toBe(12)
    })

    it('respects division precedence', () => {
        expect(evaluateExpression('10 + 6 ÷ 2')).toBe(13)
    })

    it('evaluates addition and subtraction', () => {
        expect(evaluateExpression('10 − 3 + 2')).toBe(9)
    })

    it('evaluates multiplication and division', () => {
        expect(evaluateExpression('20 ÷ 2 × 3')).toBe(30)
    })

    it('throws when dividing by zero', () => {
        expect(() => evaluateExpression('10 ÷ 0')).toThrow(
            'Cannot divide by zero',
        )
    })

    it('supports implicit multiplication before parentheses', () => {
        expect(evaluateExpression('2(2 + 2)')).toBe(8)
    })

    it('evaluates standalone percentages', () => {
        expect(evaluateExpression('2%')).toBe(0.02)
    })

    it('evaluates calculator-style percentage addition', () => {
        expect(evaluateExpression('2% + 2%')).toBe(0.0204)
    })

    it('evaluates calculator-style percentage addition with a base value', () => {
        expect(evaluateExpression('100 + 10%')).toBe(110)
    })

    it('evaluates calculator-style percentage subtraction', () => {
        expect(evaluateExpression('100 − 10%')).toBe(90)
    })

    it('evaluates square roots', () => {
        expect(evaluateExpression('√9')).toBe(3)
    })

    it('evaluates square roots with parentheses', () => {
        expect(evaluateExpression('√(9 + 7)')).toBe(4)
    })

    it('evaluates square roots inside an expression', () => {
        expect(evaluateExpression('2 × √9')).toBe(6)
    })

    it('throws when calculating the square root of a negative number', () => {
        expect(() => evaluateExpression('√−9')).toThrow(
            'Cannot calculate square root of a negative number',
        )
    })

    it('evaluates powers', () => {
        expect(evaluateExpression('2 ^ 3')).toBe(8)
    })

    it('evaluates powers before multiplication', () => {
        expect(evaluateExpression('2 × 3 ^ 2')).toBe(18)
    })

    it('evaluates powers before addition', () => {
        expect(evaluateExpression('2 + 3 ^ 2')).toBe(11)
    })

    it('evaluates powers inside parentheses', () => {
        expect(evaluateExpression('(2 + 3) ^ 2')).toBe(25)
    })

    it('evaluates powers with multiplication inside parentheses', () => {
        expect(evaluateExpression('(2 × 3) ^ 2')).toBe(36)
    })

    it('evaluates a complex expression with square root and powers', () => {
        expect(
            evaluateExpression('√(25 + 75) × 2 + 10% − 5 ÷ 2 ^ 2')
        ).toBe(20.75)
    })

    it('evaluates powers inside a larger expression', () => {
        expect(
            evaluateExpression('2 + 3 ^ 2 × 4')
        ).toBe(38)
    })

    it('evaluates parentheses with multiple operators and powers', () => {
        expect(
            evaluateExpression('(2 + 3) ^ 2 × 4 − 10')
        ).toBe(90)
    })

    it('evaluates nested parentheses with powers', () => {
        expect(
            evaluateExpression('2 × ((3 + 2) ^ 2)')
        ).toBe(50)
    })

    it('evaluates a percentage as a standalone value', () => {
        expect(evaluateExpression('2%')).toBe(0.02)
    })

    it('evaluates percentage addition', () => {
        expect(evaluateExpression('100 + 10%')).toBe(110)
    })

    it('evaluates percentage subtraction', () => {
        expect(evaluateExpression('100 − 10%')).toBe(90)
    })

    it('evaluates percentage multiplication', () => {
        expect(evaluateExpression('100 × 10%')).toBe(10)
    })

    it('evaluates percentage division', () => {
        expect(evaluateExpression('100 ÷ 10%')).toBe(1000)
    })

    it('evaluates percentage calculations with powers', () => {
        expect(evaluateExpression('100 + 10% × 2 ^ 2')).toBe(100.4)
    })

    it('evaluates a square root with a power', () => {
        expect(evaluateExpression('√9 ^ 2')).toBe(9)
    })

    it('evaluates a square root inside a powered expression', () => {
        expect(evaluateExpression('(√9 + 1) ^ 2')).toBe(16)
    })

    it('evaluates multiple square roots in an expression', () => {
        expect(evaluateExpression('√9 + √16')).toBe(7)
    })
})