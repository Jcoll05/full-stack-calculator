package domain

import (
	"errors"
	"testing"
)

// TestCalculate tests the Calculate function for various operations and edge cases.
func TestCalculate(t *testing.T) {
	tests := []struct {
		name      string
		operation Operation
		a         float64
		b         float64
		want      float64
		wantErr   error
	}{
		{
			name:      "addition",
			operation: Add,
			a:         10,
			b:         5,
			want:      15,
			wantErr:   nil,
		},
		{
			name:      "subtraction",
			operation: Subtract,
			a:         10,
			b:         5,
			want:      5,
			wantErr:   nil,
		},
		{
			name:      "multiplication",
			operation: Multiply,
			a:         10,
			b:         5,
			want:      50,
			wantErr:   nil,
		},
		{
			name:      "division",
			operation: Divide,
			a:         10,
			b:         5,
			want:      2,
			wantErr:   nil,
		},
		{
			name:      "power",
			operation: Power,
			a:         2,
			b:         3,
			want:      8,
			wantErr:   nil,
		},
		{
			name:      "square root",
			operation: SquareRoot,
			a:         25,
			want:      5,
			wantErr:   nil,
		},
		{
			name:      "percentage",
			operation: Percentage,
			a:         25,
			want:      0.25,
			wantErr:   nil,
		},
		{
			name:      "division by zero",
			operation: Divide,
			a:         10,
			b:         0,
			want:      0,
			wantErr:   ErrDivisionByZero,
		},
		{
			name:      "negative square root",
			operation: SquareRoot,
			a:         -25,
			want:      0,
			wantErr:   ErrNegativeSquareRoot,
		},
		{
			name:      "unsupported operation",
			operation: Operation("invalid"),
			a:         10,
			b:         5,
			want:      0,
			wantErr:   ErrUnsupportedOperation,
		},
	}

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.operation, tt.a, tt.b)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
