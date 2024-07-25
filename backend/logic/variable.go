package logic

import (
	"github.com/issueye/build_magic/backend/common/model"
	commonService "github.com/issueye/build_magic/backend/common/service"
	"github.com/issueye/build_magic/backend/global"
	"github.com/issueye/build_magic/backend/service"
)

// BuildProject
// 构建项目
type Variable struct{}

// 获取变量列表
func (lc *Variable) GetVariableList(ProID string, varType int) ([]*model.Variable, error) {
	srv := commonService.NewService(&service.VariableService{})
	return srv.GetVarsByKey(ProID, varType)
}

// 删除变量
func (lc *Variable) DeleteVariable(ProID, id string, varType int) error {
	srv := commonService.NewService(&service.VariableService{})
	return srv.DelVarByKey(ProID, id, varType)
}

// 新增或者修改变量
func (lc *Variable) AddOrUpdateVariable(ProID string, varType int, list []*model.VariableBase) (err error) {
	srv := commonService.NewService(&service.VariableService{})
	srv.OpenTx()
	defer func() {
		if err != nil {
			global.Log.Infof("事务回滚：%s", err.Error())
			srv.Rollback()
		} else {
			global.Log.Infof("事务提交")
			srv.Commit()
		}
	}()

	// 删除旧的变量
	err = srv.DelVarByType(ProID, varType)
	if err != nil {
		return
	}

	for _, data := range list {
		err = srv.CreateOrModify(ProID, varType, data)
		if err != nil {
			return
		}
	}

	return
}
