package application

import (
	"errors"
	"testing"

	"github.com/Jcoll05/full-stack-calculator/backend/internal/domain"
)

// TestCalculatorService_Calculate tests the Calculate method of the CalculatorService.
func TestCalculatorService_Calculate(t *testing.T) {
	service := NewCalculatorService()

	tests := []struct {
		name      string
		operation domain.Operation
		a         float64
		b         float64
		want      float64
		wantErr   error
	}{
		{
			name:      "addition",
			operation: domain.Add,
			a:         10,
			b:         5,
			want:      15,
		},
		{
			name:      "division",
			operation: domain.Divide,
			a:         10,
			b:         2,
			want:      5,
		},
		{
			name:      "division by zero",
			operation: domain.Divide,
			a:         10,
			b:         0,
			wantErr:   domain.ErrDivisionByZero,
		},
	}

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.Calculate(tt.operation, tt.a, tt.b)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}