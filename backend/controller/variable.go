package controller

import (
	"context"

	"github.com/issueye/build_magic/backend/common/model"
	"github.com/issueye/build_magic/backend/logic"
)

var variable *Variable

// 数据模型
type Variable struct {
	Ctx context.Context
}

// NewApp creates a new App application struct
func GetVariable() *Variable {
	if variable == nil {
		variable = &Variable{}
	}

	return variable
}

// 获取变量列表
func (lc *Variable) GetVariableList(ProID string, varType int) ([]*model.Variable, error) {
	return new(logic.Variable).GetVariableList(ProID, varType)
}

// 删除变量
func (lc *Variable) DeleteVariable(ProID string, id string, varType int) error {
	return new(logic.Variable).DeleteVariable(ProID, id, varType)
}

// 保存变量
func (lc *Variable) SaveVariable(ProID string, varType int, variables []*model.VariableBase) error {
	return new(logic.Variable).AddOrUpdateVariable(ProID, varType, variables)
}
