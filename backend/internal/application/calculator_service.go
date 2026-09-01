package application

import "github.com/Jcoll05/full-stack-calculator/backend/internal/domain"

// CalculatorService provides methods to perform arithmetic operations.
type CalculatorService struct{}

// NewCalculatorService creates a new instance of CalculatorService.
func NewCalculatorService() *CalculatorService {
	return &CalculatorService{}
}

// Calculate performs the specified arithmetic operation using the domain's Calculate function.
func (s *CalculatorService) Calculate(
	operation domain.Operation,
	a, b float64,
) (float64, error) {
	return domain.Calculate(operation, a, b)
}