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
	SourceName    string
	CrawledAt     time.Time
	PublishedAt   time.Time
}

const domain string = "slo-tech.com"
const baseURL = "https://" + domain
const jobsURL = baseURL + "/delo"
const sourceName = "SloTech"

func Scrape() []jobItem {
	jobs := []jobItem{}

	c := colly.NewCollector(
		colly.AllowedDomains(domain, "wwww." + domain),
		colly.Async(true),
	)

	c.OnHTML("table.forums tr", func(e *colly.HTMLElement) {
		temp := jobItem{}
		temp.AdURL = baseURL + e.ChildAttr("td.name h3 a", "href")
		temp.SourceBaseUrl = baseURL
		temp.Title = e.ChildText("td.name h3 a")
		temp.Company = e.ChildText("td.company a")
		temp.CrawledAt = time.Now()
		temp.SourceName = sourceName
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

	c.Visit(jobsURL)
	c.Wait()

	return jobs
}
