# Full-Stack Calculator

A full-stack calculator application with a React frontend and a Go REST API backend. The frontend is responsible for building mathematical expressions and presenting results, while the backend is responsible for validating and evaluating expressions.

## Overview

The application provides a calculator interface supporting basic and advanced mathematical operations.

Supported operations include:

* Addition
* Subtraction
* Multiplication
* Division
* Exponentiation
* Square root
* Percentage
* Parentheses and nested expressions
* Decimal numbers

The frontend sends complete mathematical expressions to the backend, which evaluates them and returns the result as JSON.

## Tech Stack

### Frontend

* React
* TypeScript
* Vite
* Vitest
* React Testing Library
* Testing Library User Event
* jsdom

### Backend

* Go
* Standard library HTTP server
* Go's built-in testing package

## Architecture

The project is divided into two independent applications:

```text
full-stack-calculator/
├── backend/
└── frontend/
```

### Backend Architecture

The backend follows a lightweight Clean Architecture structure:

```text
backend/
├── go.mod
├── cmd/
│   └── server/
│       └── main.go
└── internal/
    ├── domain/
    │   ├── calculator.go
    │   └── calculator_test.go
    ├── application/
    │   ├── calculator_service.go
    │   └── calculator_service_test.go
    └── infrastructure/
        └── http/
            ├── handler.go
            ├── handler_test.go
            ├── middleware.go
            ├── middleware_test.go
            ├── routes.go
            └── routes_test.go
```


The main responsibilities are separated as follows:

* **Domain:** Mathematical expression parsing, evaluation, validation, and domain errors.
* **Application:** Provides the use-case/service interface between the HTTP layer and domain logic.
* **Infrastructure:** HTTP handlers, routing, middleware, JSON serialization, and HTTP error handling.
* **cmd/server:** Application entry point and server startup.
* **go.mod:** Defines the Go module and its Go version.

### Frontend Architecture

The frontend separates UI components, API communication, styles, and shared types:

```text
frontend/src/
├── components/
│   ├── Calculator.test.tsx
│   ├── Calculator.tsx
│   ├── Display.tsx
│   └── Keypad.tsx
├── services/
│   ├── calculatorApi.test.ts
│   └── calculatorApi.ts
├── stylesheets/
│   ├── calculator.css
│   ├── display.css
│   └── keypad.css
├── test/
│   └── setup.ts
├── types/
│   └── calculator.ts
├── App.tsx
├── App.css
├── index.css
└── main.tsx
```

`Calculator` manages calculator interaction and expression state, while `Display` and `Keypad` focus primarily on presentation. API communication is isolated in `calculatorApi.ts`.

Frontend tests are co-located with the functionality they cover. `Calculator.test.tsx` verifies calculator behavior and API interaction from the user's perspective, while `calculatorApi.test.ts` verifies the frontend's HTTP communication with the backend. `test/setup.ts` configures the testing environment.

## Features

### Calculator

The calculator supports:

* Multi-digit numbers
* Decimal values
* Basic arithmetic operators
* Powers
* Square roots
* Percentages
* Parentheses
* Nested expressions
* Backspace
* Clear
* Horizontally scrolling expressions on the display
* Responsive mobile layout
* Backend-based expression evaluation
* User-facing error messages

### Expression Evaluation

Expressions are evaluated by the backend rather than by the frontend.

For example:

```text
√((16 + 9) × 4) + 50% × (20 − 5) − 2 ^ 3 ÷ 4
```

The backend evaluates the complete expression and returns the resulting value.

## Getting Started

### Prerequisites

Make sure the following are installed:

* Node.js and npm
* Go

### Clone the repository

```bash
git clone https://github.com/Jcoll05/full-stack-calculator.git
cd full-stack-calculator
```

## Running the Backend

From the project root:

```bash
cd backend
go run ./cmd/server
```

The backend starts the REST API locally.

The API base URL is:

```text
http://localhost:8080/api/v1
```

## Running the Frontend

Open a second terminal and run:

```bash
cd frontend
npm install
npm run dev
```

Vite will provide the local development URL in the terminal.

The frontend expects the backend to be available at:

```text
http://localhost:8080
```

### Building the Frontend

To create a production build:

```bash
npm run build
```


## API Documentation

### POST `/api/v1/calculate`

Evaluates a complete mathematical expression.

### HTTP Status Codes

| Status | Meaning |
|---|---|
| `200 OK` | Expression evaluated successfully |
| `400 Bad Request` | Invalid request, expression, or mathematical operation |
| `405 Method Not Allowed` | HTTP method is not supported |
| `500 Internal Server Error` | Unexpected server-side error |

#### Request

```http
POST /api/v1/calculate
Content-Type: application/json
```

Request body:

```json
{
  "expression": "8 + 5"
}
```

#### Successful Response

```json
{
  "result": 13
}
```

### Example with multiplication

```bash
curl -X POST http://localhost:8080/api/v1/calculate \
  -H "Content-Type: application/json" \
  -d "{\"expression\":\"8 × 2\"}"
```

Response:

```json
{
  "result": 16
}
```

### Example with a complex expression

```bash
curl -X POST http://localhost:8080/api/v1/calculate \
  -H "Content-Type: application/json" \
  -d "{\"expression\":\"√((16 + 9) × 4) + 50% × (20 − 5) − 2 ^ 3 ÷ 4\"}"
```

Response:

```json
{
  "result": 15.5
}
```

### Error Response

Invalid expressions and mathematical edge cases return a `400 Bad Request` response.

For example, division by zero:

```json
{
  "expression": "8 ÷ 0"
}
```

Response:

```json
{
  "error": "cannot divide by zero"
}
```

Other validation cases include:

* Invalid expressions
* Invalid numbers
* Unexpected operators
* Missing closing parentheses
* Square root of a negative number
* Missing expression

## Testing

### Backend

Run all backend tests:

```bash
cd backend
go test ./...
```

Run tests with coverage:

```bash
go test ./... -cover
```

The core backend packages currently have full statement coverage.

### Frontend

From the `frontend` directory:

```bash
npm test -- --run
```

The frontend tests use Vitest, React Testing Library, and jsdom.

The tests cover key calculator behavior including:

* Number input
* Operator input
* Clearing calculator state
* Backspace
* Decimal input
* Successful API-backed calculation
* API error handling
* API request construction and response handling

## Design Decisions

### Backend as the Source of Truth

Expression evaluation is performed exclusively by the backend.

The frontend builds the expression entered by the user and sends it to the REST API. This avoids maintaining separate mathematical evaluation implementations in both frontend and backend.

This also provides a single source of truth for mathematical behavior.

### Clean Architecture

The backend separates mathematical domain logic from application orchestration and HTTP infrastructure.

This makes the calculator logic independently testable and prevents HTTP concerns from being mixed with mathematical evaluation.

### CORS

The backend enables CORS so the React frontend can communicate with the API during local development. The frontend and backend run on separate localhost ports.

### Expression-Based API

The API accepts a complete expression instead of exposing a separate endpoint for every mathematical operation.

For example:

```json
{
  "expression": "10 + 5 × 2"
}
```

This keeps the API small while allowing the backend to support operator precedence, parentheses, nested expressions, and additional operations without requiring new endpoints.

### Frontend Componentization

The frontend separates the calculator into:

* `Calculator` — state and interaction logic
* `Display` — expression/result presentation
* `Keypad` — calculator controls

API communication is also isolated in a dedicated service.

### Testing Strategy

Frontend tests focus on user behavior and communication with the API rather than duplicating backend mathematical tests.

Backend tests focus on expression evaluation, validation, HTTP behavior, routing, and error handling.

## Assumptions

* The frontend and backend are run locally during development.
* The backend listens on port `8080`.
* The frontend communicates with the backend through the REST API.
* Mathematical expression evaluation is performed using `float64`.
* Percentages use calculator-style contextual behavior. For example, `100 + 10%` evaluates as `110`.
* Invalid mathematical expressions are treated as client errors and return HTTP `400`.
* The backend is considered the authoritative source for calculation results.

## AI Tooling

AI-assisted development tools were used during the implementation of this project.

AI assistance was primarily used for:

* Reviewing architecture and code organization
* Discussing implementation approaches
* Identifying edge cases
* Designing and refining tests
* Reviewing frontend responsiveness and UX
* Troubleshooting TypeScript, Vitest, and testing-library configuration
* Improving documentation

The final implementation was reviewed, tested, and validated manually by the developer.

AI tooling was used as a development aid rather than as a replacement for understanding or verifying the implementation.
