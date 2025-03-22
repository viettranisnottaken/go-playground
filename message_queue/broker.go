package message_queue

import "log"

type Broker[T any] struct {
	queue *Queue[T]
	done  chan bool
}

func NewBroker[T any]() *Broker[T] {
	return &Broker[T]{
		queue: NewQueue[T](),
		done:  make(chan bool),
	}
}

func (b Broker[T]) Produce(payload T) {
	go func() {
		err := b.queue.Enqueue(payload)
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func (b Broker[T]) Consume() *Message[T] {
	return b.queue.Dequeue()
}
