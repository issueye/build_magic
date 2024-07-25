package repository

import (
	"github.com/issueye/build_magic/backend/common/model"
)

// CreateBuildProject
// 创建[构建项目]
type CreateBuildProject struct {
    model.BuildProjectBase
}

// ModifyBuildProject
// 修改[构建项目]
type ModifyBuildProject struct {
    model.BuildProject
}

// QryBuildProject
// 查询[构建项目]
type QryBuildProject struct {
    Title string `json:"title" form:"title"` // 标题
    Name string `json:"name" form:"name"` // 名称
}