package controller

import (
	"context"

	"github.com/issueye/build_magic/backend/common/model"
	commonService "github.com/issueye/build_magic/backend/common/service"
	"github.com/issueye/build_magic/backend/logic"
	"github.com/issueye/build_magic/backend/repository"
	"github.com/issueye/build_magic/backend/service"
)

var buildProject *BuildProject

// 数据模型
type BuildProject struct {
	Ctx context.Context
}

// NewApp creates a new App application struct
func GetBuildProject() *BuildProject {
	if buildProject == nil {
		buildProject = &BuildProject{}
	}

	return buildProject
}

// 创建数据模型
func (app *BuildProject) Create(data *repository.CreateBuildProject) error {
	tplc := new(logic.BuildProject)
	return tplc.Create(data)
}

func (app *BuildProject) List(condition repository.QryBuildProject, page, size int) ([]*model.BuildProject, error) {
	srv := commonService.NewService(&service.BuildProjectService{})
	qry := model.NewPage(condition)
	qry.PageNum = int64(page)
	qry.PageSize = int64(size)

	return srv.List(qry)
}

// 创建数据模型
func (app *BuildProject) Delete(id string) error {
	return new(logic.BuildProject).Delete(id)
}

// Modify
// 修改代码模板信息
func (app *BuildProject) Modify(data *repository.ModifyBuildProject) error {
	tplc := new(logic.BuildProject)
	return tplc.Modify(data)
}
