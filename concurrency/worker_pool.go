package concurrency

import (
	"sync"
)

/*
- A function `ProcessJobs(jobs []int, numWorkers int) []int` should:
- worker just doubles the job value
- Distribute the work among `numWorkers`.
- Each job should be processed by a worker and return a **result**.
- Return an array of processed results.
*/

/*
job chan and res chan
job chan because num workers != num jobs
bind workers to job chan
push jobs in job chan
push res in res chan
convert res chan to res slice
*/

func worker(jobChan <-chan int, resultCh chan<- int, wg *sync.WaitGroup, cb func(int) int) {
	defer wg.Done()
	for val := range jobChan {
		resultCh <- cb(val)
	}
}

func ProcessJobs(jobs []int, numWorkers int) []int {
	return processJobSol2(jobs, numWorkers)
}

func processJobSol1(jobs []int, numWorkers int) []int {
	// does not guarantee job result order

	jobChan := make(chan int, len(jobs))
	resultCh := make(chan int, len(jobs))
	res := make([]int, 0)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(jobChan, resultCh, &wg, func(i int) int {
			return i * 2
		})
	}

	for _, job := range jobs {
		// does not guarantee order if workers process at different speed (async, api calls, etc.)
		jobChan <- job
	}

	close(jobChan)

	wg.Wait()
	close(resultCh)

	for val := range resultCh {
		res = append(res, val)
	}

	return res
}

type Job struct {
	index int
	value int
}

func worker2(jobChan <-chan Job, res *[]int, wg *sync.WaitGroup, cb func(int) int) {
	defer wg.Done()
	for job := range jobChan {
		(*res)[job.index] = cb(job.value)
	}
}

func processJobSol2(jobs []int, numWorkers int) []int {
	res := make([]int, len(jobs))
	jobChan := make(chan Job, len(jobs))
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)

		go worker2(jobChan, &res, &wg, func(num int) int {
			return num * 2
		})
	}

	for i, job := range jobs {
		jobChan <- Job{i, job}
	}

	close(jobChan)

	wg.Wait()

	return res
}
