package router

import (
	"github.com/gin-gonic/gin"
)

type GroupRouter struct{}

func (s *GroupRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	groupRouter := Router.Group("group")
	{
		groupRouter.GET("getGroupList", GroupApi.GetGroupList)            // 获取分组列表
		groupRouter.POST("updateGroupRouter", GroupApi.UpdateGroupRouter) // 更新分组路由列表
		groupRouter.POST("createGroup", GroupApi.CreateGroup)             // 创建分组
		groupRouter.POST("deleteGroup", GroupApi.DeleteGroup)             // 删除分组
		groupRouter.POST("updateGroup", GroupApi.UpdateGroup)             // 更新分组
	}
}
