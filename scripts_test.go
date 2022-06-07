package main_test

import (
	"fmt"
	"github.com/my-gin-web/core"
	"github.com/my-gin-web/global"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGo(t *testing.T) {
	core.Viper() // 初始化Viper配置库
	core.Zap()   // 初始化zap日志库
	_, err := ReadConfig()
	if err != nil {
		//fmt.Printf("original err:%T %v\n", errors.Cause(err), errors.Cause(err))
		//fmt.Printf("stack trace:\n %+v\n", err) // %+v 可以在打印的时候打印完整的堆栈信息

		global.ZapLog.Error(fmt.Sprintf("err err:%v\n", err.Error()))
		global.ZapLog.Error(fmt.Sprintf("original err:%v\n", errors.Cause(err)))
		global.ZapLog.Error(fmt.Sprintf("original err:%s \n", err))
		global.ZapLog.Error(fmt.Sprintf("stack trace:\n %+v\n", err))
		os.Exit(1)
	}
}

func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "read failed")
	}
	return buf, nil
}

func ReadConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := ReadFile(filepath.Join(home, ".settings.xml"))
	return config, errors.WithMessage(err, "cound not read config")
}
