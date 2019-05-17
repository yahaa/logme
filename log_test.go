package logme

import (
	"go.uber.org/zap"
	"testing"
	"time"
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
