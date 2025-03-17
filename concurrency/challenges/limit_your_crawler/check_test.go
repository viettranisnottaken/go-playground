package limit_your_crawler

import (
	"testing"
	"time"
)

//////////////////////////////////////////////////////////////////////
//
// DO NOT EDIT THIS PART
// Your task is to edit `main.go`
//

func TestLimitYourCrawler(t *testing.T) {
	fetchSig := fetchSignalInstance()

	start := time.Unix(0, 0)
	go func(start time.Time) {
		for {
			switch {
			case <-fetchSig:
				// Check if signal arrived earlier than a second (with error margin)
				if time.Now().Sub(start).Nanoseconds() < 950000000 {
					t.Log("There exists a two crawls that were executed less than 1 second apart.")
					t.Log("Solution is incorrect.")
					t.FailNow()
				}
				start = time.Now()
			}
		}
	}(start)

	Run()
}
