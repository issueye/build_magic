package http

import (
	"net/http"

	"github.com/dop251/goja"
)

func NewHeader(runtime *goja.Runtime, header http.Header) *goja.Object {
	o := runtime.NewObject()
	o.Set("add", header.Add)
	o.Set("del", header.Del)
	o.Set("get", header.Get)
	o.Set("set", header.Set)
	o.Set("clone", header.Clone)
	o.Set("gets", func(key string) []string {
		return header[key]
	})

	return o
}
