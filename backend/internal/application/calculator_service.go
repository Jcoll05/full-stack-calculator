package application

import "github.com/Jcoll05/full-stack-calculator/backend/internal/domain"

// CalculatorService provides methods to evaluate mathematical expressions.
type CalculatorService struct{}

// NewCalculatorService creates a new instance of CalculatorService.
func NewCalculatorService() *CalculatorService {
	return &CalculatorService{}
}

// Calculate evaluates the provided mathematical expression using the domain layer.
func (s *CalculatorService) Calculate(expression string) (float64, error) {
	return domain.EvaluateExpression(expression)
}