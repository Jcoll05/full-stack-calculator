# AI Prompts Used

This document contains a representative selection of the prompts used during development. It is not intended to be an exhaustive transcript of the development conversation, but focuses on prompts that materially influenced the architecture, implementation, testing, and refinement of the project.

Note: Some prompts have been condensed or paraphrased for readability. The document focuses on prompts that materially affected implementation decisions rather than routine troubleshooting or minor iterations.

## Architecture and Project Structure

> Review the current calculator project and suggest a maintainable full-stack architecture using React/TypeScript for the frontend and Go for the backend. The frontend should consume the backend API for expression evaluation.

> Given the assignment requirements, how should the backend be structured to keep expression evaluation independent from HTTP concerns and make the core logic easy to unit test?

> Review the proposed frontend and backend folder structures and identify whether the responsibilities are appropriately separated and maintainable for the scope of this assignment.

## Backend Implementation

> Help design an expression evaluator in Go that supports operator precedence, parentheses, exponentiation, square root, percentage, and appropriate validation/errors while keeping the domain logic independent from HTTP.

> Review the backend API design for a calculator service. The API should accept an expression, validate it, evaluate it, and return the result as JSON. What should the request/response structure and HTTP error handling look like?

> Identify appropriate domain-level errors for invalid expressions, division by zero, negative square roots, unexpected operators, invalid numbers, and unmatched parentheses.

> Review the backend implementation and identify whether the HTTP layer, application service, and expression evaluation logic have clearly separated responsibilities.

## Expression Evaluation

> Identify important edge cases for a calculator expression parser, including division by zero, invalid expressions, unmatched parentheses, consecutive operators, negative square roots, and malformed numbers.

> How should operator precedence and parentheses be handled in a calculator expression parser?

> How should implicit multiplication such as `2(3)`, `(2)3`, and `(2+2)2` be handled by the expression parser while preserving operator precedence and avoiding ambiguous invalid expressions?

> Review the percentage behavior of the calculator. It should support both standalone percentages such as `2%` and calculator-style contextual percentages such as `100 + 10%`, `100 - 10%`, `100 × 10%`, and `100 ÷ 10%`. What parsing/evaluation approach would support these cases consistently?

> Review the parser for ambiguous or malformed expressions and suggest regression tests that would prevent valid implicit multiplication from accidentally accepting invalid expressions.

> Verify the evaluation of complex expressions containing nested parentheses, square roots, percentages, implicit multiplication, exponentiation, and mixed operators. Identify any precedence or parsing issues that could produce an incorrect result.

## Frontend Implementation

> Review the calculator component and suggest how to structure the React UI into maintainable components while keeping the calculator state and expression-building logic clear.

> Review the calculator expression-building logic and identify potential problems with numbers, operators, decimals, parentheses, square roots, percentages, backspace, and clear/reset behavior.

> What calculator UI behavior would provide a good user experience for parentheses, particularly preventing users from entering an unmatched closing parenthesis?

> Review the calculator display and responsive layout and suggest improvements for desktop, tablet, and mobile screens while keeping the implementation simple and maintainable.

## Testing

> Review the calculator tests and identify meaningful user-facing behaviors and edge cases that should be covered without writing tests solely to increase code coverage.

> The calculator component is around 90%+ covered. Which remaining branches are worth testing, and which would be artificial coverage-padding tests?

> Review the calculator API service and identify the meaningful success and error paths that should be covered by unit tests, including backend-provided errors and fallback error handling.

> Review the backend tests and identify important parser, service, HTTP handler, routing, and middleware behaviors that should be covered.

> Review the current test suite and identify any missing regression tests for behavior that was recently added to the calculator parser.

## Coverage

> Review the frontend coverage results and determine whether the remaining uncovered statements and branches represent meaningful missing behavior or whether pursuing 100% coverage would result in artificial tests.

> Explain how to document test coverage in the README, including the distinction between meaningful source-code coverage and generated coverage artifacts.

The frontend test suite was intentionally not optimized for 100% coverage. The final suite contains 19 tests across 2 test files and achieves 94.07% statement coverage, 84.00% branch coverage, 100% function coverage, and 94.77% line coverage.

The remaining uncovered component code consists primarily of defensive or low-value branches. These were evaluated rather than automatically tested solely to increase the coverage percentage.

The `calculatorApi.ts` service, which contains a small amount of API-specific logic, was fully covered at 100% across statements, branches, functions, and lines.

## UI/UX and Responsive Design

> Review the calculator interface and suggest a clean, intuitive layout for a full-stack calculator while keeping the UI implementation appropriately scoped for a take-home assignment.

> Identify responsive design considerations for a calculator interface that should work on desktop and mobile without adding unnecessary complexity.

> Review the handling of long calculator expressions and suggest a way to keep the display usable when the expression exceeds the available width.

## Documentation and Review

> Review the README against the assignment requirements and identify anything missing or unclear, including setup instructions, API examples, design decisions, assumptions, testing, coverage, and AI tooling.

> Review the project as if you were evaluating it as a take-home full-stack assignment. Identify issues that could affect correctness, clarity, maintainability, testing, documentation, or overall submission quality.

> Review the final repository structure and identify generated files, unnecessary dependencies, debugging artifacts, or other files that should not be committed.

## Docker and Containerization

> Review the assignment requirements and determine whether the optional Dockerfile should containerize the frontend and backend separately or provide a simple full-stack setup.

> Design a simple Docker setup for the React/Vite frontend and Go backend that is appropriate for a take-home assignment without introducing unnecessary infrastructure or complexity.

> Review the proposed backend Dockerfile and suggest a multi-stage build that compiles the Go application and runs the resulting binary in a lightweight production image.

> Review the proposed frontend Dockerfile and suggest a multi-stage build that installs dependencies, creates the Vite production bundle, and serves the application using Nginx.

> The frontend and backend run in separate Docker containers. How should Docker Compose connect them while allowing the React application to use the same `/api/v1` path in both local development and the containerized production setup?

> Review the Nginx configuration and determine how it should serve the React production build while reverse-proxying `/api/v1/*` requests to the Go backend container through the Docker Compose network.

> Review the Docker Compose configuration for the full-stack calculator and identify whether the service dependencies, port mappings, build arguments, and container networking are appropriate for the project.

> Review the Docker setup for unnecessary files being included in the build context and suggest appropriate `.dockerignore` files for the frontend and backend.

> Troubleshoot the Docker build and runtime configuration when the backend Go version required by `go.mod` differs from the Go version used by the Docker builder image.

> Review the Dockerized application as a take-home assignment and identify whether the containerization approach is simple, reproducible, documented, and appropriate for the scope of the project.

The final Docker setup uses separate multi-stage builds for the frontend and backend and Docker Compose to run them together.

The Go backend is compiled in a Go builder image and executed from a lightweight Alpine runtime image. The React/Vite frontend is built with Node.js and served using Nginx.

Nginx also acts as a reverse proxy for `/api/v1/*` requests, forwarding them to the backend container through the Docker Compose network. This allows the frontend to use a relative API path when containerized instead of depending on a host-specific backend address.

Docker was intentionally kept as a simple optional deployment method rather than introducing additional infrastructure or orchestration complexity.


## Use of AI

AI was used as a development aid throughout the project for architecture discussions, implementation approaches, edge-case analysis, testing strategy, coverage analysis, UI/UX refinement, troubleshooting, and documentation review.

The resulting implementation was reviewed and validated manually, including running the test suites, linting, production builds, API behavior, and end-to-end calculator interactions. AI suggestions were treated as recommendations rather than automatically accepted changes.