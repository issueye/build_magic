package controller

import "context"

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
func (control *Code) RunCode(dmId string) error {
	// err := logic.NewCodeLogic(control.Ctx).RunCode(dmId, false, "")
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (control *Code) TestRunCode(dmId string, tpId string) error {
	// err := logic.NewCodeLogic(control.Ctx).RunCode(dmId, true, tpId)
	// if err != nil {
	// 	return err
	// }

	return nil
}
