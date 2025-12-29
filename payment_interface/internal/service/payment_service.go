package service

import "go-basic/payment_interface/internal/payment"

type PaymentService struct {
	gateway payment.PaymentGateway
	deps    payment.Dependencies
}

func NewPaymentService(gateway payment.PaymentGateway, deps payment.Dependencies) *PaymentService {
	return &PaymentService{
		gateway: gateway,
		deps:    deps,
	}
}

func (s *PaymentService) ProcessPayment(amount int64, currency string) (payment.Transaction, error) {
	s.deps.Info("Starting payment processing")

	tx, err := s.gateway.Charge(amount, currency)
	if err != nil {
		s.deps.Error("Payment failed: " + err.Error())
		return nil, err
	}

	if err := s.deps.Save(tx); err != nil {
		s.deps.Error("Failed to save transaction: " + err.Error())
		return nil, err
	}

	s.deps.Info("Payment processed successfully: " + tx.ID())
	return tx, nil
}
