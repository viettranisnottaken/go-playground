package concurrency

import (
	"log"
	"math/rand/v2"
	"sync"
	"time"
)

var PHARMACIES = []string{
	"phar1",
	"phar2",
	"phar4",
	"phar4",
}

func dbWorker(wg *sync.WaitGroup, mu *sync.Mutex, workerId int, dict *map[string]int) {
	defer wg.Done()

	log.Printf("worker %d working", workerId)
	time.Sleep(time.Duration(rand.IntN(3)))

	// pharId, inputString, drugName, timestamp

	data := map[string]int{
		"phar1": rand.Int(),
		"phar2": rand.Int(),
		"phar3": rand.Int(),
		"phar4": rand.Int(),
	}

	mu.Lock()
	defer mu.Unlock()
	for k, v := range data {
		(*dict)[k] = v
	}

	log.Printf("worker %d done", workerId)
}

func RunDbScript() {
	dict := make(map[string]int)
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(16)
	for i := 0; i < 16; i++ {
		go dbWorker(&wg, &mu, i, &dict)
	}

	wg.Wait()

	log.Printf("Everything done %v", dict)
}
