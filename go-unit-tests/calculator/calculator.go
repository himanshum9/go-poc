package calculator

type DiscountCalculator struct {
	minpurchasedAmount int
	discountAmount     int
}

func NewDiscountCalculator(minpurchasedAmount int, discountAmount int) *DiscountCalculator {
	return &DiscountCalculator{minpurchasedAmount: minpurchasedAmount, discountAmount: discountAmount}
}

func (dc *DiscountCalculator) CalculateDiscount(purchasedAmount int) int {
	if dc.minpurchasedAmount < purchasedAmount {

		multiplier := purchasedAmount / dc.minpurchasedAmount
		return purchasedAmount - (dc.discountAmount * multiplier)
	}
	return purchasedAmount
}
