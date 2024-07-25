package logic

import (
	"context"
	"fmt"

	"github.com/issueye/build_magic/backend/code_engine"
	"github.com/issueye/build_magic/backend/common/model"
	commonService "github.com/issueye/build_magic/backend/common/service"
	"github.com/issueye/build_magic/backend/global"
	"github.com/issueye/build_magic/backend/repository"
	"github.com/issueye/build_magic/backend/service"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type MainFunc func() (string, error)

const CONSOLE_EVENT = "console"
const CODE_PUSH_EVENT = "code_push"

type CodeLogic struct {
	ctx context.Context
}

func NewCodeLogic(ctx context.Context) *CodeLogic {
	return &CodeLogic{
		ctx: ctx,
	}
}

// 运行代码
func (lc *CodeLogic) RunCode(dmId string, isTest bool, tpCodeId string) error {
	// 结束所有事件
	defer func() {
		runtime.EventsOffAll(lc.ctx)
	}()

	global.Log.Infof("开始运行代码，dmId: %s", dmId)

	// 查询模板信息
	runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("开始获取模板[%s]信息", dmId))
	srv := commonService.NewService(&service.DataModel{})
	info, err := srv.GetById(dmId)
	if err != nil {
		runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("获取模板[%s]失败 %s", dmId, err))
		return err
	}

	// 获取数据模板变量
	runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("开始获取模板[%s]变量", dmId))
	varList, err := srv.GetVarsByKey(info.ID)
	if err != nil {
		runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("获取模板[%s]变量失败 %s", dmId, err))
		return err
	}

	// 获取模板明细信息
	runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("开始获取模板[%s]明细信息", dmId))
	columns, err := srv.GetModelInfo(dmId)
	if err != nil {
		runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("获取模板明细失败 %s", err))
		return err
	}

	// 将信息注入到虚拟机中
	runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("开始注入模板[%s]信息", dmId))

	vm := global.JsVMCore.GetRuntime()
	defer global.JsVMCore.PutRuntime(vm)

	vm.SetProperty("dm", "title", info.Title)
	vm.SetProperty("dm", "tableName", info.TBName)
	vm.SetProperty("dm", "columns", columns)

	// 注入变量
	vars := make(map[string]any)
	for _, variable := range varList {
		// fmt.Println("variable", variable)
		vars[variable.Key] = variable.Value
	}

	vm.SetProperty("dm", "vars", vars)

	vm.SetConsoleCallBack(func(args ...any) {
		// fmt.Println("args", args)
		// 去除第一个参数，第一个参数是日志级别
		runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, args[1])
	})

	// 获取代码模板的脚本
	tpSrv := commonService.NewService(&service.Template{})

	if !isTest {
		tps, err := tpSrv.GetChildrenByCode(info.SchemeId)
		if err != nil {
			return err
		}

		errArr := make([]error, 0)
		for _, tp := range tps {

			err = lc.runCode(vm, tp)
			if err != nil {
				errArr = append(errArr, err)
			}
		}

		if len(errArr) > 0 {
			runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("执行代码失败 %s", errArr))
			return fmt.Errorf("执行代码失败 %s", errArr)
		}
	} else {
		runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("开始执行测试代码[%s]", tpCodeId))
		tp, err := tpSrv.Get(tpCodeId)
		if err != nil {
			runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("获取模板[%s]失败 %s", tpCodeId, err))
			return err
		}

		err = lc.runCode(vm, tp)
		if err != nil {
			return err
		}
	}

	return nil
}

func (lc *CodeLogic) runCode(vm *code_engine.JsVM, tp *model.CodeTemplate) (err error) {

	path := tp.FileName
	global.Log.Infof("开始执行代码 %s", path)

	var fn MainFunc
	runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, "开始导出[main]函数")

	err = vm.ExportFunc("main", path, &fn)
	if err != nil {
		global.Log.Errorf("导出[main]函数失败 %s", err)
		runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("导出[main]函数失败 %s", err))
		return err
	}

	if fn != nil {
		code, err := fn()
		if err != nil {
			global.Log.Errorf("执行代码[%s]失败 %s", tp.Title, err)
			runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("执行代码[%s]失败 %s", tp.Title, err))
			return err
		}

		fmt.Printf("执行代码[%s] 结果: %s\n", tp.Title, code)
		runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("执行代码[%s] 结果: %s\n", tp.Title, code))

		push := new(repository.PushCode)
		push.Id = tp.ID
		push.CodeContent = code
		runtime.EventsEmit(lc.ctx, CODE_PUSH_EVENT, push)
	}

	return
}
