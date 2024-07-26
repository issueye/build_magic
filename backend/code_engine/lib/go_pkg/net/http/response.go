package http

import (
	"fmt"
	"net/http"

	"github.com/dop251/goja"
	"github.com/issueye/build_magic/backend/pkg/utils"
)

func NewResponse(runtime *goja.Runtime, w http.ResponseWriter) *goja.Object {
	o := runtime.NewObject()
	o.Set("header", func(call goja.FunctionCall) goja.Value {
		return NewHeader(runtime, w.Header())
	})

	o.Set("write", func(value any) (int, error) {
		var data []byte
		switch t := value.(type) {
		case []interface{}:
			data = make([]byte, len(t))
			for i, v := range t {
				if val, ok := v.(int64); ok {
					if val >= 0 && val <= 255 {
						data[i] = byte(val)
						continue
					}
				}
				panic(runtime.NewTypeError("can not convert to byte `data[%d]:%T`", i, v))
			}
		case []byte:
			data = t
		case string:
			data = utils.String2Bytes(t)
		default:
			return 0, fmt.Errorf("can not convert to byte `%T`", t)
		}

		return w.Write(data)
	})

	o.Set("writeHeader", w.WriteHeader)
	o.Set("setCookie", func(name string, value string, path string, maxAge int, httpOnly bool) {
		cookie := &http.Cookie{
			Name:     name,
			Value:    value,
			Path:     path,
			MaxAge:   maxAge,
			HttpOnly: httpOnly,
		}
		http.SetCookie(w, cookie)
	})

	o.Set("nativeType", w)

	return o
}
