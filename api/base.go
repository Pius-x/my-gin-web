package api

import (
	"github.com/gin-gonic/gin"
	"github.com/my-gin-web/model/feishu"
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

	answer.OkWithDetailed(BaseService.BuildLoginInfo(resInfo), "登录成功", c)
}

// FsLogin 飞书扫码登录/web登录
func (Api *BaseApi) FsLogin(c *gin.Context) {
	code, _ := c.GetQuery("code")

	// 登录逻辑
	if res, err := BaseService.FsLogin(code); err != nil {
		c.HTML(200, "error.html", feishu.LoginE{Err: err.Error()})
	} else {
		c.HTML(200, "login.html", BaseService.BuildFsLoginInfo(res))
	}
}

// FsBind 用户飞书ID绑定
func (Api *BaseApi) FsBind(c *gin.Context) {
	code, _ := c.GetQuery("code")
	token, _ := c.GetQuery("state")

	// parseToken 解析token包含的信息
	claims, err := utils.NewJWT().ParseToken(token)
	if err != nil || claims == nil {
		c.HTML(200, "error.html", feishu.LoginE{Err: "Token 信息有误，绑定失败"})
	}

	userId := claims.BaseClaims.ID
	if fsUserInfo, err := BaseService.BindFs(code, userId); err != nil {
		c.HTML(200, "error.html", feishu.LoginE{Err: err.Error()})
	} else {
		c.HTML(200, "bindSuccess.html", BaseService.BuildFsBindInfo(fsUserInfo))
	}
}
