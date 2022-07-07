package user

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/my-gin-web/model/common"
	"github.com/my-gin-web/model/group"
)

// ReqLogin 用户登录参数
type ReqLogin struct {
	Username string `form:"username"` // 用户名
	Password string `form:"password"` // 密码
	//Captcha   string `json:"captcha"`   // 验证码
	//CaptchaId string `json:"captchaId"` // 验证码ID
}

type ResLogin struct {
	User      LoginInfo `json:"user"`
	Token     string    `json:"token"`
	ExpiresAt int64     `json:"expires_at"`
}

type LoginInfoRes struct {
	Id            int64             `json:"id"`
	Account       string            `json:"account"`         //账号
	Gid           int64             `json:"gid"`             //分组ID
	HeadPic       int64             `json:"head_pic"`        //头像ID
	Name          string            `json:"name"`            //用户昵称
	LastLoginTime int64             `json:"last_login_time"` //最后登录时间
	RouterList    []group.RouteInfo `json:"router_list"`

	BindFs    bool   `json:"bind_fs"`     // 是否绑定了飞书
	FsHeadPic string `json:"fs_head_pic"` // 飞书账号头像
	FsName    string `json:"fs_name"`     // 飞书账号名字
	ExpiresAt int64  `json:"expires_at"`  // token 过期时间
	Token     string `json:"token"`       // token信息
}

type FsLoginToken struct {
	UserInfo string `json:"user_info"`
}

type FsBindToken struct {
	BindFs    bool   `json:"bind_fs"`
	FsHeadPic string `json:"fs_head_pic"`
	FsName    string `json:"fs_name"`
}

type ResUserList struct {
	Id         int64  `json:"id"`
	Account    string `json:"account"`     // 账号
	Gid        int64  `json:"gid"`         // 分组ID
	Name       string `json:"name"`        // 用户昵称
	Mobile     string `json:"mobile"`      // 手机号码
	CreateTime string `json:"create_time"` // 创建时间
	CreateBy   string `json:"create_by"`   // 创建人
	BindFs     bool   `json:"bind_fs"`     // 是否绑定飞书
}

type TUsers struct {
	Id            int64  `json:"id" db:"id" form:"id"`
	Account       string `json:"account" db:"account" form:"account"`                         //账号
	Gid           int64  `json:"gid" db:"gid" form:"gid"`                                     //分组ID
	Name          string `json:"name" db:"name" form:"name"`                                  //用户昵称
	Mobile        string `json:"mobile" db:"mobile" form:"mobile"`                            //手机号码
	CreateTime    int64  `json:"create_time" db:"create_time" form:"create_time"`             //创建时间
	Password      string `json:"password" db:"password" form:"password"`                      //密码
	HeadPic       int64  `json:"head_pic" db:"head_pic" form:"head_pic"`                      //头像ID
	UpdateTime    int64  `json:"update_time" db:"update_time" form:"update_time"`             //更新时间
	LastLoginTime int64  `json:"last_login_time" db:"last_login_time" form:"last_login_time"` //最后登录时间
	CreateBy      string `json:"create_by" db:"create_by" form:"create_by"`                   //创建人
	BindFs        bool   `json:"bind_fs" db:"bind_fs" form:"bind_fs"`                         //是否绑定飞书
}

type LoginInfo struct {
	TUsers
	RouterList []group.RouteInfo `json:"router_list" from:"router_list" db:"router_list"`
}

// ChangePasswordStruct Modify password structure
type ChangePasswordStruct struct {
	Id          int64  `form:"id" json:"id" db:"id"`                            // 用户ID
	Password    string `form:"password" json:"password" db:"password"`          // 原密码
	NewPassword string `form:"newPassword" json:"newPassword" db:"newPassword"` // 新密码
}

// UpdateHeadPicStruct Modify headPic structure
type UpdateHeadPicStruct struct {
	Id      int64 `form:"id" json:"id" db:"id"`                // 用户ID
	HeadPic int64 `form:"headPic" json:"headPic" db:"headPic"` // 头像ID
}

type UpdateUserInfo struct {
	Id     int64  `json:"id" form:"id" db:"id"`
	Gid    int64  `json:"gid" form:"gid" db:"gid"`          //分组ID
	Name   string `json:"name" form:"name" db:"name"`       //用户昵称
	Mobile string `json:"mobile" form:"mobile" db:"mobile"` //手机号码
}

type UpdateUserGid struct {
	Id  int64 `json:"id" form:"id" db:"id"`
	Gid int64 `json:"gid" form:"gid" db:"gid"`
}

type MultiUpdateUserGid struct {
	UserGidList []UpdateUserGid `json:"userGidList" from:"userGidList" db:"userGidList"` //用户信息列表
}

type CreateUserInfo struct {
	Account    string `form:"account" form:"account" db:"account"`             //账号
	Gid        int64  `form:"gid" form:"gid" db:"gid"`                         //分组ID
	Name       string `form:"name" form:"name" db:"name"`                      //用户昵称
	Mobile     string `form:"mobile" form:"mobile" db:"mobile"`                //手机号码
	CreateTime int64  `form:"create_time" form:"create_time" db:"create_time"` //创建时间
	CreateBy   string `form:"create_by" form:"create_by" db:"create_by"`       //创建人
}

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	ID       int64
	Account  string
	Name     string
	Password string
}

// ReqUserList Paging common input parameter structure
type ReqUserList struct {
	common.PageInfo
	Gid           int64   `form:"gid"`             // 分组id
	FilterGidList []int64 `form:"filter_gid_list"` // 筛选分组id列表
}

type RefreshUserInfo struct {
	Gid        int64             `json:"gid"` //分组ID
	RouterList []group.RouteInfo `json:"router_list"`
}
