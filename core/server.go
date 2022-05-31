package core

import (
	"fmt"
	"time"

	"github.com/my-gin-web/global"
	"github.com/my-gin-web/initialize"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.Config.System.UseMultipoint || global.Config.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.Config.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.ZapLog.Info("server run success on ", zap.String("address", address))

	global.ZapLog.Error(s.ListenAndServe().Error())
}
