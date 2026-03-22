// Package main is a simple calculator CLI that evaluates arithmetic expressions.
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ffreis/hello/calculator"
)

func main() {
	expr := strings.Join(os.Args[1:], " ")
	if expr == "" {
		fmt.Fprintln(os.Stderr, "usage: hello <expression>  e.g. hello 3 + 4")
		os.Exit(1)
	}

	result, err := calculator.ParseExpr(expr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
