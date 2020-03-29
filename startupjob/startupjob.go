package startupjob

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

const domain string = "startupjob.si"
const baseURL = "https://" + domain
const sourceName = "StartupJob"

func Scrape() []jobItem {
	jobs := []jobItem{}

	c := colly.NewCollector(
		colly.AllowedDomains(domain, "www." + domain),
		colly.Async(true),
	)

	//loading more than last 8 jobs requires javascript enabled
	//since we are interested only in latest jobs and jobs are posted infrequently that's not a deal breaker
	c.OnHTML("ul.job_listings", func(element *colly.HTMLElement) {
		element.ForEach("li.job_listing", func(i int, element *colly.HTMLElement) {
			temp := jobItem{}
			temp.AdURL = element.Attr("data-href")
			temp.SourceBaseUrl = baseURL
			temp.SourceName = sourceName
			temp.Title = element.ChildText("div.job_listing-about > div.job_listing-position job_listing__column > h3")
			temp.Company = element.ChildText("div.job_listing-about > div.job_listing-position job_listing__column > div.job_listing-company > strong")
			temp.CrawledAt = time.Now()
			jobs = append(jobs, temp)
		})
	})

	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(baseURL)
	c.Wait()

	return jobs
}
