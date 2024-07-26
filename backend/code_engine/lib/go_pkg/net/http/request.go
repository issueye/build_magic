package http

import (
	"io"
	"net/http"

	"github.com/dop251/goja"
	"github.com/issueye/build_magic/backend/code_engine/lib/go_pkg/net/url"
	"github.com/issueye/build_magic/backend/pkg/utils"
)

func NewRequest(runtime *goja.Runtime, r *http.Request) *goja.Object {
	o := runtime.NewObject()

	o.Set("getContentLength", r.ContentLength)
	o.Set("getMethod", r.Method)
	o.Set("getHost", r.Host)
	o.Set("getRemoteAddr", r.RemoteAddr)
	o.Set("getBody", r.Body)

	o.Set("getHeader", func() *goja.Object {
		return NewHeader(runtime, r.Header)
	})

	o.Set("getHeaders", map[string][]string(r.Header))
	o.Set("getUri", r.RequestURI)

	o.Set("getUrl", func(call goja.FunctionCall) goja.Value {
		return url.NewURL(runtime, r.URL)
	})

	o.Set("getForm", func() goja.Value {
		return url.NewValues(runtime, r.Form)
	})

	o.Set("formValue", r.FormValue)

	o.Set("formFile", func(key string) goja.Value {
		file, fileHeader, err := r.FormFile(key)
		if err != nil {
			return nil
		}

		return runtime.ToValue(map[string]interface{}{
			"file":   NewMultipartFile(runtime, file),
			"name":   fileHeader.Filename,
			"header": map[string][]string(fileHeader.Header),
		})
	})

	o.Set("userAgent", r.UserAgent())
	o.Set("parseForm", r.ParseForm)
	o.Set("parseMultipartForm", r.ParseMultipartForm)

	o.Set("cookie", func(name string) goja.Value {
		c, err := r.Cookie(name)
		if err != nil {
			return nil
		}
		return NewCookie(runtime, c)
	})

	o.Set("cookies", func(call goja.FunctionCall) goja.Value {
		return runtime.ToValue(r.Cookies())
	})

	o.Set("getRawBody", func() (bts []byte, err error) {
		return io.ReadAll(r.Body)
	})

	o.Set("getBodyString", func(call goja.FunctionCall) (value string, err error) {
		bts, err := io.ReadAll(r.Body)
		if err != nil {
			return
		}

		return utils.Bytes2String(bts), nil
	})

	return o
}
