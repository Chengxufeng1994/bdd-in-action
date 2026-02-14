package stepdefinitions

import (
	"fmt"

	"github.com/Chegnxufeng1994/bdd-in-action/online-banking/banking"
	"github.com/Chegnxufeng1994/bdd-in-action/online-banking/tests/acceptancetests/actions"
	"github.com/Chegnxufeng1994/bdd-in-action/online-banking/tests/acceptancetests/testcontext"
	"github.com/cucumber/godog"
)

// InterestSteps defines step definitions for interest calculation scenarios
type InterestSteps struct {
	ctx *testcontext.TestContext
}

// NewInterestSteps creates a new InterestSteps instance
func NewInterestSteps(ctx *testcontext.TestContext) *InterestSteps {
	return &InterestSteps{
		ctx: ctx,
	}
}

// RegisterSteps registers interest-related step definitions with godog
func (is *InterestSteps) RegisterSteps(sc *godog.ScenarioContext) {
	// Given steps - Interest setup
	sc.Step(`^(\w+) has a (\w+) account with \$(\d+(?:\.\d+)?)$`,
		is.clientHasAccountWithBalance)
	sc.Step(`^the interest rate for (\w+) accounts is (\d+(?:\.\d+)?)$`,
		is.interestRateIsSet)

	// When steps - Interest calculation
	sc.Step(`^the monthly interest is calculated$`,
		is.monthlyInterestIsCalculated)

	// Then steps - Interest assertions
	sc.Step(`^she should have earned \$(\d+(?:\.\d+)?)$`,
		is.shouldHaveEarned)
	sc.Step(`^she should have \$(\d+(?:\.\d+)?) in her (\w+) account$`,
		is.shouldHaveBalanceInAccount)
}

// Given steps - Setup initial state

func (is *InterestSteps) clientHasAccountWithBalance(
	clientName, accountTypeStr string,
	balance float64,
) error {
	// Create client if not exists
	if is.ctx.Client == nil {
		is.ctx.Client = banking.NewClient(clientName)
	}

	// Parse account type
	accountType, err := actions.ParseAccountType(accountTypeStr)
	if err != nil {
		return err
	}

	// Create and open account
	account := banking.BankAccountOfType(accountType).WithBalance(balance)
	is.ctx.Client.Opens(account)
	is.ctx.RegisterAccount(account)

	return nil
}

func (is *InterestSteps) interestRateIsSet(accountTypeStr string, rate float64) error {
	// Parse account type
	accountType, err := actions.ParseAccountType(accountTypeStr)
	if err != nil {
		return err
	}

	// Set interest rate using the interest calculator
	is.ctx.InterestCalculator.SetRates(accountType, rate)
	return nil
}

// When steps - Actions being tested

func (is *InterestSteps) monthlyInterestIsCalculated() error {
	if is.ctx.Client == nil {
		return fmt.Errorf("no client available")
	}

	// Calculate interest for all accounts
	for _, account := range is.ctx.Client.Accounts() {
		transaction := is.ctx.InterestCalculator.CalculateMonthlyInterestOn(account)
		is.ctx.LastTransaction = transaction
	}

	return nil
}

// Then steps - Assertions

func (is *InterestSteps) shouldHaveEarned(expectedEarnings float64) error {
	actualEarnings := is.ctx.LastTransaction.Amount()
	if !floatEquals(actualEarnings, expectedEarnings) {
		return fmt.Errorf(
			"expected earnings $%.2f, but got $%.2f",
			expectedEarnings,
			actualEarnings,
		)
	}
	return nil
}

func (is *InterestSteps) shouldHaveBalanceInAccount(
	expectedBalance float64,
	accountTypeStr string,
) error {
	// Parse account type
	accountType, err := actions.ParseAccountType(accountTypeStr)
	if err != nil {
		return err
	}

	// Get account balance
	account := is.ctx.GetAccount(accountType)
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
