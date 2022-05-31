package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/my-gin-web/global"
	"github.com/my-gin-web/model/common/response"
	systemReq "github.com/my-gin-web/model/system/request"
	systemRes "github.com/my-gin-web/model/system/response"
	"github.com/my-gin-web/model/user"
	"github.com/my-gin-web/utils"
	"go.uber.org/zap"
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
	_ = c.ShouldBindJSON(&l)
	if err := utils.Verify(l, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err, oneUser := BaseService.Login(l); err != nil {
		global.ZapLog.Error("登陆失败! "+err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		b.tokenNext(c, *oneUser)
	}
}

// 登录以后签发jwt
func (b *BaseApi) tokenNext(c *gin.Context, user user.LoginInfo) {
	j := &utils.JWT{SigningKey: []byte(global.Config.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		ID:       user.Id,
		Account:  user.Account,
		Password: user.Password,
	})

	token, err := j.CreateToken(claims)
	if err != nil {
		global.ZapLog.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	fmt.Printf("user %+v \n", user)
	if !global.Config.System.UseMultipoint {
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}

	response.OkWithDetailed(systemRes.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)
}
