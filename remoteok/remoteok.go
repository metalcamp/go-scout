package remoteok

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
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

const domain string = "remoteok.io"
const baseURL = "https://" + domain
const jobsURL = baseURL + "/remote-dev-jobs"
const sourceName = "RemoteOK"

func Scrape() []jobItem {
	jobs := []jobItem{}

	c := colly.NewCollector(
		colly.AllowedDomains(domain, "wwww."+domain),
		colly.Async(true),
	)

	c.OnHTML("tr.job", func(e *colly.HTMLElement) {
		var relativeAdUrl = e.Attr("data-url")
		if relativeAdUrl == "" {
			return
		}
		temp := jobItem{}
		temp.AdURL = baseURL + relativeAdUrl
		temp.SourceBaseUrl = baseURL
		temp.Title = e.ChildText("td > h2[itemprop=title]")
		temp.Company = e.Attr("data-company")
		temp.CrawledAt = time.Now()
		temp.SourceName = sourceName

		timeCreatedTimestamp, err := strconv.ParseInt(e.Attr("data-epoch"), 10, 64)

		if err != nil {
			panic(err)
		}

		temp.PublishedAt = time.Unix(timeCreatedTimestamp, 0)
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
