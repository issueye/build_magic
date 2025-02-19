package boltdb

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	bolt "go.etcd.io/bbolt"
)

func InitBolt() {
	require.RegisterNativeModule("db/bolt", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)

		// 打开数据库
		o.Set("open", func(path string) goja.Value {
			db, err := bolt.Open(path, 0666, nil)
			if err != nil {
				return nil
			}

			return NewBoltDb(runtime, db)
		})

	})
}

func RegisterNativeBolt(rt *goja.Runtime, name string, db *bolt.DB) {
	rt.Set(name, NewBoltDb(rt, db))
}
