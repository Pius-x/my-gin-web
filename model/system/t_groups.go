package system

type RouteInfo struct {
	Path     string `json:"path"`
	Readonly uint64 `json:"readonly"`
}

type TGroups struct {
	Gid        uint64      `json:"gid" gorm:"primaryKey"`              // 分组ID
	Gname      string      `json:"gname"`                              // 分组名
	ParentGid  uint64      `json:"parent_gid"`                         // 分组名
	RouterList []RouteInfo `json:"router_list" gorm:"serializer:json"` // 权限列表
	CreateTime uint64      `json:"create_time" gorm:"autoCreateTime"`  // 创建时间
	UpdateTime uint64      `json:"update_time" gorm:"autoUpdateTime"`  // 更新时间
	Children   []TGroups   `json:"children" gorm:"-"`                  // 子分组信息
}
