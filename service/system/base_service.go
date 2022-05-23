package system

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type BaseService struct{}

func (This *BaseService) Login(u systemReq.Login) (err error, userInter *system.LoginInfo) {
	if nil == global.GVA_DB {
		return fmt.Errorf("db not init"), nil
	}

	//获取用户个人信息
	var userInfo system.TUsers
	if err = global.GVA_DB.Where("account = ? or mobile = ?", u.Username, u.Username).First(&userInfo).Error; err != nil {
		return errors.New("账号不存在"), nil
	}

	//密码校验
	fmt.Printf("loginInfo %+v \n", userInfo)
	if ok := utils.BcryptCheck(u.Password, userInfo.Password); !ok {
		return errors.New("密码错误"), nil
	}

	//登录信息
	loginInfo := system.LoginInfo{
		TUsers:     userInfo,
		RouterList: []system.RouteInfo{},
	}

	var groupInfo system.TGroups
	_ = global.GVA_DB.Where("gid = ?", userInfo.Gid).First(&groupInfo).Error

	loginInfo.RouterList = groupInfo.RouterList

	return err, &loginInfo
}
