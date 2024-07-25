package service

import (
	"github.com/issueye/build_magic/backend/common/model"
	"github.com/issueye/build_magic/backend/common/service"
	"github.com/issueye/build_magic/backend/repository"
	"gorm.io/gorm"
)

type Scheme struct {
	service.BaseService
}

func (owner *Scheme) SetBase(base service.BaseService) {
	owner.BaseService = base
}

// 创建数据
func (owner *Scheme) Create(data *repository.CreateScheme) error {
	info := model.NewScheme(&data.SchemeBase)
	return owner.GetDB().Model(&model.Scheme{}).Create(info).Error
}

// 通过ID 获取数据
func (owner *Scheme) Get(id string) (*model.Scheme, error) {
	info := &model.Scheme{}
	err := owner.GetDB().Model(&model.Scheme{}).Where("ID = ?", id).Find(info).Error
	return info, err
}

// ProgrammeList
func (owner *Scheme) ProgrammeList() ([]*model.Scheme, error) {
	info := make([]*model.Scheme, 0)
	err := owner.GetDB().Model(&model.Scheme{}).Where("level = ? and parent_code = ?", 1, "002").Find(&info).Error
	return info, err
}

// 通过ID 获取数据
func (owner *Scheme) GetChildrenByCode(code string) (list []*model.Scheme, err error) {
	list = make([]*model.Scheme, 0)
	err = owner.GetDB().Model(&model.Scheme{}).Where("parent_code = ?", code).Find(&list).Error
	return list, err
}

// 通过ID 获取数据
func (owner *Scheme) GetByCode(code string) (*model.Scheme, error) {
	info := &model.Scheme{}
	err := owner.GetDB().Model(&model.Scheme{}).Where("code = ?", code).Find(info).Error
	return info, err
}

// 通过ID 获取数据
func (owner *Scheme) Gets(ids []string) ([]*model.Scheme, error) {
	info := make([]*model.Scheme, 0)
	err := owner.GetDB().Model(&model.Scheme{}).Where("ID in (?)", ids).Find(&info).Error
	return info, err
}

// 创建数据
func (owner *Scheme) Modify(data *repository.ModifyScheme) error {
	updateMap := make(map[string]any)
	updateMap["title"] = data.Title
	updateMap["icon"] = data.Icon
	return owner.GetDB().Model(&model.Scheme{}).Where("ID = ?", data.ID).Updates(&updateMap).Error
}

// 创建数据
func (owner *Scheme) ModifyByCode(data *repository.ModifyScheme) error {
	updateMap := make(map[string]any)
	updateMap["title"] = data.Title
	updateMap["icon"] = data.Icon
	return owner.GetDB().Model(&model.Scheme{}).Where("code = ?", data.Code).Updates(&updateMap).Error
}

// 创建数据
func (owner *Scheme) Delete(id string) error {
	return owner.GetDB().Delete(&model.Scheme{}, id).Error
}

// 创建数据
func (owner *Scheme) DeleteByCode(code string) error {
	return owner.GetDB().Where("code = ?", code).Delete(&model.Scheme{}).Error
}

// 创建数据
func (owner *Scheme) DeleteByParentCode(code string) error {
	return owner.GetDB().Where("parent_code = ?", code).Delete(&model.Scheme{}).Error
}

// 获取数据列表
func (owner *Scheme) List(data *model.Page[repository.QryScheme]) ([]*model.Scheme, error) {
	list := make([]*model.Scheme, 0)
	err := service.Filter(owner, model.Scheme{}.TableName(), data, &list, func(db *gorm.DB) (*gorm.DB, error) {

		if data.Condition.Condition != "" {
			db = db.Where("(title like ? file_name like ? file_type like ? or mark like ?)", "%"+data.Condition.Condition+"%", "%"+data.Condition.Condition+"%")
		}

		if len(data.Condition.Ids) > 0 {
			db = db.Where("id in (?)", data.Condition.Ids)
		}

		return db, nil
	})

	return list, err
}

// 获取指定节点的子节点
func (owner *Scheme) GetNodeByCode(code []string, parentCode string) (list []*model.Scheme, err error) {
	list = make([]*model.Scheme, 0)

	sql := `SELECT 1 as num,* FROM scheme WHERE nodes || ',' LIKE '%001,%' UNION SELECT 2 as num,* FROM scheme WHERE code in (?) or parent_code = ?;`
	err = owner.GetDB().Raw(sql, code, parentCode).Scan(&list).Error
	return
}

// 获取指定节点的子节点
func (owner *Scheme) GetCommonNode() (list []*model.Scheme, err error) {
	list = make([]*model.Scheme, 0)

	sql := `SELECT * FROM scheme WHERE nodes || ',' LIKE '%001,%';`
	err = owner.GetDB().Raw(sql).Scan(&list).Error
	return
}
