package pubsub

import (
	"sync"
)

type Broker struct {
	Subscribers []Subscriber
	//Topics      map[string]Subscribers
	Mutex sync.RWMutex
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
