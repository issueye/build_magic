package initialize

import (
	"path/filepath"

	"github.com/issueye/build_magic/backend/config"
	"github.com/issueye/build_magic/backend/global"
	"github.com/issueye/build_magic/backend/pkg/logger"
)

func InitLog() {
	logConf := new(logger.Config)
	logConf.Path = filepath.Join("runtime", "logs")
	logConf.MaxSize = config.GetParam(config.CfgLogMaxSize, "10").Int()
	logConf.MaxBackups = config.GetParam(config.CfgLogMaxBackups, "10").Int()
	logConf.MaxAge = config.GetParam(config.CfgLogMaxAge, "10").Int()
	logConf.Compress = config.GetParam(config.CfgLogCompress, "true").Bool()
	logConf.Level = config.GetParam(config.CfgLogLevel, "-1").Int()
	logConf.Mode = logger.LOM_DEBUG
	global.Log, global.Logger = logger.InitLogger(logConf)
}
