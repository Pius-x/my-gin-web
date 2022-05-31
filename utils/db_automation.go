package utils

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

//@author: [songzhibin97](https://github.com/songzhibin97)
//@function: ClearTable
//@description: 清理数据库表数据
//@param: dbInstance(数据库对象) *gorm.DB, tableName(表名) string, compareField(比较字段) string, interval(间隔) string
//@return: error

// ClearTable 清理数据库表数据
func ClearTable(db *sqlx.DB, tableName string, compareField string, interval string) error {
	if db == nil {
		return errors.New("dbInstance Cannot be empty")
	}
	duration, err := time.ParseDuration(interval)
	if err != nil {
		return err
	}
	if duration < 0 {
		return errors.New("parse duration < 0")
	}

	_, err = db.Exec(fmt.Sprintf("DELETE FROM %s WHERE %s < ?", tableName, compareField), time.Now().Add(-duration))

	return err
}
