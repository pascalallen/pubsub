package pubsub

import (
	"sync"
)

type Subscriber struct {
	Id       string
	Messages chan *Message
	Topics   []string
	Active   bool
	Mutex    *sync.RWMutex
}

func (s *Subscriber) AddTopic(topic string) {
	for _, t := range s.Topics {
		if t == topic {
			return
		}
	}

	s.Mutex.RLock()
	defer s.Mutex.RUnlock()
	s.Topics = append(s.Topics, topic)
}

func (s *Subscriber) RemoveTopic(topic string) {
	for i, t := range s.Topics {
		if t == topic {
			s.Topics[i] = s.Topics[len(s.Topics)-1]
		}
	}

	s.Mutex.RLock()
	defer s.Mutex.RUnlock()
	s.Topics = s.Topics[:len(s.Topics)-1]
}
