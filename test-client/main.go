package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type response struct {
	Hello  string `json: "hello"`
	Client string `json: "client"`
}

func main() {
	defer timeTrack(time.Now(), "main")
	server := []string{"A", "B", "C", "D", "E"}

	var wg sync.WaitGroup
	wg.Add(len(server))
	for _, s := range server {
		go func(s string) {
			defer wg.Done()
			res, err := http.Get("http://localhost:8000/?client=" + s)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()
			r := new(response)
			json.NewDecoder(res.Body).Decode(&r)
			log.Println(r)
		}(s)
		// go request(s)
	}
	wg.Wait()
}

func timeTrack(start time.Time, name string) {
	// ref: https://coderwall.com/p/cp5fya/measuring-execution-time-in-go
	elapsed := time.Since(start)
	fmt.Printf("!! %s took %s", name, elapsed)
}
