package actions

import (
	"fmt"

	"github.com/Chegnxufeng1994/bdd-in-action/chapter02/banking"
)

// ParseAccountType converts a string to AccountType
// Made public to be reusable across test packages
func ParseAccountType(accountTypeStr string) (banking.AccountType, error) {
	switch accountTypeStr {
	case "Current":
		return banking.AccountTypeCurrent, nil
	case "Savings":
		return banking.AccountTypeSavings, nil
	case "Investment":
		return banking.AccountTypeInvestment, nil
	case "SuperSaver":
		return banking.AccountTypeSuperSaver, nil
	default:
		return "", fmt.Errorf("unknown account type: %s", accountTypeStr)
	}
}
