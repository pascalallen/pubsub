package pubsub

import "github.com/oklog/ulid/v2"

type Message struct {
	Id   ulid.ULID
	Data []byte
}

func CreateMessage(data []byte) *Message {
	return &Message{
		Id:   ulid.Make(),
		Data: data,
	}
}
