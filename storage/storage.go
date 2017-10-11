package storage

import (
	//"goRank/models"
	"goRank/models"
)

type Engine interface {
	Save(event models.Event)
	FindEventsForSearch(string) ([]models.Event, error)
	InitStorage()
}