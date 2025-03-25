package calculator

import "testing"

func TestDiscountApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 10)
	amount := calculator.CalculateDiscount(110)

	if 100 != amount {
		t.Fail()
	}
}
