# cmd

## 描述

此包包含应用程序的命令行界面.

## 导入

```typescript
import * as cmd from 'go/cmd'
```

## 构建方法

### 1. `newExec(timeOut int)`

- timeOut: 超时时间，单位为秒；注：0 不超时.
- return: 返回执行对象的示例.

#### 示例
```typescript
let exec = cmd.newExec(10)
```

## 方法

- `setWorkDir(workDir string)`
- `getWorkDir():string`
- `setEnv(key string, value string)`
- `getEnv(key string):string`
- `run(name string, args ...string):object`
- `runCommand(name string, args []string, callback func(string))`
- `runTimeOut(name string, args []string, timeOut int64, callback func(string))`

### 1. `setWorkDir(workDir string)`

- workDir: 工作区路径.

设置工作区.

#### 示例
```javascript
let exec = cmd.newExec(0)

exec.setWorkDir("/path/to/workdir")
```

### 2. `getWorkDir():string`

- return: 工作区路径.

获取工作区.

#### 示例
```javascript
let exec = cmd.newExec(0)

exec.setWorkDir("/path/to/workdir")
let workDir = exec.getWorkDir()
```

### 3. `setEnv(key string, value string)`

- key: 环境变量名称.
- value: 环境变量值.

设置环境变量.

#### 示例
```javascript
let exec = cmd.newExec(0)

exec.setEnv("key", "value")
```

### 4. `getEnv(key string):string`

- key: 环境变量名称.
- return: 环境变量值.

获取环境变量.

#### 示例
```javascript
let exec = cmd.newExec(0)

exec.setEnv("key", "value")
let value = exec.getEnv("key")
```

### 5. `run(name string, args ...string):object`

- name: 命令名称.
- args: 命令参数.
- return: object.

运行命令，等待内容在返回值中输出.

```javascript
object = {
    ok: bool,
    stdout: string,
    stderr: string
}
```

执行命令.

#### 示例
```javascript
let exec = cmd.newExec(0)

let result = exec.run("cmd", ["/C", "dir"])
```

### 6. `runCommand(name string, args []string, callback func(string))`

- name: 命令名称.
- args: 命令参数.
- callback: 回调函数.

执行命令，并实时输出结果.

#### 示例
```javascript
let exec = cmd.newExec(0)

exec.runCommand("cmd", ["/C", "dir"], (result) => {
    console.log(result)
})
```


### 7. `runTimeOut(name string, args []string, timeOut int64, callback func(string))`

- name: 命令名称.
- args: 命令参数.
- timeOut: 超时时间，单位为秒；注：0 不超时.
- callback: 回调函数.

执行命令，并实时输出结果，超时后自动停止执行.

#### 示例
```javascript
let exec = cmd.newExec(0)

exec.runTimeOut("cmd", ["/C", "ping", "www.baidu.com"], 10, (result) => {
    console.log(result)
})
```