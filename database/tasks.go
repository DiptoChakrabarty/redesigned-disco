package database

import (
	"go-scheduler/utils"
	bolt "go.etcd.io/bbolt"
)

func AddTask(task string) (int, error) {
	var id int
	err := DbConn.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(defaultBucket)
		id64, _ := bkt.NextSequence()
		id = int(id64)
		key := utils.ConvertIntToBytes(id)
		return bkt.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := DbConn.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(defaultBucket)
		cur := bkt.Cursor()
		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			tasks = append(tasks, Task{
				Key:   utils.ConvertBytesToInt(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTasks(key int) error {
	err := DbConn.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(defaultBucket)
		return bkt.Delete(utils.ConvertIntToBytes(key))
	})
	if err != nil {
		return err
	}
	return nil
}
