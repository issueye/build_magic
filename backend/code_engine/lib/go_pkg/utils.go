package gopkg

import (
	"encoding/base64"
	"fmt"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"github.com/google/uuid"
	"github.com/issueye/build_magic/backend/pkg/utils"
)

func InitUtils() {
	require.RegisterNativeModule("go/utils", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("print", fmt.Print)
		o.Set("panic", func(v any) {
			panic(v)
		})

		o.Set("toString", utils.Bytes2String)
		o.Set("toBase64", func(value any) any {
			switch data := value.(type) {
			case []byte:
				return base64.StdEncoding.EncodeToString(data)
			case string:
				return base64.StdEncoding.EncodeToString([]byte(data))
			default:
				return value
			}
		})

		

		o.Set("uuid", uuid.NewString)
	})
}
