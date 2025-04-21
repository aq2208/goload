package mq

import "context"

type ConsumerHanlerFunc func(ctx context.Context, queueName string, payload []byte) error

type Consumer interface {
	RegisterHandler(queueName string, handlerFunc ConsumerHanlerFunc)
	Start(ctx context.Context) error
}

// build separate handler for each type of event