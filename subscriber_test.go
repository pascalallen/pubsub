package pubsub

import (
	"testing"
)

func TestThatCreateSubscriberReturnsInstanceOfSubscriber(t *testing.T) {
	c := make(chan Message)
	to := CreateTopic("my-topic", c)
	m := CreateMessage([]byte("hello, jupiter!"))
	to.AddMessage(*m)

	s := CreateSubscriber("my-subscription", *to, c)

	if s == nil {
		t.Fatal("test failed attempting to CreateSubscriber")
	}
}
