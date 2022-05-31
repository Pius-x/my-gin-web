package initialize

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/my-gin-web/config"
	"github.com/my-gin-web/global"
)

const MYSQL = "mysql"

func MysqlDb() {
	dbMap := make(map[string]*sqlx.DB)
	for _, info := range global.Config.DBList {
		if info.Disable {
			continue
		}
		dbMap[info.ConnName] = MysqlByConfig(info)
	}

	global.DbMap = dbMap
}

// MysqlByConfig 初始化Mysql数据库用过传入配置
func MysqlByConfig(m config.Mysql) *sqlx.DB {
	db := sqlx.MustConnect(MYSQL, m.Dsn())
	db.SetMaxIdleConns(m.MaxIdleConns)
	db.SetMaxOpenConns(m.MaxOpenConns)
	return db
}
