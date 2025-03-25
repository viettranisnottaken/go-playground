package message_queue

import (
	"log"
	"sync"
)

type Queue[T any] struct {
	Data chan *Message[T]
	mu   sync.Mutex
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		Data: make(chan *Message[T], 10),
	}
}

func (msq *Queue[T]) Enqueue(payload T) error {
	message, err := NewMessage(payload)

	if err != nil {
		log.Println(err)
		return err
	}
	
	msq.Data <- message
	return nil
}

func (msq *Queue[T]) Dequeue() *Message[T] {
	msq.mu.Lock()
	defer msq.mu.Unlock()

	message := <-msq.Data
	return message
}
