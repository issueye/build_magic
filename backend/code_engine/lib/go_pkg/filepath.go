package gopkg

import (
	"path/filepath"

	js "github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

func InitFilepath() {
	require.RegisterNativeModule("go/filepath", func(runtime *js.Runtime, module *js.Object) {
		o := module.Get("exports").(*js.Object)
		o.Set("abs", filepath.Abs)
		o.Set("join", filepath.Join)
		o.Set("ext", filepath.Ext)
	})
}
