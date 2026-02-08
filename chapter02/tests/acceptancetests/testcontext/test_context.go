package testcontext

import (
	"github.com/Chegnxufeng1994/bdd-in-action/chapter02/banking"
)

// TestContext holds the test state during scenario execution
// Following BDD best practices, this context maintains state between steps
type TestContext struct {
	// Client represents the current test client
	Client *banking.Client

	// InterestCalculator is used for interest calculations
	InterestCalculator *banking.InterestCalculator

	// LastError stores any error from the last operation
	LastError error

	// LastTransaction stores the most recent transaction
	LastTransaction banking.Transaction

	// Accounts is a registry of all accounts created during the test
	Accounts map[banking.AccountType]*banking.BankAccount
}

// NewTestContext creates a new test context
func NewTestContext() *TestContext {
	return &TestContext{
		InterestCalculator: banking.NewInterestCalculator(),
		Accounts:           make(map[banking.AccountType]*banking.BankAccount),
	}
}

// Reset clears the context for the next scenario
func (tc *TestContext) Reset() {
	tc.Client = nil
	tc.InterestCalculator = banking.NewInterestCalculator()
	tc.LastError = nil
	tc.LastTransaction = banking.Transaction{}
	tc.Accounts = make(map[banking.AccountType]*banking.BankAccount)
}

// GetAccount retrieves an account by type from the client
func (tc *TestContext) GetAccount(accountType banking.AccountType) *banking.BankAccount {
	if tc.Client == nil {
		return nil
	}
	return tc.Client.Get(accountType)
}

// RegisterAccount stores an account in the accounts registry
func (tc *TestContext) RegisterAccount(account *banking.BankAccount) {
	tc.Accounts[account.AccountType()] = account
}
