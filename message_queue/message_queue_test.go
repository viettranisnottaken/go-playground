package message_queue

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"reflect"
	"sort"
	"sync"
	"syscall"
	"testing"
	"time"
)

func produce(broker *Broker[int], topic Topic) {
	for i := 0; ; i++ {
		time.Sleep(time.Second * 1)

		go func() {
			log.Println("Produce", topic, i)
			err := broker.Produce(topic, i)

			if err != nil {
				log.Fatal(err)
			}
		}()
	}
}

func consumeOnce(broker *Broker[int], topic Topic) {
	for {
		time.Sleep(time.Second * 2)

		go func() {
			msg, _ := broker.ConsumeOnce(topic)

			log.Println("Consume", topic, msg)
		}()
	}
}

func consume(broker *Broker[int], topic Topic) {
	unsub := broker.Consume(topic, func(msg *Message[int]) {
		log.Println("Consume", topic, msg)
	})

	go func() {
		time.Sleep(4 * time.Second)

		unsub()
	}()
}

func TestMessageQueueConsumeOnce(t *testing.T) {
	broker := NewBroker[int]()
	sig, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGKILL)
	defer stop()

	withTimeout, cancel := context.WithTimeout(context.Background(), time.Second*11)
	defer cancel()

	go produce(broker, "topic")
	go consumeOnce(broker, "topic")

	select {
	case <-withTimeout.Done():
	case <-sig.Done():
	}
}

func TestMessageQueueConsume(t *testing.T) {
	broker := NewBroker[int]()
	sig, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGKILL)
	defer stop()

	withTimeout, cancel := context.WithTimeout(context.Background(), time.Second*11)
	defer cancel()

	go produce(broker, "topic")
	go produce(broker, "topic")
	go produce(broker, "topic")

	go consume(broker, "topic")
	//go consume(broker, "topic")
	//go consume(broker, "topic")

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

	msg, err := broker.ConsumeOnce(topic)

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

func publish(wg *sync.WaitGroup, broker *Broker[string], topic Topic, payload string) {
	defer wg.Done()

	err := broker.Publish(topic, payload)
	if err != nil {
		log.Println(err)
	}
}

func subscribe(wg *sync.WaitGroup, subCh chan *Message[string], broker *Broker[string], topic Topic) {
	defer wg.Done()

	handler := func(msg *Message[string]) {
		log.Println("subscribe", msg)
		subCh <- msg
	}

	unsub := broker.Subscribe(topic, &handler)

	go func() {
		time.Sleep(2 * time.Second)

		unsub()
	}()
}

func TestSubscribe(t *testing.T) {
	//withTimeout, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	//defer cancel()
	//
	//sigint, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGKILL)
	//defer stop()

	var wg sync.WaitGroup

	broker := NewBroker[string]()
	subCh := make(chan *Message[string], 20)

	wg.Add(3)
	go subscribe(&wg, subCh, broker, "topic")
	go subscribe(&wg, subCh, broker, "topic")
	go subscribe(&wg, subCh, broker, "topic")
	wg.Wait()

	for i := 0; i < 4; i++ {
		time.Sleep(time.Millisecond * 500)
		wg.Add(1)
		go publish(&wg, broker, "topic", fmt.Sprintf("payload %d", i))
	}
	wg.Wait()

	withSubTimeout, subCancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer subCancel()

	go func() {
		<-withSubTimeout.Done()

		close(subCh)
	}()

	res := make([]Message[string], 0)
	for msg := range subCh {
		res = append(res, *msg)
	}

	if len(res) != 9 {
		t.Errorf("Expected to get 9 values from subscription, instead got %d", len(res))
	}

	//select {
	//case <-withTimeout.Done():
	//case <-sigint.Done():
	//}
}
