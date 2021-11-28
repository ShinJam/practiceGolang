package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

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
var jobURL string = baseURL + "jobs?q=data"

func main() {
	defer timeTrack(time.Now(), "main")

	var wgCSV sync.WaitGroup
	var jobs []extrasctedJob
	c := make(chan []extrasctedJob)
	totalPages := getPages()
	log.Println("totalPages:", totalPages)
	for i := 0; i < totalPages; i++ {
		go getPage(i, c, &wgCSV)
		// jobs = append(jobs, extractedJobs...)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}
	wgCSV.Wait()
	log.Println("Done, extracted", len(jobs))
}

func writeCSV(job extrasctedJob, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.OpenFile("./jobs.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkErr(err)
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	jobSlice := []string{baseURL + "채용보기?jk=" + job.id, job.title, job.location, job.summary}
	jwErr := w.Write(jobSlice)
	checkErr(jwErr)
}

func getPage(page int, mainC chan<- []extrasctedJob, wg *sync.WaitGroup) {
	var jobs []extrasctedJob
	c := make(chan extrasctedJob)
	pageURL := jobURL + "&limit=50&start=" + strconv.Itoa(page*50)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()
	fmt.Println(pageURL, res.StatusCode)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".result")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})
	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		wg.Add(1)
		go writeCSV(job, wg)
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extrasctedJob) {
	id, exists := card.Attr("data-jk")
	if !exists {
		log.Panic("'data-jk' attribute does not exists")
	}
	title := cleanString(card.Find(".jobTitle").Text())
	location := cleanString(card.Find(".companyLocation").Text())
	summary := cleanString(card.Find(".job-snippet").Text())
	c <- extrasctedJob{
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
	res, err := http.Get(jobURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	countPage := cleanString(doc.Find("#searchCountPages").Text())
	fmt.Println(countPage)
	// re := regexp.MustCompile(`.*([0-9]+)건`)
	// // ref:  https://stackoverflow.com/questions/30483652/how-to-get-capturing-group-functionality-in-go-regular-expressions
	re := regexp.MustCompile(`.* (?P<Pages>.+)건`)
	match := re.FindStringSubmatch(countPage)[1]
	match = strings.Join(strings.Split(match, ","), "")
	fmt.Println(match)
	pages, err := strconv.Atoi(match)
	checkErr(err)
	pages = pages / 50
	if mod := pages % 50; mod > 0 {
		pages += 1
	}
	return pages
}

// func writeJobs(jobs []extrasctedJob) {
// 	file, err := os.Create("jobs.csv")
// 	checkErr(err)
// 	w := csv.NewWriter(file)
// 	defer w.Flush()

// 	headers := []string{"Link", "Title", "Location", "Summary"}
// 	wErr := w.Write(headers)
// 	checkErr(wErr)

// 	for _, job := range jobs {
// 		jobSlice := []string{baseURL + "채용보기?jk=" + job.id, job.title, job.location, job.summary}
// 		jwErr := w.Write(jobSlice)
// 		checkErr(jwErr)
// 	}
// }

func checkErr(err error) {
	switch {
	case err != nil:
		log.Panic(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Panic("request failed with status", res.StatusCode)
	}
}

func timeTrack(start time.Time, name string) {
	// ref: https://coderwall.com/p/cp5fya/measuring-execution-time-in-go
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
