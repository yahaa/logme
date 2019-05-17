### logrus 封装

### 使用

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