package main

import (
	"github.com/my-gin-web/core"
	"github.com/my-gin-web/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {

	core.Viper()           // 初始化Viper配置库
	core.Zap()             // 初始化zap日志库
	initialize.Timer()     // 初始化定时任务管理
	initialize.MysqlDb()   // 初始化数据库列表（有多个数据库实例）
	initialize.InitModel() // 初始化Model

	core.RunWindowsServer()
}
