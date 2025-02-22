package main

import "log"

func IsClosed[T any](c chan T) bool {
	select {
	case <-c:
		return true
	default:
	}
	return false
}

func main() {
	ch := make(chan int)

	//close(ch)

	//val := <-ch

	log.Println(IsClosed(ch))

	//log.Println(val)
}
