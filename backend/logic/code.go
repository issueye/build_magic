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
func (lc *CodeLogic) RunCode(tpCodeId string) error {
	// 结束所有事件
	defer func() {
		runtime.EventsOffAll(lc.ctx)
	}()

	global.Log.Infof("开始运行代码，dmId: %s", tpCodeId)

	// 获取项目
	runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, "开始获取项目信息")
	proSrv := commonService.NewService(&service.BuildProjectService{})
	proInfo, err := proSrv.GetByTpCode(tpCodeId)
	if err != nil {
		runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, "获取项目失败")
		return err
	}

	// 获取数据模板变量
	runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("开始项目[%s]环境变量", tpCodeId))
	varSrv := commonService.NewService(&service.VariableService{})
	varList, err := varSrv.GetVarsByKey(proInfo.ID, 0)
	if err != nil {
		runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("获取项目[%s]环境变量失败 %s", proInfo.Title, err))
		return err
	}

	injectVarList, err := varSrv.GetVarsByKey(proInfo.ID, 1)
	if err != nil {
		runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("获取项目[%s]嵌入变量失败 %s", proInfo.Title, err))
		return err
	}

	vm := global.JsVMCore.GetRuntime()
	defer global.JsVMCore.PutRuntime(vm)

	vm.SetProperty("magic", "title", proInfo.Title)
	vm.SetProperty("magic", "name", proInfo.Name)
	vm.SetProperty("magic", "version", proInfo.Version)
	vm.SetProperty("magic", "build_type", proInfo.BuildType)
	vm.SetProperty("magic", "build_version", proInfo.BuildVersion)
	vm.SetProperty("magic", "path", proInfo.ProjectPath)

	runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, "注入环境变量和嵌入变量")

	// 注入变量
	vars := make(map[string]any)
	for _, variable := range varList {
		vars[variable.Key] = variable.Value
	}

	injectVars := make(map[string]any)
	for _, variable := range injectVarList {
		injectVars[variable.Key] = variable.Value
	}

	vm.SetProperty("magic", "env", vars)
	vm.SetProperty("magic", "inject", injectVars)

	vm.SetConsoleCallBack(func(args ...any) {
		// fmt.Println("args", args)
		// 去除第一个参数，第一个参数是日志级别
		runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, args[1])
	})

	// 获取代码模板的脚本
	tpSrv := commonService.NewService(&service.Template{})

	runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("开始获取代码[%s]", tpCodeId))
	tp, err := tpSrv.Get(tpCodeId)
	if err != nil {
		runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("获取代码[%s]失败 %s", tpCodeId, err))
		return err
	}

	err = lc.runCode(vm, tp)
	if err != nil {
		return err
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
