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

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	printDivider("#1.1 Packages and Imports")

	something.Hello()
	// something.bye() // error: ./main.go:12:2: cannot refer to unexported name something.bye

	name1 := "jam"
	var name2 = "jam"
	var name3 string = "jam"

	fmt.Println(name1, name2, name3)

	printDivider("#1.3 functions part one")
	fmt.Println(multiply(13, 17))
	_, upperName := lenAndUpper("jam") // ignore은 underscore(_)를 사용한다
	fmt.Println(upperName)
	repeatMe("A", "B", "C")
}
