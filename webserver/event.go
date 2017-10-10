package webserver

import "time"

type Event struct {
	Search 		string `json:"search"`
	Item   		string `json:"item"`
	Timestamp 	time.Time `json:"timestamp"`
}
