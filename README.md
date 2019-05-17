## 常用第三方 log 库封装

### v1.0.0 封装了 logrus 库实现了如下功能

* 实现了常用的 Log 终端输出带 caller 字段
* 实现了多 log 实例的构造函数，支持写入到文件，并对 log 文件进行切割，设置过期时间

### v2.0.0 实现了 zap 库的封装
* 实现了默认 Log ，终端输出，并且带 caller 字段。
* * 实现了多 log 实例的构造函数，支持写入到文件，并对 log 文件进行切割，设置过期时间，对 log 进行压缩。

### v1.0.0 使用

获取代码;

```bash
$ go get "github.com/yahaa/logme@v1.0.0"
```

示例:

```go
package main

import (
	"time"
	
	"github.com/yahaa/logme"
)

func main() {
	// 使用 stdout log
	logme.Log.Infof("5555")

	// 日子输出到当前目录(./log)的 cvc.log 文件，默认 7 保留七天，一个小时一个 log 文件（要确保 ./log 文件存在)
	appLog0, err := logme.New("test0", "debug", "./log", "cvc.log")
	if err != nil {
		panic(err)
	}

	appLog0.Warn("this is test0")

	// 日子输出到当前目录(/Users/zihua/repo/logme)的 bc.log 文件，默认 5 保留五分钟，一分钟一个 log 文件（要确保目录存在)
	appLog1, err := logme.New(
		"test1",
		"debug",
		"/Users/zihua/repo/logme",
		"bc.log",
		time.Second*300,
		time.Second*60,
	)

	if err != nil {
		panic(err)
	}

	appLog1.Info("this is test 1")

	// 日子输出到当前目录(./)的 cbc.log 文件，默认 5 保留一小时，三十秒一个 log 文件（要确保目录存在)
	appLog2, err := logme.New("test2", "debug", "./", "cbc.log", time.Hour, time.Second*30)

	if err != nil {
		panic(err)
	}

	appLog2.Info("this is test 3")
}
```

### v2.0.0 使用

获取代码:
```bash
$ go get "github.com/yahaa/logme@v2.0.0"
```

示例:
```go
package logme

import (
	"testing"
	"time"
	
	"go.uber.org/zap"
)

func TestNew(t *testing.T) {
	Log.With(zap.String("svc", "test")).Info("this is test")

	newLog := New("debug", "")
	newLog.With(zap.String("svc", "newLog")).Info("this is new log")

	newLog1 := New("debug", "./xxx/test.log", 1, 1)

	for {
		newLog1.With(zap.String("svc", "newLog1")).Info("this is new log 1")
		time.Sleep(time.Second)
	}
}

```