package example

import (
	"fmt"
	"sync"
)

type Bank struct {
	balance int64
	sync.Mutex
}

func (b *Bank) Deposit(amount int64) {
	b.Lock()
	b.balance += amount
	b.Unlock()
}

func (b *Bank) Withdraw(amount int64) error {
	b.Lock()
	defer b.Unlock()

	newBalance := b.balance - amount
	if newBalance < 0 {
		return fmt.Errorf("funds insufficient")
	}

	b.balance = newBalance
	return nil
}
