package main

import (
	"fmt"
	"time"
)

// later this will be long-running aggregator service.
// For now it will be used to fetch a single feed to ensure parsing works
// It should fetch the feed found at https://www.wagslane.dev/index.xml and print the entire struct to the console.
func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("No time provided. Usage: agg <time> | where <time> is in the format: 1s/1m/1h")
	}
	timer := cmd.Args[0]
	fmt.Printf("Collecting feeds every %s", timer)
	timeBetweenRequests, err := time.ParseDuration(timer)
	if err != nil {
		return fmt.Errorf("Error parsing time value: %v", err)
	}
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeed(s)
	}
}
