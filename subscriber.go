package pubsub

import "github.com/oklog/ulid/v2"

type Subscriber struct {
	Id      ulid.ULID
	Name    string
	Topic   Topic
	Channel chan Message
}

// CreateSubscriber creates and returns a pointer instance of Subscriber.
func CreateSubscriber(name string, topic Topic, channel chan Message) *Subscriber {
	return &Subscriber{
		Id:      ulid.Make(),
		Name:    name,
		Topic:   topic,
		Channel: channel,
	}
}
