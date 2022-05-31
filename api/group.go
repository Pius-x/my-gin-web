package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my-gin-web/global"
	"github.com/my-gin-web/model/common/request"
	"github.com/my-gin-web/model/common/response"
	"github.com/my-gin-web/model/group"
	"github.com/my-gin-web/utils"
	"go.uber.org/zap"
)

type GroupApi struct{}

// CreateGroup 创建分组
func (a *GroupApi) CreateGroup(c *gin.Context) {
	var authority group.TGroups
	_ = c.ShouldBindJSON(&authority)
	if err := utils.Verify(authority, utils.CreateGroupVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := GroupService.CreateGroup(authority); err != nil {
		global.ZapLog.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {

		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGroup 删除分组
func (a *GroupApi) DeleteGroup(c *gin.Context) {
	var authority group.TGroups
	_ = c.ShouldBindJSON(&authority)
	if err := utils.Verify(authority, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := GroupService.DeleteGroup(&authority); err != nil { // 删除角色之前需要判断是否有用户正在使用此角色
		global.ZapLog.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// UpdateGroupRouter 更新分组路由列表
func (a *GroupApi) UpdateGroupRouter(c *gin.Context) {
	var auth group.GroupInfo
	_ = c.ShouldBindJSON(&auth)

	if err := utils.Verify(auth, utils.UpdateGroupRouterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//判断权限是否超过了父分组的权限
	if auth.ParentGid != 0 {
		if err := GroupService.ParentGroupRouterListVerify(auth); err != nil {
			global.ZapLog.Error("配置的权限不能超过其父分组的权限!", zap.Error(err))
			response.FailWithMessage("更新失败 "+err.Error(), c)
			return
		}
	}

	//更新自身以及子分组的权限
	if err := GroupService.UpdateGroupRouterList(auth); err != nil {
		global.ZapLog.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// UpdateGroup 更新分组
func (a *GroupApi) UpdateGroup(c *gin.Context) {
	var auth group.TGroups
	_ = c.ShouldBind(&auth)
	if err := utils.Verify(auth, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := GroupService.UpdateGroup(auth); err != nil {
		global.ZapLog.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetGroupList 获取分组列表
func (a *GroupApi) GetGroupList(c *gin.Context) {
	var params request.GetGroupListById
	_ = c.ShouldBind(&params)

	if err, list := GroupService.GetGroupInfoList(params.Gid); err != nil {
		global.ZapLog.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}
}
