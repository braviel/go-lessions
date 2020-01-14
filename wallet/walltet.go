package wallet

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

//Bcoin unit
type Bcoin int

//Stringer for print format
type Stringer interface {
	String() string
}

func (b Bcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//Wallet impl
type Wallet struct {
	balance Bcoin
}

//Balance of wallet
func (w *Wallet) Balance() Bcoin {
	return w.balance
}

//Deposit to wallet
func (w *Wallet) Deposit(amount Bcoin) {
	w.balance += amount
}

//Withdraw from wallet
func (w *Wallet) Withdraw(amount Bcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
