package service

import (
	"github.com/issueye/build_magic/backend/common/model"
	"github.com/issueye/build_magic/backend/common/service"
)

// VariableService
// 变量
type VariableService struct {
	service.BaseService
}

func (owner *VariableService) SetBase(base service.BaseService) {
	owner.BaseService = base
}

func (owner *VariableService) CreateOrModify(ProID string, data *model.VariableBase) (err error) {
	var info *model.Variable
	info, err = owner.GetVarByKey(data.ProID, data.Key)
	if err != nil {
		return err
	}

	if info.Key == "" {
		info = model.NewVariable(data)
	} else {
		info.Value = data.Value
		info.Mark = data.Mark
	}

	info.ProID = ProID

	return owner.GetDB().Model(&model.Variable{}).Where("pro_id = ? and key = ?", ProID, data.Key).Save(info).Error
}

func (owner *VariableService) GetVarByKey(md_id, key string) (*model.Variable, error) {
	info := &model.Variable{}
	return info, owner.GetDB().Model(&model.Variable{}).Where("pro_id = ? and key = ?", md_id, key).Find(info).Error
}

func (owner *VariableService) GetVarsByKey(md_id string) ([]*model.Variable, error) {
	info := make([]*model.Variable, 0)
	return info, owner.GetDB().Model(&model.Variable{}).Where("pro_id = ?", md_id).Find(&info).Error
}

func (owner *VariableService) DelVarByKey(md_id, key string) error {
	return owner.GetDB().Delete(&model.Variable{}, "pro_id = ? and key = ?", md_id, key).Error
}
