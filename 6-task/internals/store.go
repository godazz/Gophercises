package internals

import (
	"encoding/json"

	"github.com/boltdb/bolt"
)

const (
	dbFileName      = "tasks.db"
	tasksBucketName = "tasks"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

func connectDB() *bolt.DB {
	db, err := bolt.Open(dbFileName, 0600, nil)
	if err != nil {
		Exitf("%v", err)
	}
	return db
}

func CreateTask(task *Task) error {
	db := connectDB()
	defer db.Close()
	return db.Update(func(tx *bolt.Tx) error {

		b, err := tx.CreateBucketIfNotExists([]byte(tasksBucketName))
		if err != nil {
			return err
		}
		id, err := b.NextSequence()
		if err != nil {
			return err
		}
		task.ID = int(id)

		buf, err := json.Marshal(&task)
		if err != nil {
			return err
		}

		return b.Put(Itob(task.ID), buf)
	})
}

func ListTasks() (*[]Task, error) {
	return nil, nil
}

func MarkTaskAsCompleted(taskID int) error {
	return nil
}
