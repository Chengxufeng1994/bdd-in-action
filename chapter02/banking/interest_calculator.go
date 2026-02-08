package banking

import (
	"math"
	"time"
)

type InterestCalculator struct {
	interestRates map[AccountType]float64
}

func NewInterestCalculator() *InterestCalculator {
	return &InterestCalculator{
		interestRates: make(map[AccountType]float64),
	}
}

func (ic *InterestCalculator) SetRates(accountType AccountType, interestRate float64) {
	ic.interestRates[accountType] = interestRate
}

func (ic *InterestCalculator) CalculateMonthlyInterestOn(account *BankAccount) Transaction {
	rate := ic.interestRates[account.AccountType()]
	rawInterestEarned := rate * account.Balance() / 100.0 / 12.0
	roundedInterestEarned := math.Round(rawInterestEarned*100.0) / 100.0
	transaction := NewTransaction(time.Now(), "INTEREST", roundedInterestEarned)
	account.RecordTransaction(transaction)
	return transaction
}
