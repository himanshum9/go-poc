package calculator

import "testing"

func TestCalculateDiscount(t *testing.T) {
	type testCase struct {
		minpurchasedAmount int
		discountAmount     int
		purchaseAmount     int
		expected           int
	}
	testCases := []testCase{
		{minpurchasedAmount: 100, discountAmount: 10, purchaseAmount: 110, expected: 100},
		{minpurchasedAmount: 100, discountAmount: 10, purchaseAmount: 50, expected: 50}, {minpurchasedAmount: 100, discountAmount: 10, purchaseAmount: 1000, expected: 900}, {minpurchasedAmount: 100, discountAmount: 10, purchaseAmount: 500, expected: 450}, {minpurchasedAmount: 100, discountAmount: 10, purchaseAmount: 200, expected: 180},
	}
	for _, val := range testCases {
		calculator := NewDiscountCalculator(val.minpurchasedAmount, val.discountAmount)
		amount := calculator.CalculateDiscount(val.purchaseAmount)

		if amount != val.expected {
			t.Errorf("Expected:%v,got:%v", val.expected, amount) // this function is the combination of above 2 functions
		}
	}

}

// func TestDiscountApplied(t *testing.T) {
// 	calculator := NewDiscountCalculator(100, 10)
// 	amount := calculator.CalculateDiscount(110)

// 	if amount != 100 {
// 		t.Fail()
// 	}
// }

// func TestDiscountNotApplied(t *testing.T) {
// 	calculator := NewDiscountCalculator(100, 10)
// 	amount := calculator.CalculateDiscount(50)

// 	if amount != 50 {
// 		// t.Logf("Expected:50,got:%v", amount)
// 		// t.Fail()
// 		t.Errorf("Expected:50,got:%v", amount) // this function is the combination of above 2 functions
// 	}
// }
