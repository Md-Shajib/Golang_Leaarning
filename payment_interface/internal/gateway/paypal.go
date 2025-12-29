package gateway

import "go-basic/payment_interface/internal/payment"

type PaypalGateway struct {}

func (p *PaypalGateway) Charge(amount int64, currency string)(payment.Transaction, error){
	return payment.NewTransaction(
		"paypal_tx_123",
		amount,
		currency,
		"success",
	), nil
}