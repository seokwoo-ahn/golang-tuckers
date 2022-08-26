package libs

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

type AccountMutex struct {
	Balance int
}

func DepositAndWithdrawMutex(account *AccountMutex) {
	mutex.Lock()
	defer mutex.Unlock()
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Microsecond)
	account.Balance -= 1000
}

func Mutex() {
	var wg sync.WaitGroup

	account := &AccountMutex{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdrawMutex(account)
			}
		}()
	}
	wg.Wait()
}
