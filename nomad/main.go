package main

import (
	"fmt"
	"nomad/something"
)

func main() {
	fmt.Println("hi")

	something.Hello()
	// something.bye() // error: ./main.go:12:2: cannot refer to unexported name something.bye
}
