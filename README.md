# leetcode-withTest

## 说明
本仓库主要涉及少量leetcode作为学习使用，其中添加**覆盖率测试**，**表格测试**，**性能测试**，测试代码及输出内容均可查看。

## 使用方法

代码**覆盖率测试**
```go
// 生成二进制文件
go test . -coverprofile=c.out
// 以网页形式查看
go tool cover -html=c.out
```

代码**表格驱动测试**
```go
go test .
```

代码**性能测试**
```go
// 可直接进行测试
go test -bench .
// 使用pprof工具
go test -bench . -cpuprofile=cpu.out
// 使用 go tool 打开
go tool pprof cpu.out
// 选择使用 web
(pprof) web
```
启动**火焰图**的方法：
```go

import "net/http"
import _ "net/http/pprof"
func main() {
    // 主函数中添加
	http.ListenAndServe("0.0.0.0:9999", nil)
 
}
```
并且在命令行中输入：
```go
go tool pprof -http=:6666 http://localhost:9999/debug/pprof/profile
```
等待30s左右后可以在浏览器中查看

