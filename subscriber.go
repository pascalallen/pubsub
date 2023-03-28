package pubsub

type Subscriber struct {
	Name    string
	Topic   Topic
	Channel chan Message
}
