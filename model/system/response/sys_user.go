package response

import "github.com/my-gin-web/model/user"

type SysUserResponse struct {
	User user.TUsers `json:"user"`
}

type LoginResponse struct {
	User      user.LoginInfo `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}

type UserListResponse struct {
	Id         int64  `json:"id"`
	Account    string `json:"account"`     //账号
	Gid        int64  `json:"gid"`         //分组ID
	Name       string `json:"name"`        //用户昵称
	Mobile     string `json:"mobile"`      //手机号码
	CreateTime string `json:"create_time"` //创建时间
	CreateBy   string `json:"create_by"`   //创建人
}
