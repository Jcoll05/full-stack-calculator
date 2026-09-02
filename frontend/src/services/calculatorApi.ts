import type { CalculateRequest, CalculateResponse } from '../types/calculator' // Import the request and response types for calculator operations.

// Define the base URL for the calculator API.
const API_URL = 'http://localhost:8080/api/v1'

// Function to perform a calculation by sending a request to the backend API.
export async function calculate(
    request: CalculateRequest,
): Promise<CalculateResponse> {

    // Send a POST request to the /calculate endpoint with the request data.
    const response = await fetch(`${API_URL}/calculate`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(request),
    })

    // Check if the response is not OK (status code not in the range 200-299).
    if (!response.ok) {
        const error = await response.json()
        throw new Error(error.error || 'Calculation failed')
    }

    return response.json()
}