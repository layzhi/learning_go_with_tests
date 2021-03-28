package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFunds = errors.New("oh no, insufficient funds")

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}