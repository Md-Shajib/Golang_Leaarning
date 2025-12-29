package gateway

import (
	"errors"
	"go-basic/payment_interface/internal/payment"
)

type StripeGateway struct {}

func (s *StripeGateway) Charge(amount int64, currency string)(payment.Transaction, error){
	if amount <= 0 {
		return  nil, errors.New("Invalid amount")
	}
	return  payment.NewTransaction(
		"stripe_tx_123",
		amount,
		currency,
		"success",
	), nil
}

func (s *StripeGateway) Refund(transactionId string) error {
	return nil
}