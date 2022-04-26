package db

import (
	"github.com/gocql/gocql"
)

var Session *gocql.Session

func Init() {
	var err error
	cluster := gocql.NewCluster("localhost:9042")
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	Session.Query("CREATE KEYSPACE IF NOT EXISTS sleep_centre WITH REPLICATION = {'class' : 'NetworkTopologyStrategy', 'replication_factor': 3}")
}
