package mocks

import "go-basic/payment_interface/internal/payment"

type MockGateway struct{}

func (m *MockGateway) Charge (amount int64, currency string) (payment.Transaction, error) {
	return payment.NewTransaction(
		"mock_tx_123",
		amount,
		currency,
		"success",
	), nil
}

func (m *MockGateway) Refund(transactionId string) error {
	return nil
}
