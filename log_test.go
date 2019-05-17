package logme

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	newLog, _ := New("test", "debug", "./log", "abcd.log")

	pwd, _ := os.Getwd()

	fmt.Println(path.Join(pwd, "./log/bbba.log"))
	for {
		newLog.Infof("this is debug")
		time.Sleep(time.Second)
	}
}
