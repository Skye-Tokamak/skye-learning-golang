package pointer

import (
	"errors"
	"fmt"
)

var InsufficientFundsError = errors.New("money not enough")

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
func (w *Wallet) Deposit(income Bitcoin) {
	w.balance += income
}
func (w *Wallet) Withdraw(takeout Bitcoin) error {
	if w.balance-takeout > 0 {
		w.balance -= takeout
	} else {
		return InsufficientFundsError
	}
	return nil
}
