package system

type UserCoreInfo struct {
	Id         uint64 `json:"id"`
	Account    string `json:"account"`                           //账号
	Gid        uint64 `json:"gid"`                               //分组ID
	Name       string `json:"name"`                              //用户昵称
	Mobile     string `json:"mobile"`                            //手机号码
	CreateTime uint64 `json:"create_time" gorm:"autoCreateTime"` //创建时间
}

type TUsers struct {
	UserCoreInfo
	Password   string `json:"password"`                          //密码
	HeadPic    uint64 `json:"head_pic"`                          //头像ID
	UpdateTime uint64 `json:"update_time" gorm:"autoUpdateTime"` //更新时间
	CreateBy   string `json:"create_by"`                         //创建人
}

type LoginInfo struct {
	TUsers
	RouterList []RouteInfo `json:"router_list"`
}
