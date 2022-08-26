package libs

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

var wg sync.WaitGroup

func Concurrency() {
	account := &Account{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
		}()
	}
	wg.Wait()
}

func DepositAndWithdraw(account *Account) {
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Microsecond)
	account.Balance -= 10000
}
