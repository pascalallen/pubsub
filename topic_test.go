package pubsub

import (
	"testing"
)

func TestThatCreateTopicReturnsInstanceOfTopic(t *testing.T) {
	c := make(chan Message)
	to := CreateTopic("my-topic", c)

	if to == nil {
		t.Fatal("test failed attempting to CreateTopic")
	}
}

func TestThatAddMessageAppendsMessageToMessages(t *testing.T) {
	c := make(chan Message)
	to := CreateTopic("my-topic", c)
	m := CreateMessage([]byte("hello, jupiter!"))
	to.AddMessage(*m)

	if !to.HasMessage(m.Id) {
		t.Fatal("test failed attempting to AddMessage")
	}
}

func TestThatAddMessageReturnsEarlyIfMessageHasAlreadyBeenAdded(t *testing.T) {
	c := make(chan Message)
	to := CreateTopic("my-topic", c)
	m := CreateMessage([]byte("hello, jupiter!"))
	to.AddMessage(*m)
	to.AddMessage(*m)

	if len(to.Messages) != 1 {
		t.Fatal("test failed attempting to add duplicate message")
	}
}

func TestThatRemoveMessageRemovesMessageFromMessages(t *testing.T) {
	c := make(chan Message)
	to := CreateTopic("my-topic", c)
	m := CreateMessage([]byte("hello, jupiter!"))
	to.AddMessage(*m)
	to.RemoveMessage(*m)

	if to.HasMessage(m.Id) {
		t.Fatal("test failed attempting to RemoveMessage")
	}
}

func TestThatPublishSendsMessageToTheChannelAndClosesChannel(t *testing.T) {
	c := make(chan Message)
	to := CreateTopic("my-topic", c)
	m := CreateMessage([]byte("hello, jupiter!"))
	to.AddMessage(*m)

	go to.Publish()

	me := <-c
	_, ok := <-c

	if !to.HasMessage(me.Id) || ok {
		t.Fatal("test failed attempting to Publish messages")
	}
}
