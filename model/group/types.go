package group

type RouteInfo struct {
	Path     string `json:"path"`
	Readonly int64  `json:"readonly"`
}

type TGroups struct {
	Gid        int64  `json:"gid" db:"gid" form:"gid"`                         // 分组ID
	Gname      string `json:"gname" db:"gname" form:"gname"`                   // 分组名
	ParentGid  int64  `json:"parent_gid" db:"parent_gid" form:"parent_gid"`    // 分组名
	RouterList string `json:"router_list" db:"router_list" form:"router_list"` // 权限列表
	CreateTime int64  `json:"create_time" db:"create_time" form:"create_time"` // 创建时间
	UpdateTime int64  `json:"update_time" db:"update_time" form:"update_time"` // 更新时间
}

type GroupInfo struct {
	Gid        int64       `json:"gid" db:"gid" form:"gid"`                         // 分组ID
	Gname      string      `json:"gname" db:"gname" form:"gname"`                   // 分组名
	ParentGid  int64       `json:"parent_gid" db:"parent_gid" form:"parent_gid"`    // 分组名
	RouterList []RouteInfo `json:"router_list" db:"router_list" form:"router_list"` // 权限列表
	CreateTime int64       `json:"create_time" db:"create_time" form:"create_time"` // 创建时间
	UpdateTime int64       `json:"update_time" db:"update_time" form:"update_time"` // 更新时间
	Children   []GroupInfo `json:"children" db:"children" form:"children"`          // 子分组信息
}
