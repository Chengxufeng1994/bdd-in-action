package banking

type BankAccount struct {
	accountType AccountType
	balance     float64
}

func NewBankAccount(accountType AccountType) *BankAccount {
	return &BankAccount{
		accountType: accountType,
		balance:     0,
	}
}

func BankAccountOfType(accountType AccountType) *BankAccount {
	return NewBankAccount(accountType)
}

func (b *BankAccount) WithBalance(balance float64) *BankAccount {
	b.balance = balance
	return b
}

func (b *BankAccount) Deposit(amount float64) {
	b.balance += amount
}

func (b *BankAccount) Withdraw(amount float64) {
	b.balance -= amount
}

func (b *BankAccount) Balance() float64 {
	return b.balance
}

func (b *BankAccount) AccountType() AccountType {
	return b.accountType
}

func (b *BankAccount) RecordTransaction(transaction Transaction) {
	b.balance = b.balance + transaction.Amount()
}
