package pubsub

import "github.com/oklog/ulid/v2"

type Topic struct {
	Id       ulid.ULID
	Name     string
	Messages []Message
	Channel  chan Message
}

// CreateTopic creates and returns a pointer instance of Topic.
func CreateTopic(name string, channel chan Message) *Topic {
	return &Topic{
		Id:      ulid.Make(),
		Name:    name,
		Channel: channel,
	}
}

// AddMessage appends the given message to the Topic's messages field.
// Returns early if the message already exists.
func (t *Topic) AddMessage(message Message) {
	for _, m := range t.Messages {
		if m.Id == message.Id {
			return
		}
	}

	t.Messages = append(t.Messages, message)
}

// RemoveMessage removes the given message from the Topic's messages field.
func (t *Topic) RemoveMessage(message Message) {
	for i, m := range t.Messages {
		if m.Id == message.Id {
			t.Messages[i] = t.Messages[len(t.Messages)-1]
		}
	}

	t.Messages = t.Messages[:len(t.Messages)-1]
}

// HasMessage returns true if the given ID matches a message's ID in the Topic's messages field,
// returns false otherwise.
func (t *Topic) HasMessage(id ulid.ULID) bool {
	for _, m := range t.Messages {
		if m.Id == id {
			return true
		}
	}

	return false
}

// Publish sends each message to the Topic's channel and then closes the channel.
func (t *Topic) Publish() {
	for _, m := range t.Messages {
		t.Channel <- m
	}
	close(t.Channel)
}
