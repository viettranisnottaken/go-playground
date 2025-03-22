package message_queue

import (
	"github.com/nu7hatch/gouuid"
)

type Message[T any] struct {
	id      *uuid.UUID
	payload T
}

func NewMessage[T any](payload T) (*Message[T], error) {
	id, err := uuid.NewV4()

	if err != nil {
		return nil, err
	}

	return &Message[T]{
		id:      id,
		payload: payload,
	}, nil
}
