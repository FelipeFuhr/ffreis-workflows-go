// Package calculator provides basic arithmetic and expression parsing.
package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

// Add returns the sum of a and b.
func Add(a, b float64) float64 { return a + b }

// Sub returns the difference of a and b.
func Sub(a, b float64) float64 { return a - b }

// Mul returns the product of a and b.
func Mul(a, b float64) float64 { return a * b }

// Div returns a divided by b. Returns an error if b is zero.
func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// ParseExpr parses a simple "a op b" expression (e.g. "3 + 4") and returns the result.
// Supported operators: +, -, *, /.
func ParseExpr(expr string) (float64, error) {
	parts := strings.Fields(expr)
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid expression %q: expected \"a op b\"", expr)
	}

	a, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid operand %q: %w", parts[0], err)
	}

	b, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid operand %q: %w", parts[2], err)
	}

	switch parts[1] {
	case "+":
		return Add(a, b), nil
	case "-":
		return Sub(a, b), nil
	case "*":
		return Mul(a, b), nil
	case "/":
		return Div(a, b)
	default:
		return 0, fmt.Errorf("unknown operator %q", parts[1])
	}
}
