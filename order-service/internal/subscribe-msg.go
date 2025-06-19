package internal

import (
	"errors"
	"time"

	"github.com/nats-io/nats.go"
)

func WaitForMessage(nc *nats.Conn, subject string, timeout time.Duration) ([]byte, error) {
	msgCh := make(chan *nats.Msg, 1)
	sub, err := nc.Subscribe(subject, func(msg *nats.Msg) {
		msgCh <- msg
	})
	if err != nil {
		return nil, err
	}
	defer sub.Unsubscribe()

	select {
	case msg := <-msgCh:
		return msg.Data, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout waiting for message")
	}
}
