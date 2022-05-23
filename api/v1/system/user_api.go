package system

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
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

	user := &system.TUsers{
		UserCoreInfo: system.UserCoreInfo{
			Account: r.Account,
			Gid:     r.Gid,
			Name:    r.Name,
			Mobile:  r.Mobile,
		},
		Password: "123456",
		HeadPic:  0,
		CreateBy: utils.GetOperatorAccount(c),
	}

	if err, userReturn := userService.CreateUserInfo(*user); err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
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

	updateUserInfo := map[string]interface{}{
		"gid":    r.Gid,
		"name":   r.Name,
		"mobile": r.Mobile,
	}

	if err := userService.UpdateUserInfo(updateUserInfo, r.Id); err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
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

	if err := userService.MultiUpdateUserGid(userInfoList); err != nil {
		global.GVA_LOG.Error("批量更新失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("批量更新成功", c)
	}
}

// ChangePassword 修改用户密码
func (b *UserApi) ChangePassword(c *gin.Context) {
	var user systemReq.ChangePasswordStruct
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.ChangePasswordVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	u := &system.TUsers{
		UserCoreInfo: system.UserCoreInfo{Account: user.Username},
		Password:     user.Password,
	}
	if err, _ := userService.ChangePassword(u, user.NewPassword); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// UpdateHeadPic 修改用户头像
func (b *UserApi) UpdateHeadPic(c *gin.Context) {
	//account:=c.Request.Header.Get("x-account")
	var user systemReq.UpdateHeadPicStruct
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.UpdateHeadPicVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	u := &system.TUsers{UserCoreInfo: system.UserCoreInfo{Account: user.Account}, HeadPic: user.HeadPic}
	if err, _ := userService.UpdateHeadPic(u, user.HeadPic); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
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
	if err, list, total := userService.GetUserInfoList(userListInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
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
	_ = c.ShouldBindQuery(&userListInfo)

	if err, list, total := userService.GetUserListByGid(userListInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
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
	_ = c.ShouldBindJSON(&reqId)
	if err := utils.Verify(reqId, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	claims := utils.GetOperatorAccount(c)
	fmt.Printf("claims: %+v \n", claims)
	jwtId := utils.GetOperatorID(c)
	if jwtId == reqId.ID {
		response.FailWithMessage("删除用户失败，不能删除自己", c)
		return
	}
	if err := userService.DeleteUser(reqId.ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// ResetPassword 重置密码
func (b *UserApi) ResetPassword(c *gin.Context) {
	var user system.TUsers
	_ = c.ShouldBindJSON(&user)
	if err := userService.ResetPassword(user.Id); err != nil {
		global.GVA_LOG.Error("重置失败!", zap.Error(err))
		response.FailWithMessage("重置失败"+err.Error(), c)
	} else {
		response.OkWithMessage("重置成功", c)
	}
}
