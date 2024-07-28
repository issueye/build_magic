package url

import (
	"net/url"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

func NewURL(runtime *goja.Runtime, u *url.URL) *goja.Object {
	o := runtime.NewObject()
	o.Set("getForceQuery", u.ForceQuery)
	o.Set("getFragment", u.Fragment)
	o.Set("getHost", u.Host)
	o.Set("getOpaque", u.Opaque)
	o.Set("getPath", u.Path)
	o.Set("getRawPath", u.RawPath)
	o.Set("getRawQuery", u.RawQuery)
	o.Set("getScheme", u.Scheme)
	o.Set("getPort", u.Port)

	return o
}

func InitUrl() {
	require.RegisterNativeModule("go/url", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("parse", url.Parse)
		o.Set("queryEscape", url.QueryEscape)
		o.Set("queryUnescape", url.QueryUnescape)
		o.Set("parseQuery", url.ParseQuery)

		o.Set("parseRequestURI", func(rawurl string) goja.Value {
			u, err := url.ParseRequestURI(rawurl)
			if err != nil {
				return nil
			}
			return NewURL(runtime, u)
		})

		o.Set("newValues", func() goja.Value {
			return NewValues(runtime, make(url.Values))
		})
	})
}
