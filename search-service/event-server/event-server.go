package eventserver

import "context"

type IEventServer interface {
	StartConsumer(queue string, ch chan<- []byte, ctx context.Context)
	Disconnect() error
}
