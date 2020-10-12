package payment

type CreditAccount struct {
	accountNumber string
	accountOwner  string
}

func (c CreditAccount) AccountNumber() string {
	return c.accountNumber
}
func (c CreditAccount) AccountOwner() string {
	return c.accountOwner
}
func (c CreditAccount) AvailableCredit() float32 {
	return 1000
}
