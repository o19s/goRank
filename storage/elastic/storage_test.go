package elastic

import (
	"testing"
	"goRank/models"
	"time"
	"os"
)

var e struct {
	Connection Elastic
}

func TestMain(m *testing.M) {
	e.Connection.Connect()
	os.Exit(m.Run())

}

func TestInit(t *testing.T) {
	e.Connection.Version()
}

func TestSave(t *testing.T) {
	test_event := models.Event{
		Search: "test",
		Item: "Test1",
		Timestamp: time.Now(),
	}

	e.Connection.Save(test_event)
}