package service

import (
	"encoding/json"
	"errors"
	"github.com/my-gin-web/global"
	"github.com/my-gin-web/model/group"
	systemReq "github.com/my-gin-web/model/system/request"
	"github.com/my-gin-web/model/user"
	"github.com/my-gin-web/utils"
)

type BaseService struct{}

func (This *BaseService) Login(u systemReq.Login) (err error, userInter *user.LoginInfo) {

	//获取用户个人信息
	var userInfo user.TUsers
	if err = userModel.GetOneUserInfo(&userInfo, u); err != nil {
		global.ZapLog.Error(err.Error())
		return errors.New("账号不存在"), nil
	}

	//密码校验
	if ok := utils.BcryptCheck(u.Password, userInfo.Password); !ok {
		return errors.New("密码错误"), nil
	}

	//登录信息
	loginInfo := user.LoginInfo{
		TUsers:     userInfo,
		RouterList: []group.RouteInfo{},
	}

	var groupInfoString = groupModel.GetRouterList(userInfo.Gid)

	_ = json.Unmarshal([]byte(groupInfoString), &loginInfo.RouterList)

	//更新最后一次登录时间
	if err = userModel.UpdateLastLoginTime(loginInfo.Account); err != nil {
		return err, nil
	}

	return err, &loginInfo
}
