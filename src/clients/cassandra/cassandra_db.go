package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {
	// connect to Cassandra cluster: (default 127.0.0.1:9042)
	cluster := gocql.NewCluster("localhost:32769")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()

	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra connection successfully created")
	defer session.Close()
}

func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
