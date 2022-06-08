package service

import (
	"encoding/json"
	"github.com/my-gin-web/global"
	"github.com/my-gin-web/model/group"
	"github.com/my-gin-web/model/user"
	"github.com/my-gin-web/utils"
	"github.com/pkg/errors"
)

type BaseService struct{}

func (This *BaseService) Login(u user.ReqLogin) (err error, res user.ResLogin) {

	// 获取用户个人信息
	userInfo, err := userModel.GetOneUserInfo(u)
	if err != nil {
		return errors.New("账号不存在"), res
	}

	// 密码校验
	if ok := utils.BcryptCheck(u.Password, userInfo.Password); !ok {
		return errors.New("密码错误"), res
	}

	// 登录信息
	var loginInfo = user.LoginInfo{
		TUsers:     userInfo,
		RouterList: []group.RouteInfo{},
	}

	// 解析权限信息
	var groupInfoString = groupModel.GetRouterList(userInfo.Gid)
	if err = json.Unmarshal([]byte(groupInfoString), &loginInfo.RouterList); err != nil {
		panic(errors.Wrap(err, "权限列表解析失败"))
	}

	// 更新最后一次登录时间
	userModel.UpdateLastLoginTime(loginInfo.Id)

	// 登录以后签发jwt
	res = This.tokenNext(loginInfo)

	return err, res
}

// TokenNext 登录以后签发jwt
func (This *BaseService) tokenNext(loginInfo user.LoginInfo) (response user.ResLogin) {

	j := &utils.JWT{SigningKey: []byte(global.Config.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(user.BaseClaims{
		ID:       loginInfo.Id,
		Account:  loginInfo.Account,
		Password: loginInfo.Password,
	})

	token, err := j.CreateToken(claims)
	if err != nil {
		panic(errors.Wrap(err, "获取token失败"))
	}

	return user.ResLogin{
		User:      loginInfo,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}
}
