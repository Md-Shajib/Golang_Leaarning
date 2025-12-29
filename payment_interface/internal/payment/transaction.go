package payment

type BasicTransaction struct {
	id			string
	amount 		int64
	currency 	string
	status		string
}

func NewTransaction(id string, amount int64, currency, status string) *BasicTransaction {
	return &BasicTransaction{
		id: id,
		amount: amount,
		currency: currency,
		status: status,
	}
}

func (t *BasicTransaction) ID() string { return t.id}
func (t *BasicTransaction) Amount() int64 { return t.amount}
func (t *BasicTransaction) Currency() string { return t.currency}
func (t *BasicTransaction) Status() string { return t.status}