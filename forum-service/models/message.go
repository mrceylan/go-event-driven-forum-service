package models

import "time"

type Message struct {
	Id         string
	TopicId    string
	Message    string
	CreateDate time.Time
	CreatedBy  string
}
