package mq

import (
	"context"

	"github.com/aq2208/goload/internal/model"
)

type Producer interface {
	Publish(ctx context.Context, queueName string, payload []byte) error
}

// Type of Events
type DownloadTaskCreated struct {
	DownloadTask model.DownloadTask
}

// Publish event to the MQ