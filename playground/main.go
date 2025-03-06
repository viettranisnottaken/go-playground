package main

import (
	"log"
	"time"
)

type TokenBucket[TRequest any] struct {
	requests       chan TRequest
	bucket         chan time.Time
	period         time.Duration
	limit          int
	timePerRequest time.Duration
}

func NewTokenBucket[TRequest any](
	requests chan TRequest,
	period time.Duration,
	limit int,
) *TokenBucket[TRequest] {
	timePerRequest := period / time.Duration(limit)
	bucket := make(chan time.Time, limit)

	return &TokenBucket[TRequest]{
		requests:       requests,
		bucket:         bucket,
		period:         period,
		limit:          limit,
		timePerRequest: timePerRequest,
	}
}

func (rl *TokenBucket[TRequest]) HandleRequests(handler func(request TRequest, timestamp time.Time)) {
	// refill
	go func() {
		tick := time.NewTicker(rl.timePerRequest)

		for t := range tick.C {
			select {
			case rl.bucket <- t:
			default:
			}
		}
	}()

	// listen to requests and handle
	for request := range rl.requests {
		timestamp := <-rl.bucket

		go handler(request, timestamp)
	}
}

func main() {
	requests := make(chan int)
	rl := NewTokenBucket[int](requests, time.Second, 2)

	go rl.HandleRequests(func(request int, timestamp time.Time) {
		log.Println("Request no.", request, "handled at", timestamp)
	})

	for i := 0; ; i++ {
		time.Sleep(time.Millisecond * 250)
		select {
		case requests <- i:
		default:
			log.Printf("Request %d at %s denied\n", i, time.Now())
		}
	}
}
