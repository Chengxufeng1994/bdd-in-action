package banking

type Client struct {
	name     string
	accounts map[AccountType]*BankAccount
}

func NewClient(name string) *Client {
	return &Client{
		name:     name,
		accounts: make(map[AccountType]*BankAccount),
	}
}

func (c *Client) Name() string {
	return c.name
}

func (c *Client) Opens(bankAccount *BankAccount) {
	c.accounts[bankAccount.AccountType()] = bankAccount
}

func (c *Client) Get(accountType AccountType) *BankAccount {
	return c.accounts[accountType]
}

func (c *Client) Accounts() []*BankAccount {
	result := make([]*BankAccount, 0, len(c.accounts))
	for _, account := range c.accounts {
		result = append(result, account)
	}
	return result
}
