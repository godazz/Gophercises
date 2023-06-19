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
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func connectDB() *bolt.DB {
	db, err := bolt.Open(dbFileName, 0600, nil)
	if err != nil {
		Exitf("%v", err)
	}
	return db
}

func createBucket(tx *bolt.Tx) *bolt.Bucket {
	b, err := tx.CreateBucketIfNotExists([]byte(tasksBucketName))
	if err != nil {
		Exitf("%v", err)
	}
	return b
}

func CreateTask(task *Task) error {
	db := connectDB()
	defer db.Close()
	return db.Update(func(tx *bolt.Tx) error {

		b := createBucket(tx)
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

func ListTasks(isCompleted bool) ([]*Task, error) {

	db := connectDB()
	defer db.Close()

	var tasks []*Task
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(tasksBucketName))

		b.ForEach(func(_, v []byte) error {
			var task Task
			if err := json.Unmarshal(v, &task); err != nil {
				return err
			}
			if task.Completed != isCompleted {
				return nil
			}
			tasks = append(tasks, &task)
			return nil
		})
		return nil
	})

	return tasks, nil
}

func MarkTaskAsCompleted(taskID int) (string, error) {

	db := connectDB()
	defer db.Close()

	var task Task
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(tasksBucketName))
		taskRecord := b.Get(Itob(taskID))

		if taskRecord == nil {
			Exitf("task not found with Id=%d", taskID)
		}

		if err := json.Unmarshal(taskRecord, &task); err != nil {
			return err
		}

		task.Completed = true
		taskRecord, err := json.Marshal(&task)
		if err != nil {
			return err
		}
		return b.Put(Itob(task.ID), taskRecord)
	})
	return task.Title, nil
}

func DeleteTask(taskID int) (string, error) {
	db := connectDB()
	defer db.Close()

	var task Task
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(tasksBucketName))
		taskRecord := b.Get(Itob(taskID))

		if taskRecord == nil {
			Exitf("task not found with Id=%d", taskID)
		}

		if err := json.Unmarshal(taskRecord, &task); err != nil {
			return err
		}
		return b.Delete(Itob(taskID))
	})
	return task.Title, nil
}
