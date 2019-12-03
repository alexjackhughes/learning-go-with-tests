package wallet

import (
	"errors"
	"fmt"
)

// Bitcoin is an int.
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet is an object that allows
// you to store money.
type Wallet struct {
	balance Bitcoin
}

// Balance shows you how much you have in your wallet.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Deposit allows you to add money.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw allows you to take money.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("that amount is larger than your balance")
	}

	w.balance -= amount
	return nil
}
