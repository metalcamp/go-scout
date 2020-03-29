package main

import (
	"fmt"
	"scout/slotech"
	"scout/startupjob"
)

func main() {
	var slotechJobs = slotech.Scrape()
	var startupJobs = startupjob.Scrape()

	fmt.Println(slotechJobs)
	fmt.Println(startupJobs)
}
