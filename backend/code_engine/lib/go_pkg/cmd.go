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
	"golang.org/x/text/encoding/simplifiedchinese"
)

var passthroughEnvVars = []string{"HOME", "USER", "USERPROFILE", "TMPDIR", "TMP", "TEMP", "PATH"}

type Exec struct {
	workDir string
	envs    map[string]string // 环境变量
	stop    chan struct{}     // 停止信号
	timeOut int64             // 超时时间
}

func NewExec(timeOut int64) *Exec {

	ec := &Exec{
		envs:    make(map[string]string),
		stop:    make(chan struct{}),
		timeOut: timeOut,
	}

	// 设置工作目录
	for _, key := range passthroughEnvVars {
		if value := os.Getenv(key); value != "" {
			ec.envs[key] = value
		}
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

// Stop
// 停止命令行
func (e *Exec) Stop() {
	e.stop <- struct{}{}
}

// SetTimeOut
// 设置超时时间
func (e *Exec) SetTimeOut(timeOut int64) {
	e.timeOut = timeOut
}

// GetTimeOut
// 获取超时时间
func (e *Exec) GetTimeOut() int64 {
	return e.timeOut
}

func (e *Exec) newRun(name string, args []string, timeOut int64) (ctx context.Context, cancel context.CancelFunc, cmd *exec.Cmd) {

	if timeOut > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	} else if timeOut <= 0 && e.timeOut > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), time.Duration(e.timeOut)*time.Second)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}

	cmd = exec.CommandContext(ctx, name, args...)

	// 设置工作区
	if e.workDir != "" {
		cmd.Dir = e.workDir
	}

	// 设置环境变量
	cmd.Env = []string{}
	for k, v := range e.envs {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}

	return
}

// Run
// 执行命令行
func (e *Exec) Run(name string, args ...string) map[string]any {
	fmt.Println("运行命令：", name, args)

	ctx, cancel, cmd := e.newRun(name, args, 0)
	defer cancel()

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
			cancel()
			if err != nil {
				okResult = false
			}
		}
	case <-e.stop:
		{
			cancel()
			okResult = false
			_ = cmd.Process.Kill()
		}
	case <-ctx.Done():
		{
			cancel()
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

// pipeline
// 管道处理
func (e *Exec) pipeline(ctx context.Context, cancel context.CancelFunc, cmd *exec.Cmd, callback func(string)) error {

	defer cancel()

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
			select {
			case <-e.stop:
				return
			case <-ctx.Done():
				return
			default:
				{
					buf := make([]byte, 1024)
					strNum, err := stdout.Read(buf)
					if err != nil {
						break
					}

					output := buf[:strNum]

					callback(ConvertByte2String(output, GB18030))
				}
			}
		}
	}()

	select {
	case err := <-done:
		{
			cancel()
			if err != nil {
				return err
			}

			return nil
		}
	case <-ctx.Done():
		{
			cancel()
			_ = cmd.Process.Kill()
			return nil
		}
	case <-e.stop:
		{
			cancel()
			_ = cmd.Process.Kill()
			return nil
		}
	}
}

// RunTimeOut
// 运行命令行
func (e *Exec) RunTimeOut(name string, args []string, timeOut int64, callback func(string)) error {
	fmt.Println("运行命令：", name, args)

	ctx, cancel, cmd := e.newRun(name, args, timeOut)
	return e.pipeline(ctx, cancel, cmd, callback)
}

func (e *Exec) RunCommand(name string, args []string, callback func(string)) error {
	fmt.Println("运行命令：", name, args)

	ctx, cancel, cmd := e.newRun(name, args, 0)
	return e.pipeline(ctx, cancel, cmd, callback)
}

func InitCmd() {
	require.RegisterNativeModule("go/cmd", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)

		o.Set("newExec", func(timeOut int) *goja.Object {
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
	obj.Set("runTimeOut", e.RunTimeOut)
	obj.Set("stop", e.Stop)

	return obj
}

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}
