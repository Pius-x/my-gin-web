package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my-gin-web/model/user"
	"github.com/my-gin-web/utils"
	"github.com/my-gin-web/utils/answer"
)

type BaseApi struct{}

// Login 用户登录
func (Api *BaseApi) Login(c *gin.Context) {
	var req user.ReqLogin
	utils.Verify(&req, utils.LoginVerify, c)

	err, resInfo := BaseService.Login(req)
	if err != nil {
		answer.FailWithMessage(err.Error(), c)
	}

	answer.OkWithDetailed(resInfo, "登录成功", c)
}
