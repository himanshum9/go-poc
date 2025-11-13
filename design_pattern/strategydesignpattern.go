package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64)
}

type Paypal struct {
	Email string
}

func (p Paypal) Pay(amount float64) {
	fmt.Printf("Paying $%.2f using PayPal account %s\n", amount, p.Email)
}

type CreditCard struct {
	CardNumber string
}

func (c CreditCard) Pay(amount float64) {
	fmt.Printf("Paying $%.2f using Credit Card %s\n", amount, c.CardNumber)
}

type Checkout struct {
	strategy PaymentStrategy
}

func (c *Checkout) SetStrategy(s PaymentStrategy) {
	c.strategy = s
}

func (c *Checkout) Process(amount float64) {
	if c.strategy == nil {
		fmt.Println("no payment method found")
		return
	}
	c.strategy.Pay(amount)
	fmt.Println("Payment successfull")
}
