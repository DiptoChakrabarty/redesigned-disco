package database

import (
	"time"

	bolt "go.etcd.io/bbolt"
)

var defaultBucket = []byte("default")

var DbConn *bolt.DB

type Task struct {
	Key   int
	Value string
}

func Init(DbPath string) error {
	var err error
	DbConn, err = bolt.Open(DbPath, 0600, &bolt.Options{Timeout: 3 * time.Second})
	if err != nil {
		return err
	}
	return CreateBucket(defaultBucket)
}
