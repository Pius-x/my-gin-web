package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

// UserList Paging common input parameter structure
type UserList struct {
	Page     int64  `form:"page"`     // 页码
	PageSize int64  `form:"pageSize"` // 每页大小
	Gid      int64  `form:"gid"`      // 分组id
	Keyword  string `form:"keyword"`  // 关键字
}

// UserListByGid Paging common input parameter structure
type UserListByGid struct {
	Gid int64 `form:"gid"` // 分组id
}

// SearchUser  Paging common input parameter structure
type SearchUser struct {
	Condition int `json:"condition"` // 条件
}

// GetGroupListById  Paging common input parameter structure
type GetGroupListById struct {
	Gid int64 `json:"gid" form:"gid"` // 分组Id
}

// GetById Find by id structure
type GetById struct {
	ID int64 `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId string `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}
