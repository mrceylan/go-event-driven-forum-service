package models

type PublishEvent struct {
	Topic string
	Event interface{}
}
