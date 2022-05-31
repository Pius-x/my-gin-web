package service

import (
	"github.com/my-gin-web/model/group"
	"github.com/my-gin-web/model/operationlog"
	"github.com/my-gin-web/model/user"
)

var (
	userModel         *user.Model
	groupModel        *group.Model
	operationLogModel *operationlog.Model
)

func InitModel() {
	userModel = user.NewModel()
	groupModel = group.NewModel()
	operationLogModel = operationlog.NewModel()
}
