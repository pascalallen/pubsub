package pubsub

import "sync"

type Subscriber struct {
	Id       string
	Messages chan *Message
	Topics   map[string]bool
	Active   bool
	Mutex    sync.RWMutex
}
