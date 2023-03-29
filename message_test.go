package pubsub

import (
	"testing"
)

func TestThatCreateMessageReturnsInstanceOfMessage(t *testing.T) {
	m := CreateMessage([]byte("hello, jupiter!"))

	if m == nil {
		t.Fatal("test failed attempting to CreateMessage")
	}
}
