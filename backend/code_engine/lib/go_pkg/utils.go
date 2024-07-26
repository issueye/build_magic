package gopkg

import (
	"encoding/base64"
	"fmt"

	js "github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"github.com/google/uuid"
	"github.com/issueye/build_magic/backend/pkg/utils"
)

func InitUtils() {
	require.RegisterNativeModule("std/utils", func(runtime *js.Runtime, module *js.Object) {
		o := module.Get("exports").(*js.Object)
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

		o.Set("md5", func(value any) any {
			switch data := value.(type) {
			case []byte:
				return utils.MD5V(utils.Bytes2String(data))
			case string:
				return utils.MD5V(data)
			default:
				return value
			}
		})

		o.Set("sha1", func(value any) any {
			switch data := value.(type) {
			case []byte:
				return utils.Sha1(utils.Bytes2String(data))
			case string:
				return utils.Sha1(data)
			default:
				return value
			}
		})

		o.Set("uuid", uuid.NewString)
	})
}
