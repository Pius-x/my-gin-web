package service

import (
	"github.com/my-gin-web/model/common"
	"github.com/my-gin-web/model/user"
	"github.com/my-gin-web/utils"
	"github.com/pkg/errors"
	"time"
)

//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: err error, userInter model.SysUser

type UserService struct{}

func (This *UserService) CreateUserInfo(r user.CreateUserInfo) (err error) {
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

	userModel.InsertNewUser(createInfo)
	return err
}

// MultiUpdateUserGid 批量更新用户GID
func (This *UserService) MultiUpdateUserGid(userInfo user.MultiUpdateUserGid) {
	for _, item := range userInfo.UserGidList {
		userModel.UpdateUserGid(item.Id, item.Gid)
	}
}

// UpdateUserInfo 更新用户信息
func (This *UserService) UpdateUserInfo(r user.UpdateUserInfo) {
	updateInfo := map[string]any{
		"id":     r.Id,
		"gid":    r.Gid,
		"name":   r.Name,
		"mobile": r.Mobile,
	}

	userModel.UpdateUserInfo(updateInfo)
}

// ChangePassword 修改密码
func (This *UserService) ChangePassword(u user.ChangePasswordStruct) (err error) {

	var oldPwd = userModel.GetUserPwd(u.Id)
	if ok := utils.BcryptCheck(u.Password, oldPwd); !ok {
		return errors.New("原密码错误")
	}

	userModel.UpdateUserPwd(u.Id, utils.BcryptHash(u.NewPassword))

	return nil
}

// UpdateHeadPic 更新用户头像
func (This *UserService) UpdateHeadPic(u user.UpdateHeadPicStruct) {

	userModel.UpdateUserHeadPic(u.Id, u.HeadPic)
}

// GetUserInfoList 分页获取用户列表
func (This *UserService) GetUserInfoList(info user.ReqUserList) (list []user.ResUserList, total int64) {
	limit := int64(10)
	offset := int64(0)
	if info.Keyword == "" {
		limit = info.PageSize
		offset = info.PageSize * (info.Page - 1)
	}

	groupService := GroupService{}
	var gidList = groupService.GetChildrenIdListByGid(info.Gid)

	var tUsers []user.TUsers
	var userList []user.ResUserList

	if info.Keyword == "" {
		userModel.GetPageUserInfo(&tUsers, &total, gidList, limit, offset)
	} else {
		userModel.GetPageUserInfoByKey(&tUsers, &total, gidList, info.Keyword, limit, offset)
	}

	for _, oneUser := range tUsers {
		var oneUserInfo user.ResUserList
		oneUserInfo.Id = oneUser.Id
		oneUserInfo.Account = oneUser.Account
		oneUserInfo.Gid = oneUser.Gid
		oneUserInfo.Name = oneUser.Name
		oneUserInfo.Mobile = oneUser.Mobile
		oneUserInfo.CreateTime = time.Unix(oneUser.CreateTime, 0).Format("2006-01-02 15:04:05")
		oneUserInfo.CreateBy = oneUser.CreateBy
		userList = append(userList, oneUserInfo)
	}

	return userList, total
}

// GetUserListByGid 通过gid获取用户列表
func (This *UserService) GetUserListByGid(info common.GetByGid) (list []user.ResUserList, total int64) {
	tUsers, total := userModel.GetUserListByGid(info.Gid)

	for _, oneUser := range tUsers {
		var oneUserInfo user.ResUserList
		oneUserInfo.Id = oneUser.Id
		oneUserInfo.Account = oneUser.Account
		oneUserInfo.Gid = oneUser.Gid
		oneUserInfo.Name = oneUser.Name
		oneUserInfo.Mobile = oneUser.Mobile
		oneUserInfo.CreateTime = time.Unix(oneUser.CreateTime, 0).Format("2006-01-02 15:04:05")
		oneUserInfo.CreateBy = oneUser.CreateBy
		list = append(list, oneUserInfo)
	}

	return list, total
}

// DeleteUser 删除用户
func (This *UserService) DeleteUser(id int64) {

	userModel.DeleteUser(id)
}

// ResetPassword 修改用户密码
func (This *UserService) ResetPassword(id int64) {

	userModel.UpdateUserPwd(id, "123456")
}
