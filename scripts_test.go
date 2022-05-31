package main_test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/my-gin-web/core"
	"github.com/my-gin-web/initialize"
	"testing"
)

func TestGo(t *testing.T) {
	core.Viper()           // 初始化Viper配置库
	core.Zap()             // 初始化zap日志库
	initialize.Timer()     // 初始化定时任务管理
	initialize.MysqlDb()   // 初始化数据库列表（有多个数据库实例）
	initialize.InitModel() // 初始化Model

	var userInfo users.TUsers

	fmt.Printf("userInfo %+v \n", userInfo)

}
