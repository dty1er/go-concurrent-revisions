package example

import (
	"fmt"
	"sync"
)

type Account struct {
	balance int64
	sync.Mutex
}

func (a *Account) Deposit(amount int64) {
	a.Lock()
	a.balance += amount
	a.Unlock()
}

func (a *Account) Withdraw(amount int64) error {
	a.Lock()
	defer a.Unlock()

	newBalance := a.balance - amount
	if newBalance < 0 {
		return fmt.Errorf("funds insufficient")
	}

	a.balance = newBalance
	return nil
}
