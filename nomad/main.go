package main

import (
	"fmt"
	"nomad/something"
)

func main() {
	fmt.Println("hi")

	something.Hello()
	// something.bye() // error: ./main.go:12:2: cannot refer to unexported name something.bye

	name1 := "jam"
	var name2 = "jam"
	var name3 string = "jam"

	fmt.Println(name1, name2, name3)

}
