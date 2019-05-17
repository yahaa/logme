package logme

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var (
	Log      *logrus.Logger
	Args0Err = fmt.Errorf("args[0] type should be string")
	Args1Err = fmt.Errorf("args[1] type should be string")
	Args2Err = fmt.Errorf("args[2] type should be time.Duration")
	Args3Err = fmt.Errorf("args[3] type should be time.Duration")
)

const (
	goPath = "GOPATH"
)

func init() {
	// default log writer to stdout
	var err error
	Log, err = New("default", "debug")
	if err != nil {
		panic(err)
	}
}

func New(svc, level string, args ...interface{}) (log *logrus.Logger, err error) {
	var (
		path   string
		name   string
		maxAge time.Duration
		rota   time.Duration
		ok     bool

		n = len(args)
	)

	switch {
	case n >= 4:
		maxAge, ok = args[2].(time.Duration)
		if !ok {
			return nil, Args2Err
		}
		rota, ok = args[3].(time.Duration)

		if !ok {
			return nil, Args3Err
		}
		fallthrough

	case n >= 2:
		path, ok = args[0].(string)
		if !ok && len(path) == 0 {
			return nil, Args0Err
		}
		name, ok = args[1].(string)

		if !ok && len(path) == 0 {
			return nil, Args1Err
		}
	}

	return newLog(svc, level, path, name, &maxAge, &rota)
}

// newLog
func newLog(svc, logLevel, logPath, filename string, maxAge, rotaTime *time.Duration) (log *logrus.Logger, err error) {
	var (
		basePath = path.Join(logPath, filename)
		age      = time.Hour * 24 * 7
		rota     = time.Hour

		fullPath string
	)

	if maxAge != nil {
		age = *maxAge
	}

	if rotaTime != nil {
		rota = *rotaTime
	}

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}

	if basePath != "" {
		_, err := os.Stat(logPath)
		if err != nil {
			return nil, err
		}

		if !path.IsAbs(basePath) {
			pwd, err := os.Getwd()
			if err != nil {
				return nil, err
			}
			fullPath = path.Join(pwd, basePath)
		} else {
			fullPath = basePath
		}

		w, err := rotatelogs.New(
			fullPath+".%Y%m%d%H%M",
			rotatelogs.WithLinkName(fullPath), // 生成软链，指向最新日志文
			rotatelogs.WithMaxAge(age),        // 文件最大保存时间
			rotatelogs.WithRotationTime(rota), // 日志切割时间间隔
		)
		if err != nil {
			return nil, err
		}
		log = &logrus.Logger{
			Out: w,
			Formatter: &logrus.JSONFormatter{
				CallerPrettyfier: caller,
			},
			Hooks:    make(logrus.LevelHooks),
			Level:    level,
			ExitFunc: os.Exit,
		}
	} else {
		log = &logrus.Logger{
			Out: os.Stdout,
			Formatter: &logrus.TextFormatter{
				CallerPrettyfier: caller,
				FullTimestamp:    true,
			},
			Hooks:    make(logrus.LevelHooks),
			Level:    level,
			ExitFunc: os.Exit,
		}
	}

	log.AddHook(&logHook{svcName: svc})
	log.SetReportCaller(true)
	return
}

func caller(f *runtime.Frame) (string, string) {
	repPath := fmt.Sprintf("%s/src/", os.Getenv(goPath))
	filename := strings.Replace(f.File, repPath, "", -1)
	return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
}

// logHook hooks
type logHook struct {
	svcName string
}

// Fire fire hooks
func (df *logHook) Fire(entry *logrus.Entry) error {
	entry.Data["svc"] = df.svcName
	return nil
}

// Levels all levels
func (df *logHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
