package scylladb

import (
	"context"
	"fmt"
	"os"

	"github.com/gocql/gocql"
	"github.com/ramsfords/configs"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/migrate"
)

func New(conf configs.Config) gocqlx.Session {
	var cluster = gocql.NewCluster(conf.ScyllaDb.Public...)
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: conf.ScyllaDb.Username, Password: conf.ScyllaDb.Password}
	cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy(conf.ScyllaDb.Name)
	cluster.Keyspace = conf.ScyllaDb.Keyspace

	var session, err = cluster.CreateSession()

	if err != nil {
		panic("Failed to connect to cluster")
	}

	defer session.Close()

	var query = session.Query("SELECT * FROM system.clients")

	if rows, err := query.Iter().SliceMap(); err == nil {
		for _, row := range rows {
			fmt.Printf("%v\n", row)
		}
	} else {
		panic("Query error: " + err.Error())
	}
	gqlxSession, err := gocqlx.WrapSession(session, err)
	if err != nil {
		panic("Failed to connect to cluster")
	}
	files := os.DirFS("./scyllamigration")
	if err := migrate.FromFS(context.Background(), gqlxSession, files); err != nil {
		panic(err)
	}

	return gqlxSession
}
