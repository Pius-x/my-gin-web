package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.GET("getUserList", userApi.GetUserList)                // 获取用户列表
		userRouter.GET("getUserListByGid", userApi.GetUserListByGid)      // 获取用户列表通过Gid
		userRouter.POST("createUserInfo", userApi.CreateUserInfo)         // 创建用户信息
		userRouter.POST("updateUserInfo", userApi.UpdateUserInfo)         // 更新用户信息
		userRouter.POST("multiUpdateUserGid", userApi.MultiUpdateUserGid) // 批量更新用户Gid
		userRouter.POST("changePassword", userApi.ChangePassword)         // 修改用户密码
		userRouter.POST("updateHeadPic", userApi.UpdateHeadPic)           // 修改用户头像
		userRouter.POST("deleteUser", userApi.DeleteUser)                 // 删除用户
		userRouter.POST("resetPassword", userApi.ResetPassword)           // 重置密码
	}
}
