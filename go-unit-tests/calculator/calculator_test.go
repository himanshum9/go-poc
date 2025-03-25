package calculator

import "testing"

func TestDiscountApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 10)
	amount := calculator.CalculateDiscount(110)

	if amount != 100 {
		t.Fail()
	}
}

func TestDiscountNotApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 10)
	amount := calculator.CalculateDiscount(50)

	if amount != 50 {
		t.Fail()
	}
}
