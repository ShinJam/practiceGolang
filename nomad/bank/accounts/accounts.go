package accounts

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("can't withdraw you are poor")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}

	a.balance -= amount
	return nil
}

// ChangeOnwer of the account
func (a *Account) ChangeOnwer(newOnwer string) {
	a.owner = newOnwer
}

// Onwer of the account
func (a Account) Onwer() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Onwer(), "'s account.\nHas:", a.Balance())
}
