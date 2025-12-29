package repository

import "go-basic/payment_interface/internal/payment"

type InMemoryPaymentRepository struct {
	data []payment.Transaction
}

func NewInMemoryPaymentRepository() *InMemoryPaymentRepository {
	return &InMemoryPaymentRepository{
		data: make([]payment.Transaction, 0),
	}
}

func (r *InMemoryPaymentRepository) Save(tx payment.Transaction) error {
	r.data = append(r.data, tx)
	return nil
}

func (r *InMemoryPaymentRepository) FindAll() []payment.Transaction {
	return r.data
}
