package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my-gin-web/model/common"
	"github.com/my-gin-web/model/user"
	"github.com/my-gin-web/utils"
	"github.com/my-gin-web/utils/answer"
)

type UserApi struct{}

// CreateUserInfo 创建用户信息
func (b *UserApi) CreateUserInfo(c *gin.Context) {

	var r user.CreateUserInfo
	utils.Verify(&r, utils.CreateUserInfoVerify, c)

	r.CreateBy = utils.GetOperatorAccount(c)

	if err := UserService.CreateUserInfo(r); err != nil {
		answer.FailWithMessage(err.Error(), c)
	}

	answer.OkWithMessage("注册成功", c)
}

// UpdateUserInfo 更新用户信息
func (b *UserApi) UpdateUserInfo(c *gin.Context) {

	var r user.UpdateUserInfo
	utils.Verify(&r, utils.UpdateUserInfoVerify, c)

	// 更新用户信息
	UserService.UpdateUserInfo(r)

	answer.OkWithMessage("更新信息成功", c)
}

// MultiUpdateUserGid 批量更新用户Gid
func (b *UserApi) MultiUpdateUserGid(c *gin.Context) {

	var userInfoList user.MultiUpdateUserGid
	utils.Verify(&userInfoList, utils.MultiUpdateUserGidVerify, c)

	UserService.MultiUpdateUserGid(userInfoList)

	answer.OkWithMessage("批量更新所属分组成功", c)
}

// ChangePassword 修改用户密码
func (b *UserApi) ChangePassword(c *gin.Context) {

	var oneUser user.ChangePasswordStruct
	utils.Verify(&oneUser, utils.ChangePasswordVerify, c)

	if err := UserService.ChangePassword(oneUser); err != nil {
		answer.FailWithMessage(err.Error(), c)
	}

	answer.OkWithMessage("修改成功", c)
}

// UpdateHeadPic 修改用户头像
func (b *UserApi) UpdateHeadPic(c *gin.Context) {

	var oneUser user.UpdateHeadPicStruct
	utils.Verify(&oneUser, utils.UpdateHeadPicVerify, c)

	UserService.UpdateHeadPic(oneUser)

	answer.OkWithMessage("修改成功", c)
}

// GetUserList 获取用户列表
func (b *UserApi) GetUserList(c *gin.Context) {

	var userListInfo user.ReqUserList
	utils.Verify(&userListInfo, utils.UserInfoVerify, c)

	list, total := UserService.GetUserInfoList(userListInfo)

	answer.OkWithDetailed(common.PageResult{
		List:  list,
		Total: total,
	}, "获取成功", c)
}

// GetUserListByGid 获取用户列表通过Gid
func (b *UserApi) GetUserListByGid(c *gin.Context) {

	var userListInfo common.GetByGid
	utils.Verify(&userListInfo, utils.Rules{}, c)

	list, total := UserService.GetUserListByGid(userListInfo)

	answer.OkWithDetailed(common.PageResult{
		List:  list,
		Total: total,
	}, "获取成功", c)
}

// DeleteUser 通过id删除用户
func (b *UserApi) DeleteUser(c *gin.Context) {

	var reqId common.GetById
	utils.Verify(&reqId, utils.IdVerify, c)

	if utils.GetOperatorID(c) == reqId.ID {
		answer.FailWithMessage("删除用户失败，不能删除自己", c)
	}

	UserService.DeleteUser(reqId.ID)

	answer.OkWithMessage("删除成功", c)
}

// ResetPassword 重置密码
func (b *UserApi) ResetPassword(c *gin.Context) {

	var oneUser user.TUsers
	utils.Verify(&oneUser, utils.Rules{}, c)

	UserService.ResetPassword(oneUser.Id)

	answer.OkWithMessage("重置成功", c)
}
