package database

import (
	"go-scheduler/utils"
	bolt "go.etcd.io/bbolt"
)

var bucketlist = []byte("bucketlist")

func CreateBucket(bucketName []byte) error {
	return DbConn.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		err = AddBucketToBucketList(bucketName)
		return err
	})
}

func AddBucketToBucketList(bucketName []byte) error {
	err := DbConn.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(bucketlist)
		id64, _ := bkt.NextSequence()
		id := int(id64)
		key := utils.ConvertIntToBytes(id)
		return bkt.Put(key, []byte(bucketName))
	})
	if err != nil {
		return err
	}
	return nil
}

func ViewAllBuckets() ([]string, error) {
	var buckets []string
	err := DbConn.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(bucketlist)
		cur := bkt.Cursor()

		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			buckets = append(buckets, string(v))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return buckets, nil
}
