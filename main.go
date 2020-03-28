package main

import (
	"fmt"
	"scout/slotech"
)

func main() {
	var jobs = slotech.Scrape()
	fmt.Println(jobs)
}
