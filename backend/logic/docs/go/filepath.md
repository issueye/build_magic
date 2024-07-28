# filepath

## 描述

此包包含了对文件路径操作.

## 导入

```typescript
import * as fp from 'go/filepath'
```

## 方法

- `abs(path string):string`
- `join(elem ...string) string`
- `ext(path string) string`
- `rel(basepath string, targpath string)string`
- `clean(path string) string`
- `split(path string) [string, string]`
- `dir(path string) string`
- `base(path string) string`

### 1. `abs(path string):string`

- path: 文件路径.
- return: 返回文件标准路径.

获取文件的标准绝对路径； 注：对文件进行斜杠的处理.

#### 示例
```javascript
let path = fp.abs('//path/test.go')

console.log(path) // /path/test.go
```

### 2. `join(elem ...string) string`

- elem: 文件路径.
- return: 返回文件路径.

将多个路径组合成一个路径.

#### 示例
```javascript
let path = fp.join('/path', 'test.go')

console.log(path) // /path/test.go
```

### 3. `ext(path string) string`

- path: 文件路径.
- return: 返回文件扩展名.

获取文件扩展名.

#### 示例
```javascript
let ext = fp.ext('/path/test.go')

console.log(ext) // .go
```

### 4. `rel(basepath string, targpath string)string`

- basepath: 基准路径.
- targpath: 目标路径.
- return: 返回相对路径.

获取目标路径相对于基准路径的相对路径.

#### 示例
```javascript
let rel = fp.rel('/path', '/path/test.go')

console.log(rel) // test.go
```

### 5. `clean(path string) string`

- path: 文件路径.
- return: 返回文件路径.

获取文件路径的规范路径.

#### 示例
```javascript
let path = fp.clean('/path//test.go')

console.log(path) // /path/test.go
```

### 6. `split(path string) [string, string]`

- path: 文件路径.
- return: 返回文件路径.

将文件路径拆分为目录和文件名.

#### 示例
```javascript
let path = fp.split('/path/test.go')

console.log(path) // ["/path", "test.go"]
```

### 7. `dir(path string) string`

- path: 文件路径.
- return: 返回文件路径.

获取文件路径的目录部分.

#### 示例
```javascript
let path = fp.dir('/path/test.go')

console.log(path) // /path
```

### 8. `base(path string) string`

- path: 文件路径.
- return: 返回文件路径.

获取文件路径的文件名部分.

#### 示例
```javascript
let path = fp.base('/path/test.go')

console.log(path) // test.go
```