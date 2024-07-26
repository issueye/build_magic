package model

import (
	"strconv"

	"github.com/issueye/build_magic/backend/pkg/utils"
)

// BuildProject
// BuildProject [构建项目]信息模型
type BuildProject struct {
	Base
	BuildProjectBase
}

// BuildProjectBase
// BuildProject [构建项目]基础信息模型
type BuildProjectBase struct {
	Title        string `gorm:"column:title;size:50;comment:标题;" json:"title"`                    // 标题
	Name         string `gorm:"column:name;size:50;comment:名称;" json:"name"`                      // 名称
	ProjectPath  string `gorm:"column:project_path;size:300;comment:项目路径;" json:"project_path"`   // 项目路径
	BuildType    int    `gorm:"column:build_type;type:int;comment:构建类型;" json:"build_type"`       // 构建类型
	Version      string `gorm:"column:version;size:200;comment:版本;" json:"version"`               // 版本
	BuildVersion int    `gorm:"column:build_version;type:int;comment:构建版本;" json:"build_version"` // 构建版本
	CodeID       string `gorm:"column:code_id;size:50;comment:代码仓库ID;" json:"code_id"`            // 代码仓库ID
	Mark         string `gorm:"column:mark;size:2000;comment:备注;" json:"mark"`                    // 备注
}

func (BuildProject) TableName() string {
	return "build_project"
}

func (BuildProject) Copy(data *BuildProjectBase) *BuildProject {
	info := new(BuildProject)
	info.BuildProjectBase = *data
	info.ID = strconv.FormatInt(utils.GenID(), 10)
	return info
}
