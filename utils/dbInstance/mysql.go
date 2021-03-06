package dbInstance

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/my-gin-web/global"
	"github.com/pkg/errors"
)

// 数据库连接名配置
const (
	Default    = "default"     // 默认数据库连接
	Games      = "games"       // 游戏业务主库
	SlaveGames = "slave_games" // 游戏业务从库
)

// SelectConn 通过连接名获取 DbMap 中的Db实例 如果不存在则panic
func SelectConn(connName ...string) *sqlx.DB {
	var theConn string
	switch len(connName) {
	case 0:
		theConn = Default
	case 1:
		theConn = connName[0]
	default:
		panic(errors.New("params num error"))
	}

	db, ok := global.DbMap[theConn]
	if !ok || db == nil {
		panic(errors.New(fmt.Sprintf("%s no init", theConn)))
	}
	return db
}

type Db struct {
	*sqlx.DB
}

func (db *Db) Insert(dest any) {
	// 注意：这里必须是params... 不然会转成数组类型了
	//res, err := db.Exec(query, params...)
	// 打印上一次插入自增的id
	//if res != nil {
	//	fmt.Println(res.LastInsertId())
	//} else {
	//	fmt.Printf("插入失败：%v\n", err)
	//}
}

func (db *Db) Delete(sql string, params ...any) {
	res, err := db.Exec(sql, params...)
	if err != nil {
		fmt.Printf("删除失败：%v\n", err)
	} else {
		fmt.Println(res.RowsAffected())
	}
}

func (db *Db) Update(sql string, params ...any) {
	res, err := db.Exec(sql, params...)
	if err != nil {
		fmt.Printf("更新失败：%v\n", err)
	} else {
		fmt.Println(res.RowsAffected())
	}
}
