package user

import "github.com/my-gin-web/model/group"

type CoreInfo struct {
	Id         int64  `json:"id" db:"id"`
	Account    string `json:"account" db:"account"`         //账号
	Gid        int64  `json:"gid" db:"gid"`                 //分组ID
	Name       string `json:"name" db:"name"`               //用户昵称
	Mobile     string `json:"mobile" db:"mobile"`           //手机号码
	CreateTime int64  `json:"create_time" db:"create_time"` //创建时间
}

type TUsers struct {
	CoreInfo
	Password      string `json:"password" db:"password"`               //密码
	HeadPic       int64  `json:"head_pic" db:"head_pic"`               //头像ID
	UpdateTime    int64  `json:"update_time" db:"update_time"`         //更新时间
	LastLoginTime int64  `json:"last_login_time" db:"last_login_time"` //最后登录时间
	CreateBy      string `json:"create_by" db:"create_by"`             //创建人
}

type LoginInfo struct {
	TUsers
	RouterList []group.RouteInfo `json:"router_list"`
}
