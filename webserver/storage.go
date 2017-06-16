package goRank

import (
	"log"

	"github.com/gocql/gocql"
)

func Save(event Event) {
	session := getSession()
	defer session.Close()

	// insert a user
	if err := session.Query("INSERT INTO click.clicks (search, item, timestamp) VALUES (?, ?, now())", event.Search, event.Item).Exec(); err != nil {
		log.Fatal(err)
	}
}

func FindEventsForSearch(search string) (events []Event, err error) {
	session := getSession()
	defer session.Close()

	iter := session.Query("SELECT * FROM click.clicks WHERE search = ?", search).Iter()
	print("queried!\n")
	var searchresult, item string
	var timestamp gocql.UUID
	for iter.Scan(&searchresult, &item, &timestamp) {
		events = append(events, Event{searchresult, item, timestamp.Time()})
	}
	err = iter.Close()

	return
}

func InitStorage() {
	session := getSession()
	defer session.Close()

	if err := session.Query("CREATE KEYSPACE IF NOT EXISTS click WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}").Exec(); err != nil {
		log.Fatal(err)
	}

	if err := session.Query("CREATE TABLE IF NOT EXISTS click.clicks ( search text, item text, timestamp timeuuid, PRIMARY KEY (search, item, timestamp) )").Exec(); err != nil {
		log.Fatal(err)
	}
}

func getSession() *gocql.Session {
	// connect to the cluster
	cluster := gocql.NewCluster("cassandra")
	// workaround for issue connecting to dockerized cassandra
	// https://github.com/gocql/gocql/issues/575
	cluster.DisableInitialHostLookup = true
	session, err := cluster.CreateSession()

	if err != nil {
		log.Fatal(err)
	}
	return session
}
