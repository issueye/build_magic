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
func (lc *Variable) GetVariableList(ProID string) ([]*model.Variable, error) {
	srv := commonService.NewService(&service.VariableService{})
	return srv.GetVarsByKey(ProID)
}

// 删除变量
func (lc *Variable) DeleteVariable(ProID, id string) error {
	srv := commonService.NewService(&service.VariableService{})
	return srv.DelVarByKey(ProID, id)
}

// 新增或者修改变量
func (lc *Variable) AddOrUpdateVariable(ProID string, list []*model.VariableBase) (err error) {
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

	for _, data := range list {
		err = srv.CreateOrModify(ProID, data)
		if err != nil {
			return
		}
	}

	return
}
