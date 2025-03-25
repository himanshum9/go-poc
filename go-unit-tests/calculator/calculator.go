package calculator

import "errors"

type DiscountCalculator struct {
	minpurchasedAmount int
	discountAmount     int
}

func NewDiscountCalculator(minpurchasedAmount int, discountAmount int) (*DiscountCalculator, error) {
	if minpurchasedAmount == 0 {
		return &DiscountCalculator{}, errors.New("0 value not allowed")
	}
	return &DiscountCalculator{minpurchasedAmount: minpurchasedAmount, discountAmount: discountAmount}, nil
}

func (dc *DiscountCalculator) CalculateDiscount(purchasedAmount int) int {
	if dc.minpurchasedAmount < purchasedAmount {

		multiplier := purchasedAmount / dc.minpurchasedAmount
		return purchasedAmount - (dc.discountAmount * multiplier)
	}
	return purchasedAmount
}
