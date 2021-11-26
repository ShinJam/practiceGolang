package main

import (
	"bank/accounts"
	"fmt"
)

func main() {
	account := accounts.NewAccount("jam")
	account.Deposit(100)
	fmt.Println(account.Balance())
	err := account.Withdraw(110)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Onwer())
	fmt.Println(account)
}
