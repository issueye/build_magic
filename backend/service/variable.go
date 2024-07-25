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

func (owner *VariableService) CreateOrModify(pro_id string, varType int, data *model.VariableBase) (err error) {
	var info *model.Variable
	info, err = owner.GetVarByKey(data.ProID, data.Key, data.VarType)
	if err != nil {
		return err
	}

	if info.Key == "" {
		info = model.NewVariable(data)
	} else {
		info.Value = data.Value
		info.Mark = data.Mark
	}

	info.ProID = pro_id

	return owner.GetDB().
		Model(&model.Variable{}).
		Where("pro_id = ? and key = ? and var_type = ?", pro_id, data.Key, varType).
		Save(info).Error
}

func (owner *VariableService) GetVarByKey(pro_id, key string, varType int) (*model.Variable, error) {
	info := &model.Variable{}
	return info, owner.GetDB().
		Model(&model.Variable{}).Where("pro_id = ? and key = ? and var_type = ?", pro_id, key, varType).
		Find(info).Error
}

func (owner *VariableService) GetVarsByKey(pro_id string, varType int) ([]*model.Variable, error) {
	info := make([]*model.Variable, 0)
	return info, owner.GetDB().
		Model(&model.Variable{}).
		Where("pro_id = ? and var_type = ?", pro_id, varType).
		Find(&info).Error
}

func (owner *VariableService) DelVarByKey(pro_id, key string, varType int) error {
	return owner.GetDB().
		Delete(&model.Variable{}, "pro_id = ? and key = ? and var_type = ?", pro_id, key, varType).Error
}

func (owner *VariableService) DelVarByType(pro_id string, varType int) error {
	return owner.GetDB().
		Delete(&model.Variable{}, "pro_id = ? and var_type = ?", pro_id, varType).Error
}
