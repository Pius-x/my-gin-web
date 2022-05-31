package router

import (
	"github.com/gin-gonic/gin"
	"github.com/my-gin-web/middleware"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	{
		userRouter.GET("getUserList", UserApi.GetUserList)                // 获取用户列表
		userRouter.GET("getUserListByGid", UserApi.GetUserListByGid)      // 获取用户列表通过Gid
		userRouter.POST("createUserInfo", UserApi.CreateUserInfo)         // 创建用户信息
		userRouter.POST("updateUserInfo", UserApi.UpdateUserInfo)         // 更新用户信息
		userRouter.POST("multiUpdateUserGid", UserApi.MultiUpdateUserGid) // 批量更新用户Gid
		userRouter.POST("changePassword", UserApi.ChangePassword)         // 修改用户密码
		userRouter.POST("updateHeadPic", UserApi.UpdateHeadPic)           // 修改用户头像
		userRouter.POST("deleteUser", UserApi.DeleteUser)                 // 删除用户
		userRouter.POST("resetPassword", UserApi.ResetPassword)           // 重置密码
	}
}
