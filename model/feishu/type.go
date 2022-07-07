package feishu

import "github.com/my-gin-web/model/user"

type AccessReq struct {
	AccessToken      string `json:"access_token"`       // 飞书服务器授权的access_token，用于调用其他接口
	TokenType        string `json:"token_type"`         // OAuth 2.0协议规定的Token类型，固定为 Bearer
	ExpiresIn        int    `json:"expires_in"`         // access_token 的有效期，三方应用服务器需要根据此返回值来控制access_token的有效时间
	RefreshToken     string `json:"refresh_token"`      // 当 access_token 过期时，通过 refresh_token来刷新，获取新的 access_token
	RefreshExpiresIn int    `json:"refresh_expires_in"` // refresh_token 的有效期
}

type FsUserInfo struct {
	FsId         int64  `json:"fs_id" db:"fs_id"`                 // 飞书ID序号，(递增)
	Sub          string `json:"sub" db:"sub"`                     // 用户在应用内的唯一标识，等同于open_id
	Name         string `json:"name" db:"name"`                   // 用户姓名
	Picture      string `json:"picture" db:"picture"`             // 用户头像，等同于avatar_url
	OpenId       string `json:"open_id" db:"open_id"`             // 用户在应用内的唯一标识, 等同于sub
	UnionId      string `json:"union_id" db:"union_id"`           // 用户统一ID，在同一租户开发的所有应用内的唯一标识
	EnName       string `json:"en_name" db:"en_name"`             // 用户英文名称
	TenantKey    string `json:"tenant_key" db:"tenant_key"`       // 当前企业标识
	AvatarUrl    string `json:"avatar_url" db:"avatar_url"`       // 用户头像，等同于picture
	AvatarThumb  string `json:"avatar_thumb" db:"avatar_thumb"`   // 用户头像 72x72
	AvatarMiddle string `json:"avatar_middle" db:"avatar_middle"` // 用户头像 240x240
	AvatarBig    string `json:"avatar_big" db:"avatar_big"`       // 用户头像 640x640
	UserId       string `json:"user_id" db:"user_id"`             // 用户 user id，申请了邮箱获取权限(获取用户 user Id)的应用会返回该字段
	Email        string `json:"email" db:"email"`                 // 用户邮箱，申请了邮箱获取权限(获取用户邮箱信息)的应用会返回该字段
	Mobile       string `json:"mobile" db:"mobile"`               // 用户手机号，申请了手机号获取权限(获取用户手机号)的应用会返回该字段
	CmsUserId    int64  `json:"cms_user_id" db:"cms_user_id"`     // 管理后台关联的用户ID
}

func (FsUserInfo) TableName() string {
	return "fs_user_info"
}

type LoginUserInfo struct {
	FsUserInfo
	user.TUsers
}

type LoginU struct {
	Test      string
	JWT       string
	ExpiresAt int64
}

type LoginE struct {
	Err string
}
