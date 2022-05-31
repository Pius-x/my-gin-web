package service

import (
	"errors"
	"github.com/my-gin-web/global"
	"github.com/my-gin-web/model/common/request"
	systemReq "github.com/my-gin-web/model/system/request"
	systemRes "github.com/my-gin-web/model/system/response"
	"github.com/my-gin-web/model/user"
	"github.com/my-gin-web/utils"
	"go.uber.org/zap"
	"time"
)

//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: err error, userInter model.SysUser

type UserService struct{}

func (This *UserService) CreateUserInfo(r systemReq.CreateUserInfo) (err error) {
	var curTime = time.Now().Unix()

	createInfo := user.TUsers{
		CoreInfo: user.CoreInfo{
			Account:    r.Account,
			Gid:        r.Gid,
			Name:       r.Name,
			Mobile:     r.Mobile,
			CreateTime: curTime,
		},
		Password:      utils.BcryptHash("123456"),
		HeadPic:       0,
		CreateBy:      r.CreateBy,
		UpdateTime:    curTime,
		LastLoginTime: curTime,
	}
	// 判断用户名是否注册
	if total := userModel.GetUserCountByAccount(createInfo.Account); total != 0 {
		return errors.New("账户名已注册")
	}

	// 判断电话号码是否注册
	if total := userModel.GetUserCountByMobile(createInfo.Mobile); total != 0 {
		return errors.New("手机号码已注册")
	}

	err = userModel.InsertNewUser(createInfo)
	return err
}

// MultiUpdateUserGid 批量更新用户GID
func (This *UserService) MultiUpdateUserGid(userInfo systemReq.MultiUpdateUserGid) (err error) {
	for _, item := range userInfo.UserGidList {
		if err = userModel.UpdateUserGid(item.Id, item.Gid); err != nil {
			return errors.New("批量更新失败")
		}
	}

	return err
}

// UpdateUserInfo 更新用户信息
func (This *UserService) UpdateUserInfo(r systemReq.UpdateUserInfo) error {
	updateInfo := map[string]interface{}{
		"id":     r.Id,
		"gid":    r.Gid,
		"name":   r.Name,
		"mobile": r.Mobile,
	}

	return userModel.UpdateUserInfo(updateInfo)
}

//@function: ChangePassword
//@description: 修改用户密码
//@param: u *model.SysUser, newPassword string
//@return: err error, userInter *model.SysUser

func (This *UserService) ChangePassword(u systemReq.ChangePasswordStruct) (err error) {

	var oldPwd = userModel.GetUserPwd(u.Id)
	if ok := utils.BcryptCheck(u.Password, oldPwd); !ok {
		return errors.New("原密码错误")
	}

	return userModel.UpdateUserPwd(u.Id, utils.BcryptHash(u.NewPassword))
}

//@function: UpdateHeadPic
//@description: 更新头像
//@param: u *model.SysUser, newPassword string
//@return: err error, userInter *model.SysUser

func (This *UserService) UpdateHeadPic(u systemReq.UpdateHeadPicStruct) (err error) {

	return userModel.UpdateUserHeadPic(u.Id, u.HeadPic)
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (This *UserService) GetUserInfoList(info request.UserList) (err error, list []systemRes.UserListResponse, total int64) {
	limit := int64(10)
	offset := int64(0)
	if info.Keyword == "" {
		limit = info.PageSize
		offset = info.PageSize * (info.Page - 1)
	}

	groupService := GroupService{}
	var gidList = groupService.GetChildrenIdListByGid(info.Gid)

	var tUsers []user.TUsers
	var userList []systemRes.UserListResponse

	if info.Keyword == "" {
		if err = userModel.GetPageUserInfo(&tUsers, &total, gidList, limit, offset); err != nil {
			global.ZapLog.Error("get page user info err", zap.Error(err))
			return err, nil, 0
		}
	} else {
		if err = userModel.GetPageUserInfoByKey(&tUsers, &total, gidList, info.Keyword, limit, offset); err != nil {
			global.ZapLog.Error("get page user info by key err", zap.Error(err))
			return err, nil, 0
		}
	}

	for _, oneUser := range tUsers {
		var oneUserInfo systemRes.UserListResponse
		oneUserInfo.Id = oneUser.Id
		oneUserInfo.Account = oneUser.Account
		oneUserInfo.Gid = oneUser.Gid
		oneUserInfo.Name = oneUser.Name
		oneUserInfo.Mobile = oneUser.Mobile
		oneUserInfo.CreateTime = time.Unix(oneUser.CreateTime, 0).Format("2006-01-02 15:04:05")
		oneUserInfo.CreateBy = oneUser.CreateBy
		userList = append(userList, oneUserInfo)
	}

	return err, userList, total
}

// GetUserListByGid 通过gid获取用户列表
func (This *UserService) GetUserListByGid(info request.UserListByGid) (list []systemRes.UserListResponse, total int64, err error) {
	tUsers, total, err := userModel.GetUserListByGid(info.Gid)
	if err != nil {
		return nil, 0, err
	}

	for _, oneUser := range tUsers {
		var oneUserInfo systemRes.UserListResponse
		oneUserInfo.Id = oneUser.Id
		oneUserInfo.Account = oneUser.Account
		oneUserInfo.Gid = oneUser.Gid
		oneUserInfo.Name = oneUser.Name
		oneUserInfo.Mobile = oneUser.Mobile
		oneUserInfo.CreateTime = time.Unix(oneUser.CreateTime, 0).Format("2006-01-02 15:04:05")
		oneUserInfo.CreateBy = oneUser.CreateBy
		list = append(list, oneUserInfo)
	}

	return list, total, err
}

// DeleteUser 删除用户
func (This *UserService) DeleteUser(id int64) (err error) {

	return userModel.DeleteUser(id)
}

// ResetPassword 修改用户密码
func (This *UserService) ResetPassword(id int64) (err error) {

	var pwd = "123456"
	return userModel.UpdateUserPwd(id, pwd)
}
