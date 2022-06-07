package service

import (
	"encoding/json"
	"github.com/my-gin-web/global"
	"github.com/my-gin-web/model/group"
	systemReq "github.com/my-gin-web/model/system/request"
	systemRes "github.com/my-gin-web/model/system/response"
	"github.com/my-gin-web/model/user"
	"github.com/my-gin-web/utils"
	"github.com/pkg/errors"
)

type BaseService struct{}

func (This *BaseService) Login(u systemReq.Login) (err error, response systemRes.LoginResponse) {

	// 获取用户个人信息
	var userInfo user.TUsers
	if err = userModel.GetOneUserInfo(&userInfo, u); err != nil {
		return errors.New("账号不存在"), response
	}

	// 密码校验
	if ok := utils.BcryptCheck(u.Password, userInfo.Password); !ok {
		return errors.New("密码错误"), response
	}

	// 登录信息
	var loginInfo = user.LoginInfo{
		TUsers:     userInfo,
		RouterList: []group.RouteInfo{},
	}

	var groupInfoString = groupModel.GetRouterList(userInfo.Gid)

	_ = json.Unmarshal([]byte(groupInfoString), &loginInfo.RouterList)

	// 更新最后一次登录时间
	userModel.UpdateLastLoginTime(loginInfo.Id)

	// 登录以后签发jwt
	response = This.tokenNext(loginInfo)

	return err, response
}

// TokenNext 登录以后签发jwt
func (This *BaseService) tokenNext(user user.LoginInfo) (response systemRes.LoginResponse) {
	j := &utils.JWT{SigningKey: []byte(global.Config.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		ID:       user.Id,
		Account:  user.Account,
		Password: user.Password,
	})

	token, err := j.CreateToken(claims)
	if err != nil {
		panic(errors.Wrap(err, "获取token失败"))
	}

	return systemRes.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}
}
