package models

import "time"

type Topic struct {
	Id         string
	Header     string
	CreateDate time.Time
	CreatedBy  string
}
