package operationlog

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/my-gin-web/utils/dbInstance"
	"strings"
	"time"
)

type Model struct {
	db *sqlx.DB
}

func NewModel() *Model {
	return &Model{
		db: dbInstance.SelectConn(),
	}
}

func (This *Model) InsertNewRecord(recordInfo SysOperationRecord) (err error) {
	curTime := time.Now().Unix()
	recordInfo.CreatedAt = curTime
	recordInfo.UpdatedAt = curTime

	var sql = `INSERT INTO gva.sys_operation_records (created_at, updated_at, ip, method, path, status, latency, agent, error_message,
                                       body, resp, user_account)
values (:created_at, :updated_at, :ip, :method, :path, :status, :latency, :agent, :error_message, :body, :resp,
        :user_account)`

	_, err = This.db.NamedExec(sql, recordInfo)
	return err
}

func (This *Model) GetPageRecordByKey(info SysOperationRecordSearch) (recordList []SysOperationRecord, total int64, err error) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 如果有条件搜索 下方会自动创建搜索语句
	var where = "WHERE true "
	if info.UserAccount != "" {
		where += fmt.Sprintf("AND user_account = '%s' ", strings.TrimSpace(info.UserAccount))
	}
	if info.Path != "" {
		where += fmt.Sprintf("AND path LIKE '%%%s%%'", strings.TrimSpace(info.Path))
	}

	if err = This.db.Get(&total, "SELECT COUNT(1) FROM gva.sys_operation_records "+where); err != nil {
		return nil, 0, err
	}

	if err = This.db.Select(&recordList, "SELECT * FROM gva.sys_operation_records "+where+" Order By id desc LIMIT ? OFFSET ? ", limit, offset); err != nil {
		return nil, 0, err
	}

	return recordList, total, err
}
