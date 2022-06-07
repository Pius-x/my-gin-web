package utils

import (
	"fmt"
	"github.com/my-gin-web/global"
)

func ZapInfoLog(info interface{}) {
	global.ZapCallerLog.Info(fmt.Sprintf("%+v \n", info))
}

func ZapErrorLog(errInfo interface{}) {
	global.ZapLog.Error(fmt.Sprintf("stacktrace:\n %+v\n", errInfo))
}
