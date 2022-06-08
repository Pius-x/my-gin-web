package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my-gin-web/model/common"
	"github.com/my-gin-web/model/group"
	"github.com/my-gin-web/utils"
	"github.com/my-gin-web/utils/answer"
)

type GroupApi struct{}

// CreateGroup 创建分组
func (Api *GroupApi) CreateGroup(c *gin.Context) {

	var authority group.TGroups
	utils.Verify(&authority, utils.CreateGroupVerify, c)

	// 创建新分组
	GroupService.CreateGroup(authority)

	answer.OkWithMessage("创建成功", c)
}

// DeleteGroup 删除分组
func (Api *GroupApi) DeleteGroup(c *gin.Context) {

	var authority group.TGroups
	utils.Verify(&authority, utils.AuthorityIdVerify, c)

	// 删除角色之前需要判断是否有用户正在使用此角色
	if err := GroupService.DeleteGroup(&authority); err != nil {
		answer.FailWithMessage(err.Error(), c)
	}

	answer.OkWithMessage("删除成功", c)
}

// UpdateGroupRouter 更新分组路由列表
func (Api *GroupApi) UpdateGroupRouter(c *gin.Context) {

	var auth group.GroupInfo
	utils.Verify(&auth, utils.UpdateGroupRouterVerify, c)

	//判断权限是否超过了父分组的权限
	if auth.ParentGid != 0 {
		GroupService.ParentGroupRouterListVerify(auth)
	}

	//更新自身以及子分组的权限
	GroupService.UpdateGroupRouterList(auth)

	answer.OkWithMessage("更新成功", c)
}

// UpdateGroup 更新分组
func (Api *GroupApi) UpdateGroup(c *gin.Context) {

	var auth group.TGroups
	utils.Verify(&auth, utils.AuthorityVerify, c)

	GroupService.UpdateGroup(auth)

	answer.OkWithMessage("更新成功", c)
}

// GetGroupList 获取分组列表
func (Api *GroupApi) GetGroupList(c *gin.Context) {

	var params common.GetByGid
	utils.Verify(&params, utils.Rules{}, c)

	list := GroupService.GetGroupInfoList(params.Gid)

	answer.OkWithDetailed(common.PageResult{
		List: list,
	}, "获取成功", c)
}
