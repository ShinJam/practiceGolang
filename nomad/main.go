package main

import (
	"fmt"
	"nomad/something"
	"strings"
)

func printDivider(word string) {
	repats := strings.Repeat("=", 10)
	fmt.Printf("%v %v %v\n", repats, word, repats)
}

func main() {
	printDivider("#1.1 Packages and Imports")

	something.Hello()
	// something.bye() // error: ./main.go:12:2: cannot refer to unexported name something.bye

	name1 := "jam"
	var name2 = "jam"
	var name3 string = "jam"

	fmt.Println(name1, name2, name3)

}
