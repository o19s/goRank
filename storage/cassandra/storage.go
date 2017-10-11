package cassandra

import (
	"log"
	"github.com/gocql/gocql"
	"goRank/models"
)

type Cassandra struct {}

func (c Cassandra) Save(event models.Event) {
	session := c.getSession()
	defer session.Close()

	// insert a user
	if err := session.Query("INSERT INTO click.clicks (search, item, timestamp) VALUES (?, ?, now())", event.Search, event.Item).Exec(); err != nil {
		log.Fatal(err)
	}
}

func (c Cassandra)FindEventsForSearch(search string) (events []models.Event, err error) {
	session := c.getSession()
	defer session.Close()

	iter := session.Query("SELECT * FROM click.clicks WHERE search = ?", search).Iter()
	print("queried!\n")
	var searchresult, item string
	var timestamp gocql.UUID
	for iter.Scan(&searchresult, &item, &timestamp) {
		events = append(events, models.Event{searchresult, item, timestamp.Time()})
	}
	err = iter.Close()

	return
}

func (c Cassandra)InitStorage() {
	session := c.getSession()
	defer session.Close()

	if err := session.Query("CREATE KEYSPACE IF NOT EXISTS click WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}").Exec(); err != nil {
		panic(err)
	}

	if err := session.Query("CREATE TABLE IF NOT EXISTS click.clicks ( search text, item text, timestamp timeuuid, PRIMARY KEY (search, item, timestamp) )").Exec(); err != nil {
		panic(err)
	}
}

func (c Cassandra)getSession() *gocql.Session {
	// connect to the cluster
	cluster := gocql.NewCluster("cassandra")
	// workaround for issue connecting to dockerized cassandra
	// https://github.com/gocql/gocql/issues/575
	cluster.DisableInitialHostLookup = true
	session, err := cluster.CreateSession()

	if err != nil {
		panic(err)
	}
	return session
}
