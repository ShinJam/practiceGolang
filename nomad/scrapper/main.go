package main

import (
	"os"
	"scrapper/scrapper"
	"strings"

	"github.com/labstack/echo/v4"
)

const fileName = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
	// return c.String(http.StatusOK, "Hello, World!")
}

func handleScrape(c echo.Context) error {
	defer os.Remove("jobs.csv")
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
	// scrapper.Scrape("python")
}
