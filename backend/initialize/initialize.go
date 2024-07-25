package initialize

import (
	"context"

	"github.com/issueye/build_magic/backend/config"
	"github.com/issueye/build_magic/backend/global"
)

func Initialize() {
	ctx := context.Background()
	// 初始化运行文件
	InitRuntime()
	// 配置参数
	config.InitConfig()
	// 日志
	InitLog()
	// 数据
	InitData()
	// 初始化js虚拟机
	InitJSVM()

	global.Logger.Sync()
	// 关闭监听
	ctx.Done()
}

var (
	AppName string
	Branch  string
	Commit  string
	Date    string
	Version string
)
