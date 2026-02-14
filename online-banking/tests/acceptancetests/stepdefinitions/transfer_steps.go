package stepdefinitions

import (
	"fmt"
	"strconv"

	"github.com/Chegnxufeng1994/bdd-in-action/online-banking/banking"
	"github.com/Chegnxufeng1994/bdd-in-action/online-banking/tests/acceptancetests/actions"
	"github.com/Chegnxufeng1994/bdd-in-action/online-banking/tests/acceptancetests/testcontext"
	"github.com/cucumber/godog"
)

// TransferSteps defines step definitions for money transfer scenarios
// Following the Java reference implementation's TransferStepDefinitions
type TransferSteps struct {
	ctx         *testcontext.TestContext
	transferApi *actions.TransferApi
}

// NewTransferSteps creates a new TransferSteps instance
func NewTransferSteps(ctx *testcontext.TestContext) *TransferSteps {
	return &TransferSteps{
		ctx:         ctx,
		transferApi: actions.NewTransferApi(ctx),
	}
}

// RegisterSteps registers transfer-related step definitions with godog
func (ts *TransferSteps) RegisterSteps(sc *godog.ScenarioContext) {
	// Given steps - Account setup
	sc.Step(`^(\w+) has a (\w+) account with \$(\d+(?:\.\d+)?)$`,
		ts.clientHasAccountWithBalance)
	sc.Step(`^a (\w+) account with \$(\d+(?:\.\d+)?)$`,
		ts.clientHasAdditionalAccount)
	sc.Step(`^(\w+) has the following accounts:$`,
		ts.clientHasFollowingAccounts)

	// When steps - Transfer actions
	sc.Step(`^she transfers \$(\d+(?:\.\d+)?) from the (\w+) account to the (\w+) account$`,
		ts.transferBetweenAccounts)

	// Then steps - Transfer assertions
	sc.Step(`^she should have \$(\d+(?:\.\d+)?) in her (\w+) account$`,
		ts.shouldHaveBalanceInAccount)
	sc.Step(`^she should receive an 'insufficient funds' error$`,
		ts.shouldReceiveInsufficientFundsError)
	sc.Step(`^her accounts should look like this:$`,
		ts.accountsShouldLookLike)
}

// Given steps - Setup initial state

func (ts *TransferSteps) clientHasAccountWithBalance(
	clientName, accountTypeStr string,
	balance float64,
) error {
	// Create client if not exists
	if ts.ctx.Client == nil {
		ts.ctx.Client = banking.NewClient(clientName)
	}

	// Parse account type
	accountType, err := actions.ParseAccountType(accountTypeStr)
	if err != nil {
		return err
	}

	// Create and open account
	account := banking.BankAccountOfType(accountType).WithBalance(balance)
	ts.ctx.Client.Opens(account)
	ts.ctx.RegisterAccount(account)

	return nil
}

func (ts *TransferSteps) clientHasAdditionalAccount(
	accountTypeStr string,
	balance float64,
) error {
	if ts.ctx.Client == nil {
		return fmt.Errorf("client must be created first")
	}

	// Parse account type
	accountType, err := actions.ParseAccountType(accountTypeStr)
	if err != nil {
		return err
	}

	// Create and open account
	account := banking.BankAccountOfType(accountType).WithBalance(balance)
	ts.ctx.Client.Opens(account)
	ts.ctx.RegisterAccount(account)

	return nil
}

func (ts *TransferSteps) clientHasFollowingAccounts(
	clientName string,
	table *godog.Table,
) error {
	// Create client if not exists
	if ts.ctx.Client == nil {
		ts.ctx.Client = banking.NewClient(clientName)
	}

	// Parse table data
	accounts, err := parseInitialAccountsTable(table)
	if err != nil {
		return err
	}

	// Open all accounts
	for _, acc := range accounts {
		account := banking.BankAccountOfType(acc.AccountType).WithBalance(acc.Balance)
		ts.ctx.Client.Opens(account)
		ts.ctx.RegisterAccount(account)
	}

	return nil
}

// When steps - Actions being tested

func (ts *TransferSteps) transferBetweenAccounts(
	amount float64,
	fromTypeStr, toTypeStr string,
) error {
	// Parse account types
	fromType, err := actions.ParseAccountType(fromTypeStr)
	if err != nil {
		return err
	}

	toType, err := actions.ParseAccountType(toTypeStr)
	if err != nil {
		return err
	}

	// Use TransferApi for fluent transfer
	return ts.transferApi.TransferBetweenAccountTypes(amount, fromType, toType)
}

// Then steps - Assertions

func (ts *TransferSteps) shouldHaveBalanceInAccount(
	expectedBalance float64,
	accountTypeStr string,
) error {
	// Parse account type
	accountType, err := actions.ParseAccountType(accountTypeStr)
	if err != nil {
		return err
	}

	// Get account balance
	account := ts.ctx.GetAccount(accountType)
	if account == nil {
		return fmt.Errorf("account %s not found", accountTypeStr)
	}

	actualBalance := account.Balance()
	if !floatEquals(actualBalance, expectedBalance) {
		return fmt.Errorf(
			"expected balance $%.2f in %s account, but got $%.2f",
			expectedBalance,
			accountTypeStr,
			actualBalance,
		)
	}

	return nil
}

func (ts *TransferSteps) shouldReceiveInsufficientFundsError() error {
	if ts.ctx.LastError == nil {
		return fmt.Errorf("expected 'insufficient funds' error, but got none")
	}

	if ts.ctx.LastError.Error() != "insufficient funds" {
		return fmt.Errorf(
			"expected 'insufficient funds' error, but got: %v",
			ts.ctx.LastError,
		)
	}

	return nil
}

func (ts *TransferSteps) accountsShouldLookLike(table *godog.Table) error {
	// Skip header row
	for i := 1; i < len(table.Rows); i++ {
		row := table.Rows[i]
		accountTypeStr := row.Cells[0].Value
		expectedBalanceStr := row.Cells[1].Value

		expectedBalance, err := strconv.ParseFloat(expectedBalanceStr, 64)
		if err != nil {
			return fmt.Errorf("invalid expected balance: %s", expectedBalanceStr)
		}

		// Parse account type
		accountType, err := actions.ParseAccountType(accountTypeStr)
		if err != nil {
			return err
		}

		// Get account balance
		account := ts.ctx.GetAccount(accountType)
		if account == nil {
			return fmt.Errorf("account %s not found", accountTypeStr)
		}

		actualBalance := account.Balance()
		if !floatEquals(actualBalance, expectedBalance) {
			return fmt.Errorf(
				"expected balance $%.2f in %s account, but got $%.2f",
				expectedBalance,
				accountTypeStr,
				actualBalance,
			)
		}
	}

	return nil
}
