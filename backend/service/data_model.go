package service

import (
	"github.com/issueye/build_magic/backend/common/model"
	"github.com/issueye/build_magic/backend/common/service"
	"github.com/issueye/build_magic/backend/repository"
	"gorm.io/gorm"
)

type DataModel struct {
	service.BaseService
}

func (owner *DataModel) SetBase(base service.BaseService) {
	owner.BaseService = base
}

// 创建数据
func (owner *DataModel) Create(data *repository.CreateDataModel) error {
	info := model.NewDataModel(&model.DataModelBase{
		Title:    data.Title,
		MakeType: data.MakeType,
		TBName:   data.TBName,
		SchemeId: data.SchemeId,
		Mark:     data.Mark,
	})

	return owner.GetDB().Model(&model.DataModel{}).Create(info).Error
}

// 创建数据
func (owner *DataModel) Modify(data *repository.RequestModifyDataModel) error {
	updateMap := make(map[string]any)
	updateMap["title"] = data.Title
	updateMap["make_type"] = data.MakeType
	updateMap["table_name"] = data.TBName
	updateMap["scheme_id"] = data.SchemeId
	updateMap["mark"] = data.Mark
	return owner.GetDB().Model(&model.DataModel{}).Where("ID = ?", data.ID).Updates(&updateMap).Error
}

// 创建数据
func (owner *DataModel) Delete(id string) error {
	return owner.GetDB().Delete(&model.DataModel{}, id).Error
}

func (owner *DataModel) GetById(id string) (*model.DataModel, error) {
	info := &model.DataModel{}
	return info, owner.GetDB().Model(&model.DataModel{}).Where("ID = ?", id).First(info).Error
}

// 获取数据列表
func (owner *DataModel) List(data *model.Page[repository.RequestDataModelQuery]) ([]*model.DataModel, error) {
	list := make([]*model.DataModel, 0)
	err := service.Filter(owner, model.DataModel{}.TableName(), data, &list, func(db *gorm.DB) (*gorm.DB, error) {

		if data.Condition.Condition != "" {
			db = db.Where("(title like ? or mark like ?)", "%"+data.Condition.Condition+"%", "%"+data.Condition.Condition+"%")
		}

		return db, nil
	})

	return list, err
}

// SaveModelInfo
// 保存数据模型明细信息
func (owner *DataModel) SaveModelInfo(modelId string, data []*repository.RequestModelInfoSave) error {

	// 先删除模型明细信息
	err := owner.GetDB().Delete(&model.ModelInfo{}, "data_model_id = ?", modelId).Error
	if err != nil {
		return err
	}

	saveData := make([]*model.ModelInfo, len(data))
	// 保存模型明细信息
	for index, item := range data {
		saveData[index] = model.NewModelInfo(&model.ModelInfoBase{
			DataModelId: modelId,
			Title:       item.Title,
			Name:        item.Name,
			ColumnType:  item.ColumnType,
			Size:        item.Size,
			IsPk:        item.IsPk,
			Extension:   item.Extension,
			ControlType: item.ControlType,
			Mark:        item.Mark,
		})
	}

	return owner.GetDB().Model(&model.ModelInfo{}).Create(saveData).Error
}

// GetModelInfo
// 获取数据模型明细信息
func (owner *DataModel) GetModelInfo(modelId string) ([]*model.ModelInfo, error) {
	list := make([]*model.ModelInfo, 0)
	return list, owner.GetDB().Model(&model.ModelInfo{}).Where("data_model_id = ?", modelId).Find(&list).Error
}

func (owner *DataModel) GetVarByKey(md_id, key string) (*model.Variable, error) {
	info := &model.Variable{}
	return info, owner.GetDB().Model(&model.Variable{}).Where("md_id = ? and key = ?", md_id, key).Find(info).Error
}

func (owner *DataModel) GetVarsByKey(md_id string) ([]*model.Variable, error) {
	info := make([]*model.Variable, 0)
	return info, owner.GetDB().Model(&model.Variable{}).Where("md_id = ?", md_id).Find(&info).Error
}

func (owner *DataModel) DelVarByKey(md_id, key string) error {
	return owner.GetDB().Delete(&model.Variable{}, "md_id = ? and key = ?", md_id, key).Error
}
