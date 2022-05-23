package system

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	BaseApi
	UserApi
	GroupApi
	OperationRecordApi
}

var (
	BaseService            = service.ServiceGroupApp.SystemServiceGroup.BaseService
	userService            = service.ServiceGroupApp.SystemServiceGroup.UserService
	groupService           = service.ServiceGroupApp.SystemServiceGroup.GroupService
	operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
)
