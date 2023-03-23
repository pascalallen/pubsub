package pubsub

import (
	"sync"
)

type Broker struct {
	Subscribers []Subscriber
	Mutex       sync.RWMutex
}

func (b *Broker) AddSubscriber(subscriber Subscriber) {
	for _, s := range b.Subscribers {
		if s.Id == subscriber.Id {
			return
		}
	}

	b.Mutex.Lock()
	defer b.Mutex.Unlock()
	b.Subscribers = append(b.Subscribers, subscriber)
}

func (b *Broker) RemoveSubscriber(subscriber Subscriber) {
	for _, t := range subscriber.Topics {
		subscriber.RemoveTopic(t)
	}

	b.Mutex.Lock()
	for i, s := range b.Subscribers {
		if s.Id == subscriber.Id {
			b.Subscribers[i] = b.Subscribers[len(b.Subscribers)-1]
		}
	}
	b.Subscribers = b.Subscribers[:len(b.Subscribers)-1]
	b.Mutex.Unlock()

	subscriber.Mutex.RLock()
	defer subscriber.Mutex.RUnlock()
	close(subscriber.Messages)
}
