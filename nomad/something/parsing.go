package something

import (
	"fmt"
	"reflect"
	"strconv"
)

func NormalParsing() {
	/*
	 * Ref: https://ithub.tistory.com/331
	 */

	// 1. int to string - 숫자(정수)를 문자열로 변환
	a := strconv.Itoa(100)
	fmt.Println("a: ", a)                      // a: 100
	fmt.Println("type a: ", reflect.TypeOf(a)) // type a: string

	// 1-1. int to string - 100을 10진수 문자열로 변환
	aa := strconv.FormatInt(100, 10)
	fmt.Println("aa: ", aa)                      // aa: 100
	fmt.Println("type aa: ", reflect.TypeOf(aa)) // type aa: string

	// 2. string to int - 문자열을 숫자(정수) 변환
	b, _ := strconv.Atoi("100")
	fmt.Println("b: ", b)                      // b:  100
	fmt.Println("type b: ", reflect.TypeOf(b)) // type b: int

	bb, _ := strconv.ParseInt("100", 10, 64)
	fmt.Println("bb: ", bb)                      // bb: 100
	fmt.Println("type bb: ", reflect.TypeOf(bb)) // type bb: int64

	// 3. bool to string - 불을 문자열로 변환
	c := strconv.FormatBool(true)
	fmt.Println("c: ", c)                      // c: true
	fmt.Println("type c: ", reflect.TypeOf(c)) // type c: string

	// 4. flot to string - 숫자(실수)를 문자열로 변환
	d := strconv.FormatFloat(1.3, 'f', -1, 32)
	fmt.Println("d: ", d)                      // d: 1.3
	fmt.Println("type d: ", reflect.TypeOf(d)) //type d: string

	// 5. int -> int32, int32 -> int64
	var e int = 11
	f := int32(e)
	fmt.Println("f: ", f)                      // f: 11
	fmt.Println("type f: ", reflect.TypeOf(f)) // type f: int32

	g := int64(f)
	fmt.Println("g: ", g)                      // g:  11
	fmt.Println("type g: ", reflect.TypeOf(g)) // type g: int64
}

func ParsingWithSprint() string {
	/*
	 * Ref: https://stackoverflow.com/questions/11123865/format-a-go-string-without-printing
	 */

	what := "apple"
	when := 10

	return fmt.Sprintf("%v %v", what, when)
}
