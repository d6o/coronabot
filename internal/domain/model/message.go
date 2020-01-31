package model

import "context"

type (
	Message struct {
		text string
	}

	MessageSender interface {
		Send(context.Context, Message) error
	}
)

func NewMessage(t string) Message {
	return Message{
		text: t,
	}
}

func (m Message) Text() string {
	return m.text
}
