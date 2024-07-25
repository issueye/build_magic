package code_engine

import (
	_ "embed"

	"github.com/issueye/build_magic/backend/pkg/utils"
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztype"
)

// ts v4.9.3.js

//go:embed ts.js
var tsTranspile []byte

func (vm *JsVM) TranspileFile(file string) ([]byte, error) {
	code, err := zfile.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return vm.Transpile(code)
}

func (vm *JsVM) Transpile(code []byte) ([]byte, error) {
	compiler := `{"strict":false,"target":"ES5","module":"CommonJS"}`

	s := `ts.transpileModule(atob("` + zstring.Base64EncodeString(zstring.Bytes2String(code)) + `"), {
		"compilerOptions": ` + compiler + `,
		})`

	res, err := vm.vm.RunString(s)
	if err != nil {
		return nil, err
	}

	return utils.String2Bytes(ztype.ToString(res.ToObject(vm.vm).Get("outputText").Export()) + ""), nil
}
