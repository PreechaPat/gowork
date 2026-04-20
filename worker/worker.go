package worker

import "context"

type WorkBackend interface {
	Enqueue(ctx context.Context, work Work) error
	Dequeue(ctx context.Context, work Work) error
}

type Work struct {
	ID      string
	Type    string
	Payload []byte
	Status  string
}
