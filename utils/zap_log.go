package utils

import (
	"fmt"
	"github.com/my-gin-web/global"
)

func ZapInfoLog(info any) {
	global.ZapCallerLog.Info(fmt.Sprintf("%+v \n", info))
}

func ZapErrorLog(errInfo any) {
	global.ZapLog.Error(fmt.Sprintf("stacktrace:\n %+v\n", errInfo))
}
