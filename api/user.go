package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my-gin-web/global"
	"github.com/my-gin-web/model/common/request"
	"github.com/my-gin-web/model/common/response"
	systemReq "github.com/my-gin-web/model/system/request"
	"github.com/my-gin-web/model/user"
	"github.com/my-gin-web/utils"
	"go.uber.org/zap"
)

type UserApi struct{}

// CreateUserInfo 创建用户信息
func (b *UserApi) CreateUserInfo(c *gin.Context) {
	var r systemReq.CreateUserInfo
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.CreateUserInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	r.CreateBy = utils.GetOperatorAccount(c)

	if err := UserService.CreateUserInfo(r); err != nil {
		global.ZapLog.Error("注册失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("注册成功", c)
	}
}

// UpdateUserInfo 更新用户信息
func (b *UserApi) UpdateUserInfo(c *gin.Context) {
	var r systemReq.UpdateUserInfo
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.UpdateUserInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := UserService.UpdateUserInfo(r); err != nil {
		global.ZapLog.Error("设置失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}

// MultiUpdateUserGid 批量更新用户Gid
func (b *UserApi) MultiUpdateUserGid(c *gin.Context) {
	var userInfoList systemReq.MultiUpdateUserGid
	_ = c.ShouldBindJSON(&userInfoList)
	if err := utils.Verify(userInfoList, utils.MultiUpdateUserGidVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := UserService.MultiUpdateUserGid(userInfoList); err != nil {
		global.ZapLog.Error("批量更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("批量更新成功", c)
	}
}

// ChangePassword 修改用户密码
func (b *UserApi) ChangePassword(c *gin.Context) {
	var oneUser systemReq.ChangePasswordStruct
	_ = c.ShouldBind(&oneUser)
	if err := utils.Verify(oneUser, utils.ChangePasswordVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := UserService.ChangePassword(oneUser); err != nil {
		global.ZapLog.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// UpdateHeadPic 修改用户头像
func (b *UserApi) UpdateHeadPic(c *gin.Context) {
	var oneUser systemReq.UpdateHeadPicStruct
	_ = c.ShouldBind(&oneUser)
	if err := utils.Verify(oneUser, utils.UpdateHeadPicVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := UserService.UpdateHeadPic(oneUser); err != nil {
		global.ZapLog.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// GetUserList 获取用户列表
func (b *UserApi) GetUserList(c *gin.Context) {
	var userListInfo request.UserList
	_ = c.ShouldBind(&userListInfo)

	if err := utils.Verify(userListInfo, utils.UserInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := UserService.GetUserInfoList(userListInfo); err != nil {
		global.ZapLog.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, "获取成功", c)
	}
}

// GetUserListByGid 获取用户列表通过Gid
func (b *UserApi) GetUserListByGid(c *gin.Context) {
	var userListInfo request.UserListByGid
	_ = c.ShouldBind(&userListInfo)

	if list, total, err := UserService.GetUserListByGid(userListInfo); err != nil {
		global.ZapLog.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, "获取成功", c)
	}
}

// DeleteUser 通过id删除用户
func (b *UserApi) DeleteUser(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBind(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	jwtId := utils.GetOperatorID(c)
	if jwtId == reqId.ID {
		response.FailWithMessage("删除用户失败，不能删除自己", c)
		return
	}

	if err := UserService.DeleteUser(reqId.ID); err != nil {
		global.ZapLog.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// ResetPassword 重置密码
func (b *UserApi) ResetPassword(c *gin.Context) {
	var oneUser user.TUsers
	_ = c.ShouldBindJSON(&oneUser)
	if err := UserService.ResetPassword(oneUser.Id); err != nil {
		global.ZapLog.Error("重置失败!", zap.Error(err))
		response.FailWithMessage("重置失败"+err.Error(), c)
	} else {
		response.OkWithMessage("重置成功", c)
	}
}
