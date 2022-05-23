package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GroupRouter struct{}

func (s *GroupRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	groupRouter := Router.Group("group").Use(middleware.OperationRecord())
	groupApi := v1.ApiGroupApp.SystemApiGroup.GroupApi
	{
		groupRouter.GET("getGroupList", groupApi.GetGroupList)            // 获取分组列表
		groupRouter.POST("updateGroupRouter", groupApi.UpdateGroupRouter) // 更新分组路由列表
		groupRouter.POST("createGroup", groupApi.CreateGroup)             // 创建分组
		groupRouter.POST("deleteGroup", groupApi.DeleteGroup)             // 删除分组
		groupRouter.POST("updateGroup", groupApi.UpdateGroup)             // 更新分组
	}
}
