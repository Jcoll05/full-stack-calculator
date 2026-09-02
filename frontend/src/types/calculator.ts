// This file contains TypeScript types for the calculator application.
export type CalculatorOperation =
    | 'add'
    | 'subtract'
    | 'multiply'
    | 'divide'
    | 'power'
    | 'sqrt'
    | 'percentage'

// Define the request structure for a calculation operation.
export interface CalculateRequest {
    operation: CalculatorOperation
    a: number
    b?: number
}

// Define the response structure for a calculation operation.
export interface CalculateResponse {
    result: number
}