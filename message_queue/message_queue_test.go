package message_queue

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func produce(broker *Broker[int]) {
	for i := 0; ; i++ {
		time.Sleep(time.Second * 2)

		go func() {
			log.Println("Produce", i)
			broker.Produce(i)
		}()
	}
}

func consume(broker *Broker[int]) {
	for {
		time.Sleep(time.Second * 1)

		go func() {
			msg := broker.Consume()

			log.Println("Consume", msg)
		}()
	}
}

func TestMessageQueue(t *testing.T) {
	broker := NewBroker[int]()
	sig, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGKILL)
	defer stop()

	withTimeout, cancel := context.WithTimeout(context.Background(), time.Second*11)
	defer cancel()

	go produce(broker)
	go consume(broker)

	select {
	case <-withTimeout.Done():
	case <-sig.Done():
	}
}
