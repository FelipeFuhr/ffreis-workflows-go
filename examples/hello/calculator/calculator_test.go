package calculator_test

import (
	"testing"

	"github.com/ffreis/hello/calculator"
)

func TestAdd(t *testing.T) {
	if got := calculator.Add(2, 3); got != 5 {
		t.Fatalf("Add(2,3) = %v, want 5", got)
	}
}

func TestSub(t *testing.T) {
	if got := calculator.Sub(5, 3); got != 2 {
		t.Fatalf("Sub(5,3) = %v, want 2", got)
	}
}

func TestMul(t *testing.T) {
	if got := calculator.Mul(3, 4); got != 12 {
		t.Fatalf("Mul(3,4) = %v, want 12", got)
	}
}

func TestDivOk(t *testing.T) {
	got, err := calculator.Div(10, 2)
	if err != nil || got != 5 {
		t.Fatalf("Div(10,2) = %v, %v; want 5, nil", got, err)
	}
}

func TestDivByZero(t *testing.T) {
	_, err := calculator.Div(1, 0)
	if err == nil {
		t.Fatal("Div(1,0): expected error, got nil")
	}
}

func TestParseExpr(t *testing.T) {
	cases := []struct {
		expr string
		want float64
		err  bool
	}{
		{"3 + 4", 7, false},
		{"10 - 3", 7, false},
		{"3 * 4", 12, false},
		{"10 / 2", 5, false},
		{"1 / 0", 0, true},
		{"bad", 0, true},
		{"a + b", 0, true},
		{"1 % 2", 0, true},
	}

	for _, tc := range cases {
		got, err := calculator.ParseExpr(tc.expr)
		if tc.err && err == nil {
			t.Errorf("ParseExpr(%q): want error, got %v", tc.expr, got)
		}
		if !tc.err && err != nil {
			t.Errorf("ParseExpr(%q): unexpected error: %v", tc.expr, err)
		}
		if !tc.err && got != tc.want {
			t.Errorf("ParseExpr(%q) = %v, want %v", tc.expr, got, tc.want)
		}
	}
}

func FuzzParseExpr(f *testing.F) {
	f.Add("3 + 4")
	f.Add("10 / 0")
	f.Add("bad input here")
	f.Fuzz(func(t *testing.T, expr string) {
		// Must not panic regardless of input.
		_, _ = calculator.ParseExpr(expr)
	})
}
