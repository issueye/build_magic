package service

import (
	"fmt"

	"github.com/issueye/build_magic/backend/common/model"
	"github.com/issueye/build_magic/backend/common/service"
	"github.com/issueye/build_magic/backend/repository"
	"gorm.io/gorm"
)

// BuildProjectService
// 构建项目
type BuildProjectService struct {
	service.BaseService
}

func (owner *BuildProjectService) SetBase(base service.BaseService) {
	owner.BaseService = base
}

// Create
// 添加[构建项目]信息
func (owner *BuildProjectService) Create(req *repository.CreateBuildProject) error {
	table := model.BuildProject{}.Copy(&req.BuildProjectBase)
	return owner.GetDB().Create(table).Error
}

// Modify
// 修改[构建项目]信息
func (owner *BuildProjectService) Modify(req *repository.ModifyBuildProject) error {
	m := make(map[string]any)
	m["title"] = req.Title
	m["name"] = req.Name
	m["project_path"] = req.ProjectPath
	m["build_type"] = req.BuildType
	m["version"] = req.Version
	m["build_version"] = req.BuildVersion
	m["mark"] = req.Mark

	return owner.GetDB().Model(&model.BuildProject{}).Where("ID = ?", req.ID).Updates(&m).Error
}

// Delete
// 删除[构建项目]信息
func (owner *BuildProjectService) Delete(id string) error {
	data := new(model.BuildProject)
	return owner.GetDB().Model(&model.BuildProject{}).Where("ID = ?", id).Delete(data).Error
}

// GetById
// 通过ID查询[构建项目]数据
func (owner *BuildProjectService) GetById(id string) (*model.BuildProject, error) {
	data := new(model.BuildProject)
	err := owner.GetDB().Model(&model.BuildProject{}).Where("ID = ?", id).Find(data).Error
	return data, err
}

// GetById
// 通过ID查询[构建项目]数据
func (owner *BuildProjectService) GetByTpCode(id string) (*model.BuildProject, error) {
	data := new(model.BuildProject)
	err := owner.GetDB().Model(&model.BuildProject{}).Where("code_id = ?", id).Find(data).Error
	return data, err
}

// GetByName
// 通过名称[构建项目]数据
func (owner *BuildProjectService) GetByName(name string) (*model.BuildProject, error) {
	data := new(model.BuildProject)
	err := owner.GetDB().Model(&model.BuildProject{}).Where("name = ?", name).Find(data).Error
	return data, err
}

// List
// 获取[构建项目]列表
func (owner *BuildProjectService) List(req *model.Page[repository.QryBuildProject]) ([]*model.BuildProject, error) {
	list := make([]*model.BuildProject, 0)

	err := owner.DataFilter(model.BuildProject{}.TableName(), req, &list, func(db *gorm.DB) (*gorm.DB, error) {
		query := db.Order("ID")

		// 标题
		if req.Condition.Title != "" {
			query = query.Where("title like ?", fmt.Sprintf("%%%s%%", req.Condition.Title))
		}
		// 名称
		if req.Condition.Name != "" {
			query = query.Where("name = ?", req.Condition.Name)
		}

		return query, nil
	})

	return list, err
}
