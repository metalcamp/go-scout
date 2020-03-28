package slotech

import (
	"fmt"
	"github.com/gocolly/colly"
	"time"
)

type jobItem struct {
	Title         string
	Company       string
	AdURL         string
	SourceBaseUrl string
	CrawledAt     time.Time
	PublishedAt   time.Time
}

var baseURL = "https://slo-tech.com"

func Scrape() []jobItem {
	jobs := []jobItem{}

	c := colly.NewCollector(
		colly.AllowedDomains("slo-tech.com"),
		colly.Async(true),
	)

	c.OnHTML("table.forums tr", func(e *colly.HTMLElement) {
		temp := jobItem{}
		temp.AdURL = baseURL + e.ChildAttr("td.name h3 a", "href")
		temp.SourceBaseUrl = baseURL
		temp.Title = e.ChildText("td.name h3 a")
		temp.Company = e.ChildText("td.company a")
		temp.CrawledAt = time.Now()
		//temp.PublishedAt = time.Now().UTC.format((e.ChildAttr("td.last_msg time", 'datetime'))
		jobs = append(jobs, temp)
	})

	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(baseURL + "/delo")
	c.Wait()

	//fmt.Println(jobs)
	return jobs
}

