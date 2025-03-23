package message_queue

import (
	"fmt"
	"sync"
)

type Topic string

type Broker[T any] struct {
	queues map[Topic]*Queue[T]
	done   chan bool
	mu     sync.Mutex
}

func NewBroker[T any]() *Broker[T] {
	return &Broker[T]{
		queues: make(map[Topic]*Queue[T]),
		done:   make(chan bool),
	}
}

func (b *Broker[T]) Produce(topic Topic, payload T) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	_, ok := b.queues[topic]

	if !ok {
		b.queues[topic] = NewQueue[T]()
	}

	err := b.queues[topic].Enqueue(payload)

	return err
}

func (b *Broker[T]) Consume(topic Topic) (*Message[T], error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	queue, ok := b.queues[topic]

	if !ok {
		return nil, fmt.Errorf("the topic %s does not exist in the message queue", topic)
	}

	return queue.Dequeue(), nil
}

func (b *Broker[T]) Shutdown() {
	// todo
}
