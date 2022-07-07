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
func (Api *UserApi) CreateUserInfo(c *gin.Context) {

	var r user.CreateUserInfo
	utils.Verify(&r, utils.CreateUserInfoVerify, c)

	r.CreateBy = utils.GetOperatorAccount(c)
	if r.Gid == 0 && r.CreateBy != "root" {
		answer.FailWithMessage("请选择用户分组", c)
	}

	if err := UserService.CreateUserInfo(r); err != nil {
		answer.FailWithMessage(err.Error(), c)
	}

	answer.OkWithMessage("注册成功", c)
}

// UpdateUserInfo 更新用户信息
func (Api *UserApi) UpdateUserInfo(c *gin.Context) {

	var r user.UpdateUserInfo
	utils.Verify(&r, utils.UpdateUserInfoVerify, c)

	// 更新用户信息
	UserService.UpdateUserInfo(r)

	answer.OkWithMessage("更新信息成功", c)
}

// MultiUpdateUserGid 批量更新用户Gid
func (Api *UserApi) MultiUpdateUserGid(c *gin.Context) {

	var userInfoList user.MultiUpdateUserGid
	utils.Verify(&userInfoList, utils.MultiUpdateUserGidVerify, c)

	UserService.MultiUpdateUserGid(userInfoList)

	answer.OkWithMessage("批量更新所属分组成功", c)
}

// ChangePassword 修改用户密码
func (Api *UserApi) ChangePassword(c *gin.Context) {

	var oneUser user.ChangePasswordStruct
	utils.Verify(&oneUser, utils.ChangePasswordVerify, c)

	if err := UserService.ChangePassword(oneUser); err != nil {
		answer.FailWithMessage(err.Error(), c)
	}

	answer.OkWithMessage("修改成功", c)
}

// UpdateHeadPic 修改用户头像
func (Api *UserApi) UpdateHeadPic(c *gin.Context) {

	var oneUser user.UpdateHeadPicStruct
	utils.Verify(&oneUser, utils.UpdateHeadPicVerify, c)

	UserService.UpdateHeadPic(oneUser)

	answer.OkWithMessage("修改成功", c)
}

// GetUserList 获取用户列表
func (Api *UserApi) GetUserList(c *gin.Context) {

	var userListInfo user.ReqUserList
	utils.Verify(&userListInfo, utils.UserInfoVerify, c)

	account := utils.GetOperatorAccount(c)

	var gidList = userListInfo.FilterGidList
	if len(gidList) == 0 {
		gidList = GroupService.GetChildrenIdListByGid(userListInfo.Gid, account)
	}

	list, total := UserService.GetUserInfoList(userListInfo, gidList)

	answer.OkWithDetailed(common.PageResult{
		List:  list,
		Total: total,
	}, "获取成功", c)
}

// GetUserListByGid 获取用户列表通过Gid
func (Api *UserApi) GetUserListByGid(c *gin.Context) {

	var userListInfo common.GetByGid
	utils.Verify(&userListInfo, utils.Rules{}, c)

	list, total := UserService.GetUserListByGid(userListInfo)

	answer.OkWithDetailed(common.PageResult{
		List:  list,
		Total: total,
	}, "获取成功", c)
}

// DeleteUser 通过id删除用户
func (Api *UserApi) DeleteUser(c *gin.Context) {

	var reqId common.GetById
	utils.Verify(&reqId, utils.IdVerify, c)

	if utils.GetOperatorID(c) == reqId.Id {
		answer.FailWithMessage("删除用户失败，不能删除自己", c)
	}

	UserService.DeleteUser(reqId.Id)

	answer.OkWithMessage("删除成功", c)
}

// ResetPassword 重置密码
func (Api *UserApi) ResetPassword(c *gin.Context) {

	var reqId common.GetById
	utils.Verify(&reqId, utils.IdVerify, c)

	UserService.ResetPassword(reqId.Id)

	answer.OkWithMessage("重置成功", c)
}

// UnBindFs 用户飞书ID解绑
func (Api *UserApi) UnBindFs(c *gin.Context) {
	var reqId common.GetById
	utils.Verify(&reqId, utils.IdVerify, c)

	UserService.UnbindFs(reqId.Id)

	answer.OkWithMessage("解绑成功", c)
}

// GetUserInfoById 获取用户信息
func (Api *UserApi) GetUserInfoById(c *gin.Context) {
	var reqId common.GetById
	utils.Verify(&reqId, utils.IdVerify, c)

	gid, err := UserService.GetUserInfoById(reqId.Id)

	routerList := GroupService.GetGroupRouter(gid)

	if err != nil {
		answer.FailWithMessage("获取信息失败", c)
	} else {
		answer.OkWithDetailed(user.RefreshUserInfo{
			Gid:        gid,
			RouterList: routerList,
		}, "获取成功", c)
	}
}
