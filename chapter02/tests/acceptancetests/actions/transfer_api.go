package actions

import (
	"errors"
	"fmt"

	"github.com/Chegnxufeng1994/bdd-in-action/chapter02/banking"
	"github.com/Chegnxufeng1994/bdd-in-action/chapter02/tests/acceptancetests/testcontext"
)

// TransferApi provides a fluent API for transferring money between accounts
// Following the Java reference implementation's fluent interface pattern
type TransferApi struct {
	ctx         *testcontext.TestContext
	amount      float64
	fromAccount *banking.BankAccount
	toAccount   *banking.BankAccount
	amountSet   bool
	fromSet     bool
}

// NewTransferApi creates a new TransferApi instance
func NewTransferApi(ctx *testcontext.TestContext) *TransferApi {
	return &TransferApi{
		ctx: ctx,
	}
}

// TheAmount sets the transfer amount and returns self for chaining
func (ta *TransferApi) TheAmount(amount float64) *TransferApi {
	ta.amount = amount
	ta.amountSet = true
	return ta
}

// From sets the source account and returns self for chaining
func (ta *TransferApi) From(account *banking.BankAccount) *TransferApi {
	ta.fromAccount = account
	ta.fromSet = true
	return ta
}

// To sets the destination account and executes the transfer
// This completes the fluent chain and performs the actual operation
func (ta *TransferApi) To(account *banking.BankAccount) error {
	ta.toAccount = account

	// Validate all required parameters are set
	if !ta.amountSet {
		return errors.New("transfer amount not set")
	}
	if !ta.fromSet {
		return errors.New("source account not set")
	}
	if ta.toAccount == nil {
		return errors.New("destination account not set")
	}

	// Check for sufficient funds
	if ta.fromAccount.Balance() < ta.amount {
		ta.ctx.LastError = errors.New("insufficient funds")
		return nil // Don't return error, store it in context
	}

	// Execute the transfer
	ta.fromAccount.Withdraw(ta.amount)
	ta.toAccount.Deposit(ta.amount)

	// Clear error on successful transfer
	ta.ctx.LastError = nil

	return nil
}

// Transfer is a convenience method that performs a complete transfer in one call
func (ta *TransferApi) Transfer(amount float64, from, to *banking.BankAccount) error {
	return ta.TheAmount(amount).From(from).To(to)
}

// TransferBetweenAccountTypes transfers between accounts identified by type
func (ta *TransferApi) TransferBetweenAccountTypes(
	amount float64,
	fromType, toType banking.AccountType,
) error {
	if ta.ctx.Client == nil {
		return fmt.Errorf("no client available")
	}

	fromAccount := ta.ctx.Client.Get(fromType)
	toAccount := ta.ctx.Client.Get(toType)

	if fromAccount == nil {
		return fmt.Errorf("source account %s not found", fromType)
	}

	if toAccount == nil {
		return fmt.Errorf("destination account %s not found", toType)
	}

	return ta.Transfer(amount, fromAccount, toAccount)
}
