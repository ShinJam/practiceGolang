package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extrasctedJob struct {
	id       string
	title    string
	location string
	// salary   string
	summary string
}

var baseURL string = "https://kr.indeed.com/"
var jobURL string = baseURL + "jobs?q=python"

func main() {
	var jobs []extrasctedJob
	totalPages := getPages()
	fmt.Println("totalPages:", totalPages)
	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...)
	}
	writeJobs(jobs)
	log.Println("finish")
}

func writeJobs(jobs []extrasctedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Link", "Title", "Location", "Summary"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{baseURL + "채용보기?jk=" + job.id, job.title, job.location, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func getPage(page int) []extrasctedJob {
	var jobs []extrasctedJob
	pageURL := jobURL + "&limit=" + strconv.Itoa(page*50)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()
	fmt.Println(pageURL, res.StatusCode)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCard := doc.Find(".result")

	searchCard.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})
	return jobs
}

func extractJob(card *goquery.Selection) extrasctedJob {
	id, exists := card.Attr("data-jk")
	if !exists {
		log.Panic("'data-jk' attribute does not exists")
	}
	title := cleanString(card.Find(".jobTitle").Text())
	location := cleanString(card.Find(".companyLocation").Text())
	summary := cleanString(card.Find(".job-snippet").Text())
	return extrasctedJob{
		id:       id,
		title:    title,
		location: location,
		summary:  summary,
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(str), " ")
}

func getPages() int {
	pages := 0
	res, err := http.Get(jobURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Panic("request failed with status", res.StatusCode)
	}
}
