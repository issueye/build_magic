package controller

import (
	"context"

	"github.com/issueye/build_magic/backend/logic"
)

var code *Code

type Code struct {
	Ctx context.Context
}

func GetCode() *Code {
	if code == nil {
		code = &Code{}
	}

	return code
}

// 运行代码
func (control *Code) RunCode(code string) error {
	err := logic.NewCodeLogic(control.Ctx).RunCode(code)
	if err != nil {
		return err
	}

	return nil
}

func (control *Code) TestRunCode(dmId string, tpId string) error {
	err := logic.NewCodeLogic(control.Ctx).RunCode(tpId)
	if err != nil {
		return err
	}

	return nil
}
