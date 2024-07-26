package boltdb

import (
	"github.com/dop251/goja"
	bolt "go.etcd.io/bbolt"
)

func NewBucket(runtime *goja.Runtime, bucket *bolt.Bucket) *goja.Object {
	o := runtime.NewObject()
	// create
	o.Set("create", func(name string) goja.Value {
		b, err := bucket.CreateBucket([]byte(name))
		if err != nil {
			return nil
		}

		return NewBucket(runtime, b)
	})

	// get
	o.Set("get", func(key string) string {
		b := bucket.Get([]byte(key))
		return string(b)
	})

	// put
	o.Set("put", func(key string, value string) error {
		err := bucket.Put([]byte(key), []byte(value))
		if err != nil {
			return err
		}

		return nil
	})

	// delete
	o.Set("delete", func(key string) error {
		return bucket.Delete([]byte(key))
	})

	// forEach
	o.Set("foreach", func(callback func(key, value string) error) goja.Value {
		bucket.ForEach(func(k, v []byte) error {
			return callback(string(k), string(v))
		})
		return nil
	})

	// forEachBucket
	o.Set("foreachBucket", func(callback func(key string) error) goja.Value {
		bucket.ForEachBucket(func(k []byte) error {
			return callback(string(k))
		})
		return nil
	})

	// cursor
	o.Set("cursor", func() goja.Value {
		corsor := bucket.Cursor()
		return NewCursor(runtime, corsor)
	})

	return o
}

func NewCursor(runtime *goja.Runtime, cursor *bolt.Cursor) *goja.Object {
	o := runtime.NewObject()

	// first
	o.Set("first", func() map[string]string {
		key, value := cursor.First()
		return map[string]string{
			string(key): string(value),
		}
	})

	// last
	o.Set("last", func() map[string]string {
		key, value := cursor.First()
		return map[string]string{
			string(key): string(value),
		}
	})

	// next
	o.Set("next", func() map[string]string {
		key, value := cursor.Next()
		return map[string]string{
			string(key): string(value),
		}
	})

	// seek
	o.Set("seek", func(prefix string) map[string]string {
		key, value := cursor.Seek([]byte(prefix))
		return map[string]string{
			string(key): string(value),
		}
	})

	// prev
	o.Set("prev", func() map[string]string {
		key, value := cursor.Prev()
		return map[string]string{
			string(key): string(value),
		}
	})

	return o
}
