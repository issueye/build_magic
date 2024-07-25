package service

import (
	"github.com/issueye/build_magic/backend/common/model"
	"github.com/issueye/build_magic/backend/common/service"
	"github.com/issueye/build_magic/backend/repository"
	"gorm.io/gorm"
)

type Template struct {
	service.BaseService
}

func (owner *Template) SetBase(base service.BaseService) {
	owner.BaseService = base
}

// 创建数据
func (owner *Template) Create(data *repository.CreateTemplate) error {
	info := model.NewCodeTemplate(&data.CodeTemplateBase)
	return owner.GetDB().Model(&model.CodeTemplate{}).Create(info).Error
}

// 通过ID 获取数据
func (owner *Template) Get(id string) (*model.CodeTemplate, error) {
	info := &model.CodeTemplate{}
	return info, owner.GetDB().Model(&model.CodeTemplate{}).Where("ID = ?", id).First(info).Error
}

// 通过ID 获取数据
func (owner *Template) GetBySchemeCode(code string) (*model.CodeTemplate, error) {
	info := &model.CodeTemplate{}
	return info, owner.GetDB().Model(&model.CodeTemplate{}).Where("scheme_code = ?", code).First(info).Error
}

// 通过ID 获取数据
func (owner *Template) GetChildrenByCode(code string) (list []*model.CodeTemplate, err error) {
	list = make([]*model.CodeTemplate, 0)
	err = owner.GetDB().Model(&model.CodeTemplate{}).Where("scheme_parent_code = ?", code).Find(&list).Error
	return
}

// 通过ID 获取数据
func (owner *Template) Gets(ids []string) ([]*model.CodeTemplate, error) {
	info := make([]*model.CodeTemplate, 0)
	err := owner.GetDB().Model(&model.CodeTemplate{}).Where("ID in (?)", ids).Find(&info).Error
	return info, err
}

// 创建数据
func (owner *Template) Modify(data *repository.ModifyTemplate) error {
	updateMap := make(map[string]any)
	updateMap["title"] = data.Title
	updateMap["file_name"] = data.FileName
	updateMap["mark"] = data.Mark
	return owner.GetDB().Model(&model.CodeTemplate{}).Where("ID = ?", data.ID).Updates(&updateMap).Error
}

// 创建数据
func (owner *Template) Delete(id string) error {
	return owner.GetDB().Delete(&model.CodeTemplate{}, id).Error
}

// 创建数据
func (owner *Template) DeleteBySchemeCode(code string) error {
	return owner.GetDB().Where("scheme_code = ?", code).Delete(&model.CodeTemplate{}).Error
}

// 创建数据
func (owner *Template) DeleteByParentSchemeCode(code string) error {
	return owner.GetDB().Where("scheme_parent_code = ?", code).Delete(&model.CodeTemplate{}).Error
}

// 获取数据列表
func (owner *Template) List(data *model.Page[repository.QryTemplate]) ([]*model.CodeTemplate, error) {
	list := make([]*model.CodeTemplate, 0)
	err := service.Filter(owner, model.CodeTemplate{}.TableName(), data, &list, func(db *gorm.DB) (*gorm.DB, error) {

		if data.Condition.Condition != "" {
			db = db.Where("(title like ? file_name like ? file_type like ? or mark like ?)", "%"+data.Condition.Condition+"%", "%"+data.Condition.Condition+"%")
		}

		if data.Condition.ParentCode != "" {
			db.Where("scheme_code = ?", data.Condition.ParentCode).Or("scheme_parent_code = ?", data.Condition.ParentCode)
		}

		if len(data.Condition.Ids) > 0 {
			db = db.Where("id in (?)", data.Condition.Ids)
		}

		return db, nil
	})

	return list, err
}
