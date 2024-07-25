package initialize

import (
	"path/filepath"

	"github.com/issueye/build_magic/backend/code_engine"
	"github.com/issueye/build_magic/backend/global"
)

func InitJSVM() {
	logPath := filepath.Join("runtime", "logs")
	global.JsVMCore = code_engine.NewCore(
		code_engine.OptionLog(logPath, global.Logger.Named("code_engine")),
	)

	globalPath := filepath.Join("runtime", "static", "code")
	global.JsVMCore.SetGlobalPath(globalPath)
}
