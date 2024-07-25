package logic

import (
	"github.com/issueye/build_magic/backend/common/model"
	commonService "github.com/issueye/build_magic/backend/common/service"
	"github.com/issueye/build_magic/backend/repository"
	"github.com/issueye/build_magic/backend/service"
)

// BuildProject
// 构建项目
type BuildProject struct{}

// Create
// 添加[构建项目]信息
func (owner *BuildProject) Create(data *repository.CreateBuildProject) error {
	srv := commonService.NewService(&service.BuildProjectService{})
	return srv.Create(data)
}

// Modify
// 修改[构建项目]信息
func (owner *BuildProject) Modify(data *repository.ModifyBuildProject) error {
	srv := commonService.NewService(&service.BuildProjectService{})
	return srv.Modify(data)
}

// Delete
// 删除[构建项目]信息
func (owner *BuildProject) Delete(id string) error {
	srv := commonService.NewService(&service.BuildProjectService{})
	return srv.Delete(id)
}

// GetById
// 通过ID查询[构建项目]数据
func (owner *BuildProject) GetById(id string) (*model.BuildProject, error) {
	srv := commonService.NewService(&service.BuildProjectService{})
	return srv.GetById(id)
}

// GetByName
// 通过名称[构建项目]数据
func (owner *BuildProject) GetByName(name string) (*model.BuildProject, error) {
	srv := commonService.NewService(&service.BuildProjectService{})
	return srv.GetByName(name)
}

// List
// 获取[构建项目]列表
func (owner *BuildProject) List(condition repository.QryBuildProject, page, size int) ([]*model.BuildProject, error) {
	srv := commonService.NewService(&service.BuildProjectService{})

	qry := model.NewPage(condition)
	qry.PageNum = int64(page)
	qry.PageSize = int64(size)

	return srv.List(qry)
}
