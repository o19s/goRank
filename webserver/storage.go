package goRank

import (
	"github.com/gocql/gocql"
	"log"
)

func Save(event Event) {
	session := getSession()
	defer session.Close()

	// insert a user
	if err := session.Query("INSERT INTO clicks (search, item, timestamp) VALUES (?, ?, now())", event.Search, event.Item).Exec(); err != nil {
		log.Fatal(err)
	}
}

func FindEventsForSearch(search string) (events []Event, err error){
	session := getSession()
	defer session.Close()

	iter := session.Query("SELECT * FROM clicks WHERE search = ?", search).Iter()
	var searchresult, item string
	var timestamp gocql.UUID
	for iter.Scan(&searchresult, &item, &timestamp) {
		events = append(events, Event{searchresult, item, timestamp.Time()})
	}
	err = iter.Close()

	return
}

func getSession() *gocql.Session {
	// connect to the cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "click"
	session, _ := cluster.CreateSession()

	return session
}
