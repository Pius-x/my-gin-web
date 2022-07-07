package operationlog

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/my-gin-web/utils/dbInstance"
	"github.com/pkg/errors"
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

func (This *Model) InsertNewRecord(recordInfo TOperationLog) {
	curTime := time.Now().Unix()
	recordInfo.CreatedAt = curTime
	recordInfo.UpdatedAt = curTime

	var sql = `INSERT INTO crazy_cms.cms_operation_log (created_at, updated_at, ip, method, path, status, latency, agent,
                                         error_message, body, resp, user_id, user_name)
values (:created_at, :updated_at, :ip, :method, :path, :status, :latency, :agent, :error_message, :body, :resp,
        :user_id, :user_name)`
	if _, err := This.db.NamedExec(sql, recordInfo); err != nil {
		panic(errors.Wrap(err, "新增操作记录失败"))
	}
}

func (This *Model) GetPageRecordByKey(info SearchOperationLog) (recordList []TOperationLog, total int64) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 如果有条件搜索 下方会自动创建搜索语句
	var where = "WHERE true "
	if info.UserName != "" {
		where += fmt.Sprintf("AND user_name = '%s' ", strings.TrimSpace(info.UserName))
	}
	if info.Path != "" {
		where += fmt.Sprintf("AND path LIKE '%%%s%%'", strings.TrimSpace(info.Path))
	}

	if err := This.db.Get(&total, "SELECT COUNT(1) FROM crazy_cms.cms_operation_log "+where); err != nil {
		panic(errors.Wrap(err, "列表条数获取失败"))
	}

	if err := This.db.Select(&recordList, "SELECT * FROM crazy_cms.cms_operation_log "+where+" Order By id desc LIMIT ? OFFSET ? ", limit, offset); err != nil {
		panic(errors.Wrap(err, "操作记录列表获取失败"))
	}

	return recordList, total
}
