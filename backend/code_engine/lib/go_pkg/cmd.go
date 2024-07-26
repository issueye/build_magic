package gopkg

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"github.com/issueye/build_magic/backend/pkg/utils"
)

var passthroughEnvVars = []string{"HOME", "USER", "USERPROFILE", "TMPDIR", "TMP", "TEMP", "PATH"}

type Exec struct {
	workDir string
	ctx     context.Context
	cancel  context.CancelFunc
	envs    map[string]string // 环境变量
}

func NewExec(timeOut int64) *Exec {

	ec := &Exec{
		envs: make(map[string]string),
	}

	// 设置工作目录
	for _, key := range passthroughEnvVars {
		if value := os.Getenv(key); value != "" {
			ec.envs[key] = value
		}
	}

	if timeOut > 0 {
		ec.ctx, ec.cancel = context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	} else {
		ec.ctx, ec.cancel = context.WithCancel(context.Background())
	}

	return ec
}

// SetWorkDir
// 设置工作区
func (e *Exec) SetWorkDir(workDir string) {
	e.workDir = workDir
}

// GetWorkDir
// 获取工作区
func (e *Exec) GetWorkDir() string {
	return e.workDir
}

// SetEnv
// 设置环境变量
func (e *Exec) SetEnv(key, value string) {
	e.envs[key] = value
}

// GetEnv
// 获取环境变量
func (e *Exec) GetEnv(key string) string {
	return e.envs[key]
}

// Run
// 执行命令行
func (e *Exec) Run(name string, args ...string) map[string]any {
	fmt.Println("运行命令：", name, args)

	cmd := exec.Command(name, args...)

	// 设置工作区
	if e.workDir != "" {
		cmd.Dir = e.workDir
	}

	// 设置环境变量
	cmd.Env = []string{}
	for k, v := range e.envs {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}

	var outBuf, errBuf strings.Builder // or bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	var (
		stderr   string
		okResult = true
	)

	err := cmd.Start()
	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()

	select {
	case err := <-done:
		{
			if err != nil {
				okResult = false
			}
		}
	case <-e.ctx.Done():
		{
			okResult = false
			_ = cmd.Process.Kill()
		}
	}

	if err != nil {
		stderr = errBuf.String()
		okResult = false
	}

	return map[string]any{"ok": okResult, "stdout": strings.Trim(outBuf.String(), "\n"), "stderr": stderr}
}

func (e *Exec) RunCommand(name string, args []string, callback func(string)) error {
	fmt.Println("运行命令：", name, args)
	cmd := exec.Command(name, args...)

	// 设置工作区
	if e.workDir != "" {
		cmd.Dir = e.workDir
	}

	// 设置环境变量
	cmd.Env = []string{}
	for k, v := range e.envs {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	cmd.Stderr = cmd.Stdout

	err = cmd.Start()
	if err != nil {
		return err
	}

	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()
	go func() {
		for {
			buf := make([]byte, 1024)
			strNum, err := stdout.Read(buf)
			if err != nil {
				break
			}

			output := buf[:strNum]

			callback(utils.Bytes2String(output))
		}
	}()

	select {
	case err := <-done:
		{
			if err != nil {
				return err
			}
		}
	case <-e.ctx.Done():
		{
			_ = cmd.Process.Kill()
			return nil
		}
	}

	return nil
}

func InitCmd() {
	require.RegisterNativeModule("go/cmd", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)

		o.Set("NewExec", func(timeOut int) *goja.Object {
			return NewExecJs(runtime, NewExec(int64(timeOut)))
		})
	})
}

func NewExecJs(vm *goja.Runtime, e *Exec) *goja.Object {
	obj := vm.NewObject()
	obj.Set("setWorkDir", e.SetWorkDir)
	obj.Set("getWorkDir", e.GetWorkDir)
	obj.Set("setEnv", e.SetEnv)
	obj.Set("getEnv", e.GetEnv)
	obj.Set("run", e.Run)
	obj.Set("runCommand", e.RunCommand)

	return obj
}
