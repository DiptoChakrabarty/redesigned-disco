package database

import (
	"go-scheduler/types"
	"go-scheduler/utils"
	bolt "go.etcd.io/bbolt"
)

func AddTask(task string, group []byte) (int, error) {
	var id int
	err := DbConn.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(group)
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

func GetAllTasks(group []byte) ([]types.Task, error) {
	var tasks []types.Task
	err := DbConn.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(group)
		cur := bkt.Cursor()
		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			tasks = append(tasks, types.Task{
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

func DeleteTasks(key int, group []byte) error {
	err := DbConn.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket(group)
		return bkt.Delete(utils.ConvertIntToBytes(key))
	})
	if err != nil {
		return err
	}
	return nil
}
