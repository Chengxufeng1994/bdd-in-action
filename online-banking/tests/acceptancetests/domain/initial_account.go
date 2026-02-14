package domain

import "github.com/Chegnxufeng1994/bdd-in-action/online-banking/banking"

// InitialAccount represents an initial account configuration for testing
type InitialAccount struct {
	AccountType banking.AccountType
	Balance     float64
}

// NewInitialAccount creates a new InitialAccount
func NewInitialAccount(accountType banking.AccountType, balance float64) InitialAccount {
	return InitialAccount{
		AccountType: accountType,
		Balance:     balance,
	}
}
