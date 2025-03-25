package calculator

import "testing"

func TestCalculateDiscount(t *testing.T) {
	type testCase struct {
		name               string
		minpurchasedAmount int
		discountAmount     int
		purchaseAmount     int
		expected           int
	}
	testCases := []testCase{
		{name: "Discount of 10", minpurchasedAmount: 100, discountAmount: 10, purchaseAmount: 110, expected: 100},
		{name: "Discount of 0", minpurchasedAmount: 100, discountAmount: 10, purchaseAmount: 50, expected: 50}, {name: "Discount of 100", minpurchasedAmount: 100, discountAmount: 10, purchaseAmount: 1000, expected: 900}, {name: "Discount of 50", minpurchasedAmount: 100, discountAmount: 10, purchaseAmount: 500, expected: 450}, {name: "Discount of 20", minpurchasedAmount: 100, discountAmount: 10, purchaseAmount: 200, expected: 180},
	}

	for _, val := range testCases {
		t.Run(val.name, func(t *testing.T) {
			calculator, _ := NewDiscountCalculator(val.minpurchasedAmount, val.discountAmount)
			amount := calculator.CalculateDiscount(val.purchaseAmount)

			if amount != val.expected {
				t.Errorf("Expected:%v,got:%v", val.expected, amount) // this function is the combination of above 2 functions
			}
		})

	}

}

func TestCalculateDiscountWithZeroMinimumValue(t *testing.T) {

	_, err := NewDiscountCalculator(0, 10)
	if err == nil {
		t.Fatalf("Should not initiate the calculation")
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
