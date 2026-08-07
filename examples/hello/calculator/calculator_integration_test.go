//go:build integration

package calculator_test

import (
	"testing"

	"github.com/ffreis/hello/calculator"
)

// TestParseExprIntegration exercises ParseExpr end-to-end across every
// supported operator. It only compiles under the "integration" build tag —
// its presence is what go-integration-coverage.yml's self-test job (see
// .github/workflows/self-test.yml) detects to decide the fleet's
// integration-coverage modality has real tests to gate against, rather than
// silently skipping.
func TestParseExprIntegration(t *testing.T) {
	cases := []struct {
		expr string
		want float64
	}{
		{"3 + 4", 7},
		{"10 - 3", 7},
		{"3 * 4", 12},
		{"10 / 2", 5},
	}

	for _, tc := range cases {
		got, err := calculator.ParseExpr(tc.expr)
		if err != nil {
			t.Fatalf("ParseExpr(%q) returned unexpected error: %v", tc.expr, err)
		}
		if got != tc.want {
			t.Errorf("ParseExpr(%q) = %v, want %v", tc.expr, got, tc.want)
		}
	}
}
