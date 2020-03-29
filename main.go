package main

import (
	"fmt"
	"scout/remoteok"
	"scout/slotech"
	"scout/startupjob"
)

func main() {
	var slotechJobs = slotech.Scrape()
	var startupJobs = startupjob.Scrape()
	var remoteokjobs = remoteok.Scrape()

	fmt.Println(slotechJobs)
	fmt.Println(startupJobs)
	fmt.Println(remoteokjobs)
}
