package pubsub

import "github.com/oklog/ulid/v2"

type Topic struct {
	Id       ulid.ULID
	Name     string
	Messages []Message
	Channel  chan Message
}

func CreateTopic(name string, channel chan Message) *Topic {
	return &Topic{
		Id:      ulid.Make(),
		Name:    name,
		Channel: channel,
	}
}

func (t *Topic) AddMessage(message Message) {
	for _, m := range t.Messages {
		if m.Id == message.Id {
			return
		}
	}

	t.Messages = append(t.Messages, message)
}

func (t *Topic) RemoveMessage(message Message) {
	for i, m := range t.Messages {
		if m.Id == message.Id {
			t.Messages[i] = t.Messages[len(t.Messages)-1]
		}
	}

	t.Messages = t.Messages[:len(t.Messages)-1]
}

func (t *Topic) HasMessage(id ulid.ULID) bool {
	for _, m := range t.Messages {
		if m.Id == id {
			return true
		}
	}

	return false
}

func (t *Topic) Publish() {
	for _, m := range t.Messages {
		t.Channel <- m
	}
	close(t.Channel)
}
