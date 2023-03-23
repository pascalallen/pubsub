package pubsub

import "sync"

type Subscribers map[string]*Subscriber

type Broker struct {
	Subscribers Subscribers
	Topics      map[string]Subscribers
	Mutex       sync.RWMutex
}
