package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		client := c.Query("client")
		fmt.Println("== request from:", client)

		// Calling NewTicker method
		Ticker := time.NewTicker(1 * time.Second)

		// Creating channel using make
		// keyword
		mychannel := make(chan bool)

		// Go function
		go func() {

			// Using for loop
			for {

				// Select statement
				select {

				// Case statement
				case <-mychannel:
					return

				// Case to print current time
				case tm := <-Ticker.C:
					fmt.Println(tm)
				}
			}
		}()

		// Calling Sleep() method
		time.Sleep(3 * time.Second)

		// Calling Stop() method
		Ticker.Stop()

		// Setting the value of channel
		mychannel <- true

		ret := fmt.Sprintf("== End: client: %v", client)
		log.Println(ret)
		return c.SendString(ret)
	})

	app.Listen(":3000")
}
