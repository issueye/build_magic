package code_engine

import (
	"os"
	"path/filepath"

	"github.com/dop251/goja_nodejs/require"
)

func (vm *JsVM) sourceLoader(_ string) func(string) ([]byte, error) {
	return func(filename string) ([]byte, error) {
		// fmt.Println("sourceLoader -> filename", filename)
		data, err := os.ReadFile(filename)
		if err != nil {
			for _, v := range []string{filename, filename + ".ts"} {
				data, err = os.ReadFile(v)
				if err == nil {
					// fmt.Println("v", v)
					break
				}
			}
		}

		if err != nil {
			return nil, require.ModuleFileDoesNotExistError
		}
		ext := filepath.Ext(filename)
		if ext == ".ts" || ext == "" {
			data, err = vm.Transpile(data)
		}

		return data, err
	}
}
