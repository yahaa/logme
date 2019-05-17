package logme

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	// default log
	Log = New("debug", "")
)

// New
// path 日志文件路径
// l 日志级别
// args[0] 每份 log 文件大小，单位 MB，默认 512M
// args[1] 日志保存天数，默认 7 天
// args[2] 滚动保存的文件数量，默认 3 份
// args[3] 是否需要压缩，默认开启压缩 args[3]=0 不开启，args[3]=1 开启
func New(level string, path string, args ...int) *zap.Logger {
	if len(path) == 0 {
		return newStdLog()
	}

	return newLog(path, level, args...)
}

func newStdLog() *zap.Logger {
	log, _ := zap.NewProduction(
		zap.AddCaller(),
		zap.AddStacktrace(zap.FatalLevel),
	)

	return log
}

// newLog
// path 日志文件路径
// l 日志级别
// args[0] 每份 log 文件大小，单位 MB，默认 512M
// args[1] 日志保存天数，默认 7 天
// args[2] 滚动保存的文件数量，默认 3 份
// args[3] 是否需要压缩，默认开启压缩 args[3]=0 不开启，args[3]=1 开启
func newLog(path, l string, args ...int) *zap.Logger {
	var (
		maxSize  = 512
		maxAge   = 7
		maxFiles = 3
		compress = false
		n        = len(args)
	)

	switch {
	case n >= 4:
		compress = args[3] == 1

	case n >= 3:
		maxFiles = args[2]
		fallthrough
	case n >= 2:
		maxAge = args[1]
		fallthrough
	case n >= 1:
		maxSize = args[0]
	}

	hook := lumberjack.Logger{
		Filename:   path,
		MaxSize:    maxSize,
		MaxBackups: maxFiles,
		MaxAge:     maxAge,
		Compress:   compress,
	}
	w := zapcore.AddSync(&hook)

	var level zapcore.Level
	switch l {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		w,
		level,
	)

	logger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zap.FatalLevel),
	)

	return logger
}
