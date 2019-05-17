package logme

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	Log.Info("test test")
	newLog, err := New("test", "debug", "", "bbb.log")
	if err != nil {
		fmt.Println(err)
		return
	}

	newLog.Info("44444")

	//pwd, _ := os.Getwd()
	//
	//fmt.Println(path.Join(pwd, "./log/bbba.log"))
	//for {
	//	newLog.Infof("this is debug")
	//	time.Sleep(time.Second)
	//}

}
