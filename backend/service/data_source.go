package service

import (
	"github.com/issueye/build_magic/backend/common/model"
	"github.com/issueye/build_magic/backend/common/service"
	"github.com/issueye/build_magic/backend/repository"
	"gorm.io/gorm"
)

type DataSource struct {
	service.BaseService
}

func (owner *DataSource) SetBase(base service.BaseService) {
	owner.BaseService = base
}

// 创建数据
func (owner *DataSource) Create(data *repository.CreateDataSource) error {
	info := model.NewDataSource(&model.DataSourceBase{
		Title:    data.Title,
		Host:     data.Host,
		Port:     data.Port,
		UserName: data.UserName,
		Password: data.Password,
		DbType:   data.DbType,
		Database: data.Database,
		Schema:   data.Schema,
		Path:     data.Path,
	})

	return owner.GetDB().Model(&model.DataSource{}).Create(info).Error
}

// 创建数据
func (owner *DataSource) Modify(data *repository.UpdateDataSource) error {
	updateMap := make(map[string]any)
	updateMap["title"] = data.Title
	updateMap["host"] = data.Host
	updateMap["port"] = data.Port
	updateMap["user_name"] = data.UserName
	updateMap["password"] = data.Password
	updateMap["db_type"] = data.DbType
	updateMap["database"] = data.Database
	updateMap["schema"] = data.Schema
	updateMap["path"] = data.Path

	return owner.GetDB().Model(&model.DataSource{}).Where("ID = ?", data.ID).Updates(&updateMap).Error
}

// 创建数据
func (owner *DataSource) Delete(id string) error {
	return owner.GetDB().Delete(&model.DataSource{}, id).Error
}

func (owner *DataSource) GetById(id string) (*model.DataSource, error) {
	info := &model.DataSource{}
	return info, owner.GetDB().Model(&model.DataSource{}).Where("ID = ?", id).First(info).Error
}

// 获取数据列表
func (owner *DataSource) List(data *model.Page[repository.QueryDataSource]) ([]*model.DataSource, error) {
	list := make([]*model.DataSource, 0)
	err := service.Filter(owner, model.DataSource{}.TableName(), data, &list, func(db *gorm.DB) (*gorm.DB, error) {

		if data.Condition.Condition != "" {
			db = db.Where("(title like ? or mark like ?)", "%"+data.Condition.Condition+"%", "%"+data.Condition.Condition+"%")
		}

		return db, nil
	})

	return list, err
}
