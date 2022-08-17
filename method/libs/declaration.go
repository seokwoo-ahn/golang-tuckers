package libs

import "fmt"

type Account struct {
	Balance int
}

func WithdrawFunc(a *Account, amount int) {
	a.Balance -= amount
	fmt.Println("balance is:", a.Balance)
}

func (a *Account) WithdrawMethod(amount int) {
	a.Balance -= amount
	fmt.Println("balance is:", a.Balance)
}
