package initialize

import (
	"fmt"
	"path/filepath"

	"github.com/issueye/build_magic/backend/common/model"
	commonService "github.com/issueye/build_magic/backend/common/service"
	"github.com/issueye/build_magic/backend/global"
	"github.com/issueye/build_magic/backend/logic"
	"github.com/issueye/build_magic/backend/pkg/db"
)

// 初始化其他数据
func InitData() {
	path := filepath.Join("runtime", "data", "data.db")
	var err error
	commonService.DB, err = db.InitSqlite(path, global.Logger)
	if err != nil {
		panic(fmt.Sprintf("数据库初始化失败 %s", err.Error()))
	}

	// 初始化表
	err = commonService.DB.AutoMigrate(
		&model.UserInfo{},      // 用户
		&model.UserGroupInfo{}, // 用户组
		&model.DataModel{},     // 数据模型
		&model.ModelInfo{},     // 模型信息
		&model.CodeTemplate{},  // 代码模板
		&model.DataSource{},    // 数据源
		&model.Scheme{},        // 代码模板方法
		&model.Variable{},      // 变量
		&model.BuildProject{},  // 构建项目
	)

	if err != nil {
		panic(fmt.Errorf("初始化表失败 %s", err.Error()))
	}

	// 初始化
	tpLc := &logic.Template{}
	tpLc.InitScriptScheme()
}
