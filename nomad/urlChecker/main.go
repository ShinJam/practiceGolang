package main

import (
	"fmt"
	"net/http"
	"time"
)

type status struct {
	code   int
	reason string
}

type result struct {
	url    string
	status status
}

// var errRequestFailed = errors.New("request failed")

func main() {

	// var results map[string]string
	//   ;cannot asign to uninitialized map
	//   ;panic: assignment to entry in nil map

	// var results = make(map[string]string) // map[string]string{}
	c := make(chan result)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	go printResult(c)
	for _, url := range urls {
		go hitURL(url, c)
	}
	time.Sleep(5 * time.Second)
	close(c)
}

func hitURL(url string, c chan<- result) {
	fmt.Println("Checking url:", url)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	c <- result{
		url: url,
		status: status{
			code:   resp.StatusCode,
			reason: http.StatusText(resp.StatusCode),
		},
	}
}

func printResult(c chan result) {
	for result := range c {
		fmt.Println(result)
	}
}
