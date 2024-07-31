package main

import "fmt"

type PaymentMethod interface {
	Pay(amount float32) string
}

type CreditCard struct {
	Number string
}

func (c CreditCard) Pay(amount float32) string {
	return fmt.Sprint("Paid %.2f using Credit Card ending in %s", amount, c.Number[len(c.Number)-4:])
}

type PayPal struct {
	Email string
}

func (p PayPal) Pay(amount float32) string {
	return fmt.Sprint("Paid %.2f using paypal account %s", amount, p.Email)
}

type BitCoin struct {
	BitcoinAddress string
}

func (b BitCoin) Pay(amount float32) string {
	return fmt.Sprint("Paid %.2f using bitcoin address %s", amount, b.BitcoinAddress)
}

//func to process payment with any payment method

func processPayment(method PaymentMethod, amount float32) string {
	return method.Pay(amount)
}


//use like below.
cc := CreditCard{Number:"12345678912345"}
pp := PayPal{Email: "abc@gmail.com"}
bb := BitCoin{BitcoinAddress: "12345678912345"}

fmt.println(processPayment(cc, 100.0))
fmt.println(processPayment(pp, 24.9))
fmt.println(processPayment(bb, 0.1))



