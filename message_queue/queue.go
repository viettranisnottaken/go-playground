package message_queue

import (
	"log"
	"sync"
)

type Queue[T any] struct {
	Data        chan *Message[T]
	Subscribers map[*func(msg *Message[T])]func(msg *Message[T])
	mu          sync.RWMutex
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		Data:        make(chan *Message[T], 10),
		Subscribers: make(map[*func(msg *Message[T])]func(msg *Message[T])),
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

func (msq *Queue[T]) Publish(payload T) error {
	message, err := NewMessage(payload)

	if err != nil {
		log.Println(err)
		return err
	}

	msq.mu.Lock()
	msq.mu.Unlock()

	if len(msq.Subscribers) > 0 {
		msq.CallSubscribers(message)

		return nil
	}

	msq.Data <- message
	return nil
}

func (msq *Queue[T]) Dequeue() *Message[T] {
	//msq.mu.Lock()
	//defer msq.mu.Unlock()

	message := <-msq.Data
	return message
}

func (msq *Queue[T]) AddSubscriber(key *func(msg *Message[T]), val func(msg *Message[T])) {
	msq.mu.Lock()
	defer msq.mu.Unlock()

	msq.Subscribers[key] = val
}

func (msq *Queue[T]) RemoveSubscriber(key *func(msg *Message[T])) {
	msq.mu.Lock()
	defer msq.mu.Unlock()

	delete(msq.Subscribers, key)
}

func (msq *Queue[T]) CallSubscribers(msg *Message[T]) {
	msq.mu.RLock()
	defer msq.mu.RUnlock()

	for _, fn := range msq.Subscribers {
		fn(msg)
	}
}
