package something

import "fmt"

// 소문자로 시작하면 private
func bye() {
	fmt.Println("This is private bye function")
}

// 대문로 시작하면 public
func Hello() {
	fmt.Println("This is public Hello function")
}
