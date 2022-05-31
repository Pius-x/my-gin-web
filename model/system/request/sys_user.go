package request

// User login structure
type Login struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	//Captcha   string `json:"captcha"`   // 验证码
	//CaptchaId string `json:"captchaId"` // 验证码ID
}

// Modify password structure
type ChangePasswordStruct struct {
	Id          int64  `form:"id"`          // 用户ID
	Password    string `form:"password"`    // 原密码
	NewPassword string `form:"newPassword"` // 新密码
}

// Modify password structure
type UpdateHeadPicStruct struct {
	Id      int64 `form:"id"`      // 用户ID
	HeadPic int64 `form:"headPic"` // 头像ID
}

// Modify  user's auth structure
type SetUserAuth struct {
	AuthorityId string `json:"authorityId"` // 角色ID
}

// Modify  user's auth structure
type SetUserAuthorities struct {
	ID           int64
	AuthorityIds []string `json:"authorityIds"` // 角色ID
}

type UpdateUserInfo struct {
	Id     int64  `json:"id"`
	Gid    int64  `json:"gid"`    //分组ID
	Name   string `json:"name"`   //用户昵称
	Mobile string `json:"mobile"` //手机号码
}

type UpdateUserGid struct {
	Id  int64 `json:"id,omitempty"`
	Gid int64 `json:"gid,omitempty"`
}

type MultiUpdateUserGid struct {
	UserGidList []UpdateUserGid `json:"userGidList"` //用户信息列表
}

type CreateUserInfo struct {
	Id         int64  `json:"id"`
	Account    string `json:"account"`     //账号
	Gid        int64  `json:"gid"`         //分组ID
	Name       string `json:"name"`        //用户昵称
	Mobile     string `json:"mobile"`      //手机号码
	CreateTime int64  `json:"create_time"` //创建时间
	CreateBy   string `json:"create_by"`   //创建人
}
