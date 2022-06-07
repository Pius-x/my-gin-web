package request

// User login structure
type Login struct {
	Username string `json:"username" form:"username" db:"username"` // 用户名
	Password string `json:"password" form:"password" db:"password"` // 密码
	//Captcha   string `json:"captcha"`   // 验证码
	//CaptchaId string `json:"captchaId"` // 验证码ID
}

// Modify password structure
type ChangePasswordStruct struct {
	Id          int64  `form:"id" json:"id" db:"id"`                            // 用户ID
	Password    string `form:"password" json:"password" db:"password"`          // 原密码
	NewPassword string `form:"newPassword" json:"newPassword" db:"newPassword"` // 新密码
}

// Modify password structure
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
