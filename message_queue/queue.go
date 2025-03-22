package message_queue

import (
	"sync"
)

type Queue[T any] struct {
	data chan *Message[T]
	mu   sync.Mutex
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		data: make(chan *Message[T], 10),
	}
}

func (msq *Queue[T]) Enqueue(payload T) error {
	message, err := NewMessage(payload)

	if err != nil {
		return err
	}

	//msq.mu.Lock()
	//defer msq.mu.Unlock()

	//msq.data = append(msq.data, message)

	msq.data <- message
	return nil
}

func (msq *Queue[T]) Dequeue() *Message[T] {
	message := <-msq.data
	return message
}
