package pubsub

type Topic struct {
	Name     string
	Messages []Message
	Channel  chan Message
}

func (t *Topic) Publish() {
	for _, m := range t.Messages {
		t.Channel <- m
	}
	close(t.Channel)
}
