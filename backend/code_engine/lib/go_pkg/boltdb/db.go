package boltdb

import (
	"github.com/dop251/goja"
	bolt "go.etcd.io/bbolt"
)

type CallBackFunc = func(goja.FunctionCall) goja.Value

func NewBoltDb(rt *goja.Runtime, db *bolt.DB) *goja.Object {
	o := rt.NewObject()

	// createBucket
	// 创建 bucket
	o.Set("createBucket", func(name string) error {
		err := db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(name))
			if err != nil {
				return err
			}
			return nil
		})

		return err
	})

	// view
	// 查询数据
	o.Set("view", func(name string, callback func(*goja.Object) error) error {
		return db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(name))
			return callback(NewBucket(rt, b))
		})
	})

	// update
	// 更新数据
	o.Set("update", func(name string, callback func(*goja.Object) error) error {
		return db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(name))
			return callback(NewBucket(rt, b))
		})
	})

	return o
}
