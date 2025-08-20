package main

import (
	"errors"
	"fmt"
)

// Custom error type
type DivideError struct {
	A, B float64
	Msg  string
}

func (e *DivideError) Error() string {
	return fmt.Sprintf("cannot divide %.2f by %.2f: %s", e.A, e.B, e.Msg)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		// wrap a custom error inside another
		baseErr := &DivideError{A: a, B: b, Msg: "denominator must not be zero"}
		return 0, fmt.Errorf("divide failed: %w", baseErr)
	}
	return a / b, nil
}

func main() {
	data, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)

		// 1. Unwrap: get the base error
		fmt.Println("Unwrapped error:", errors.Unwrap(err))

		// 2. Is: check if error matches a known error
		if errors.Is(err, errors.New("denominator must not be zero")) {
			// ❌ This won't match because a *new* error is not equal
			fmt.Println("Matched plain error (won't print)")
		}

		// 3. As: check if it’s of type *DivideError
		var divErr *DivideError
		if errors.As(err, &divErr) {
			fmt.Println("DivideError detected, values were:", divErr.A, "/", divErr.B)
		}
	}

	fmt.Printf("Result: %.2f\n", data)
}
