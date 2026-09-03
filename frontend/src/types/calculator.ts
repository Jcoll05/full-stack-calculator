// This file contains TypeScript types for the calculator application.

// The types defined here are used for request and response payloads in the calculator API.
export interface CalculateRequest {
    expression: string
}

// The response from the calculator API contains the result of the evaluated expression.
export interface CalculateResponse {
    result: number
}