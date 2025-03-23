package message_queue

import (
	"context"
	"log"
	"os/signal"
	"reflect"
	"sort"
	"sync"
	"syscall"
	"testing"
	"time"
)

func produce(broker *Broker[int]) {
	for i := 0; ; i++ {
		time.Sleep(time.Second * 1)

		go func() {
			log.Println("Produce", i)
			err := broker.Produce("topic", i)

			if err != nil {
				log.Fatal(err)
			}
		}()
	}
}

func consume(broker *Broker[int]) {
	for {
		time.Sleep(time.Second * 2)

		go func() {
			msg, _ := broker.Consume("topic")

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

func produceWithTopic(wg *sync.WaitGroup, broker *Broker[string], topic Topic, payload string) error {
	defer wg.Done()

	err := broker.Produce(topic, payload)

	if err != nil {
		return err
	}

	return nil
}

func consumeWithTopic(wg *sync.WaitGroup, broker *Broker[string], topic Topic) (*Message[string], error) {
	defer wg.Done()

	msg, err := broker.Consume(topic)

	if err != nil {
		return nil, err
	}

	return msg, nil
}

func case1(broker *Broker[string], t *testing.T) {
	log.Println("Message queue should produce topics T1, T2 and T3, and consumer should be able to receive all values corresponding to topics")

	var prodWg sync.WaitGroup
	var conWg sync.WaitGroup
	resultCh := make(chan string, 3)
	result := make([]string, 0)

	prodWg.Add(3)
	go func() {
		err := produceWithTopic(&prodWg, broker, "T1", "Message for T1")
		if err != nil {
			t.Error(err)
		}
	}()
	go func() {
		err := produceWithTopic(&prodWg, broker, "T2", "Message for T2")
		if err != nil {
			t.Error(err)

		}
	}()
	go func() {
		err := produceWithTopic(&prodWg, broker, "T3", "Message for T3")
		if err != nil {
			t.Error(err)
		}
	}()

	prodWg.Wait()

	conWg.Add(3)
	go func() {
		msg, err := consumeWithTopic(&conWg, broker, "T1")
		if err != nil {
			t.Error(err)
		}

		resultCh <- msg.payload
	}()
	go func() {
		msg, err := consumeWithTopic(&conWg, broker, "T2")
		if err != nil {
			t.Error(err)
		}

		resultCh <- msg.payload
	}()
	go func() {
		msg, err := consumeWithTopic(&conWg, broker, "T3")
		if err != nil {
			t.Error(err)
		}

		resultCh <- msg.payload
	}()

	conWg.Wait()

	for res := range resultCh {
		result = append(result, res)

		if len(result) == 3 {
			close(resultCh)
		}
	}

	sort.Strings(result)

	expect := []string{"Message for T1", "Message for T2", "Message for T3"}
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Expected ouput to be %v, but got %v", expect, result)
	}

	log.Println("Result for case 1", result)
	log.Println("======= Case 1 done")
	log.Println("")
}

func case2(broker *Broker[string], t *testing.T) {
	log.Println("Consumer should return error when it cannot find topic")

	var conWg sync.WaitGroup

	conWg.Add(1)

	go func() {
		_, err := consumeWithTopic(&conWg, broker, "T1")
		if err == nil {
			t.Error(err)
		}
	}()

	conWg.Wait()
	log.Println("======= Case 2 done")
	log.Println("")
}

func TestBrokerWithTopic(t *testing.T) {
	log.Println("Message queue with topic")

	broker1 := NewBroker[string]()
	case1(broker1, t)

	broker2 := NewBroker[string]()
	case2(broker2, t)
}
