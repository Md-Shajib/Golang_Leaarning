package payment

type PaymentGateway interface {
	Charge(amount int64, currency string) (Transaction, error)
}

type Transaction interface {
	ID() string
	Amount() int64
	Currency() string
	Status() string
}

type Logger interface {
	Info(msg string)
	Error(msg string)
}

type PaymentRepository interface {
	Save(transaction Transaction) error
}

type Dependencies interface {
	Logger
	PaymentRepository
}