package common

type PageResult struct {
	List  any   `json:"list"`
	Total int64 `json:"total"`
}

// GetByGid  Find by gid structure
type GetByGid struct {
	Gid int64 `json:"gid" form:"gid"` // 分组Id
}

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int64  `json:"page" form:"page"`         // 页码
	PageSize int64  `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

// GetById Find by id structure
type GetById struct {
	Id int64 `json:"id" form:"id"` // 主键ID
}
