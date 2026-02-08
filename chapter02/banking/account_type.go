package banking

type AccountType string

const (
	AccountTypeCurrent    AccountType = "current"
	AccountTypeSavings    AccountType = "savings"
	AccountTypeInvestment AccountType = "investment"
	AccountTypeSuperSaver AccountType = "super_saver"
)
