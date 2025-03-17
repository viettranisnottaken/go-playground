package limit_your_crawler

import (
	"fmt"
	"sync"
	"time"
)

//////////////////////////////////////////////////////////////////////
//
// Your task is to change the code to limit the crawler to at most one
// page per second, while maintaining concurrency (in other words,
// Crawl() must be called concurrently)
//
// @hint: you can achieve this by adding 3 lines
//

// Crawl uses `fetcher` from the `mockfetcher.go` file to imitate a
// real crawler. It crawls until the maximum depth has reached.
func Crawl(url string, depth int, wg *sync.WaitGroup, rm <-chan time.Time) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	wg.Add(len(urls))
	for _, u := range urls {
		// Do not remove the `go` keyword, as Crawl() must be
		// called concurrently
		<-rm // block here
		go Crawl(u, depth-1, wg, rm)
	}
}

func Run() {
	var wg sync.WaitGroup
	tokenBucket := time.Tick(time.Second / 1) // rate limiter

	wg.Add(1)
	Crawl("http://golang.org/", 4, &wg, tokenBucket)
	wg.Wait()
}
