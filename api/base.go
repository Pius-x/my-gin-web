package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my-gin-web/model/common/response"
	systemReq "github.com/my-gin-web/model/system/request"
	"github.com/my-gin-web/utils"
)

type BaseApi struct{}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body systemReq.Login true "用户名, 密码, 验证码"
// @Success 200 {object} response.Response{data=systemRes.LoginResponse,msg=string} "返回包括用户信息,token,过期时间"
// @Router /base/login [post]

func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	utils.Verify(&l, utils.LoginVerify, c)

	err, responseInfo := BaseService.Login(l)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	response.OkWithDetailed(responseInfo, "登录成功", c)
}
