package code_engine

import (
	"net/http"

	"github.com/dop251/goja"
	gopkg "github.com/issueye/build_magic/backend/code_engine/lib/go_pkg"
	"github.com/issueye/build_magic/backend/code_engine/lib/go_pkg/boltdb"
	libhttp "github.com/issueye/build_magic/backend/code_engine/lib/go_pkg/net/http"
	bolt "go.etcd.io/bbolt"
)

func init() {
	gopkg.InitStrings()  // 字符串
	gopkg.InitFmt()      // fmt
	gopkg.InitCrypto()   // 加密
	gopkg.InitFilepath() // 文件路径
	gopkg.InitUtils()
	gopkg.InitCmd()
	boltdb.InitBolt()

}

func NewRequest(runtime *goja.Runtime, r *http.Request) *goja.Object {
	return libhttp.NewRequest(runtime, r)
}

func NewResponse(runtime *goja.Runtime, w http.ResponseWriter) *goja.Object {
	return libhttp.NewResponse(runtime, w)
}

func RegisterNativeBolt(rt *goja.Runtime, name string, db *bolt.DB) {
	boltdb.RegisterNativeBolt(rt, name, db)
}
