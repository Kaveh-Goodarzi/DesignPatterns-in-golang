package main

import "fmt"

type OrderVerification struct {}

func (ov *OrderVerification) VerifyOrder(orderID int) bool {
	return true
}

type PaymentProcessor struct {}

func (pp *PaymentProcessor) ProcessPayment(amount float64) bool {
	return true
}

type OrderFullfillment struct {}

func (off *OrderFullfillment) FullfillOrder(orderID int) bool {
	return true
}

type OrderFacade struct {
	verifier       *OrderVerification
	processor        *PaymentProcessor
	fullfillment   *OrderFullfillment
}

func (od *OrderFacade) Constructor(orderID int, amount float64) {
	od.verifier = &OrderVerification{}
	od.processor = &PaymentProcessor{}
	od.fullfillment = &OrderFullfillment{}
}

func (od *OrderFacade) PlaceOrder(orderID int, amount float64) bool {
	verified := od.verifier.VerifyOrder(orderID)

	if !verified {
		return false
	}

	paid := od.processor.ProcessPayment(amount)

	if !paid {
		return false
	}
	
	fulfilled := od.fullfillment.FullfillOrder(orderID)
	
	if !fulfilled {
		return false
	}

	return true
}

func main() {
	facade := &OrderFacade{}
	facade.PlaceOrder(123, 100000.00)

	if facade.PlaceOrder(123, 100000.00) {
		fmt.Println("order placed successfully")
	} else {
		fmt.Println("failed to process order")
	}
	// Output: true
}