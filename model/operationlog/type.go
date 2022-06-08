package operationlog

import (
	"github.com/my-gin-web/global"
	"github.com/my-gin-web/model/common"
)

// TOperationLog 操作记录表结构
type TOperationLog struct {
	global.GvaModel
	Ip           string `json:"ip" form:"ip" db:"ip"`                               // 请求ip
	Method       string `json:"method" form:"method" db:"method"`                   // 请求方法
	Path         string `json:"path" form:"path" db:"path"`                         // 请求路径
	Status       uint64 `json:"status" form:"status" db:"status"`                   // 请求状态
	Latency      uint64 `json:"latency" form:"latency" db:"latency"`                // 延迟
	Agent        string `json:"agent" form:"agent" db:"agent"`                      // 代理
	ErrorMessage string `json:"error_message" db:"error_message"`                   // 错误信息
	Body         string `json:"body" form:"body" db:"body"`                         // 请求Body
	Resp         string `json:"resp" form:"resp" db:"resp"`                         // 响应Body
	UserAccount  string `json:"user_account" form:"user_account" db:"user_account"` // 用户账号
}

type SearchOperationLog struct {
	TOperationLog
	common.PageInfo
}
