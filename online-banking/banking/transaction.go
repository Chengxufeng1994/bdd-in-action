package banking

import "time"

type Transaction struct {
	time        time.Time
	description string
	amount      float64
}

func (t Transaction) Time() time.Time {
	return t.time
}

func (t Transaction) Description() string {
	return t.description
}

func (t Transaction) Amount() float64 {
	return t.amount
}

func NewTransaction(time time.Time, description string, amount float64) Transaction {
	return Transaction{
		time:        time,
		description: description,
		amount:      amount,
	}
}
