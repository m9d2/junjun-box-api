package model

import "time"

type Event struct {
	Event string    `json:"event"`
	Id    int64     `json:"id"`
	Data  string    `json:"data"`
	Time  time.Time `json:"time"`
}
