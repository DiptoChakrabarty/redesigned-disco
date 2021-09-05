package database

import (
	bolt "go.etcd.io/bbolt"
)

func CreateBucket(bucketName []byte) error {
	return DbConn.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	})
}
