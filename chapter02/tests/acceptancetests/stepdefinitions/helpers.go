package stepdefinitions

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Chegnxufeng1994/bdd-in-action/chapter02/tests/acceptancetests/actions"
	"github.com/Chegnxufeng1994/bdd-in-action/chapter02/tests/acceptancetests/domain"
	"github.com/cucumber/godog"
)

// parseInitialAccountsTable converts a Godog table to InitialAccount objects
// This follows the Java reference implementation's DataTableType pattern
func parseInitialAccountsTable(
	table *godog.Table,
) ([]domain.InitialAccount, error) {
	accounts := make([]domain.InitialAccount, 0)

	// Skip header row
	for i := 1; i < len(table.Rows); i++ {
		row := table.Rows[i]
		accountTypeStr := row.Cells[0].Value
		balanceStr := row.Cells[1].Value

		balance, err := strconv.ParseFloat(balanceStr, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid balance: %s", balanceStr)
		}

		accountType, err := actions.ParseAccountType(accountTypeStr)
		if err != nil {
			return nil, err
		}

		account := domain.NewInitialAccount(accountType, balance)
		accounts = append(accounts, account)
	}

	return accounts, nil
}

// floatEquals checks if two floats are approximately equal
// Allows for floating point precision errors
func floatEquals(a, b float64) bool {
	const epsilon = 0.01 // Allow 1 cent difference
	return math.Abs(a-b) < epsilon
}
