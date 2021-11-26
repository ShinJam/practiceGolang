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

func lenAndUpperWithDefer(name string) (length int, uppercase string) { // 반환값을 반환 인자로 사용해줄 수 있다. return은 꼭 있어야 한다.
	defer fmt.Println("I'm done") // 함수가 끝나고 호출되는 함수, 여러개라면 아래서 부터 순차적으로 실행
	length, uppercase = len(name), strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0
	for index, number := range numbers {
		fmt.Println("index :", index)
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 { // 초기값을 할당 해줄수 있다.
		return false
	}
	return true
}

func canIDrinkSwitch(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 20:
		return true
	}
	return false
}

func main() {
	var ret interface{}

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

	printDivider("#1.4 functions part two")
	lenAndUpperWithDefer("jam")

	printDivider("#1.5 range args")
	ret = superAdd(1, 2, 3, 4, 5)
	fmt.Println(ret)

	printDivider("#1.6 if twist")
	ret = canIDrink(30)
	fmt.Println(ret)

	printDivider("#1.7 switch")
	ret = canIDrinkSwitch(30)
	fmt.Println(ret)

	printDivider("#1.9 Arrays and Slices")
	names := []string{"apple", "banana", "grape"}
	fmt.Println(names)
	names[1] = "strawberry"
	fmt.Println(names)
	names = append(names, "watermelon")
	fmt.Println(names, names[1:3])

	printDivider("#1.10 maps")
	profile := map[string]string{"name": "jam", "age": "12"}
	fmt.Println(profile)
	for key, value := range profile {
		fmt.Println(key, value)
	}
}
