package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/my-gin-web/global"
	"github.com/my-gin-web/model/user"
	"go.uber.org/zap"
)

func GetClaims(c *gin.Context) (*user.CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.ZapLog.Error(err.Error(), zap.Error(err))
	}
	return claims, err
}

// GetOperatorID 从Gin的Context中获取从jwt解析出来的用户ID
func GetOperatorID(c *gin.Context) int64 {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.ID
		}
	} else {
		waitUse := claims.(*user.CustomClaims)
		return waitUse.ID
	}
}

// GetOperatorAccount  从Gin的Context中获取从jwt解析出来的用户Account
func GetOperatorAccount(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.Account
		}
	} else {
		waitUse := claims.(*user.CustomClaims)
		return waitUse.Account
	}
}

//// 从Gin的Context中获取从jwt解析出来的用户UUID
//func GetUserUuid(c *gin.Context) uuid.UUID {
//	if claims, exists := c.Get("claims"); !exists {
//		if cl, err := GetClaims(c); err != nil {
//			return uuid.UUID{}
//		} else {
//			return cl.UUID
//		}
//	} else {
//		waitUse := claims.(*systemReq.CustomClaims)
//		return waitUse.UUID
//	}
//}
//
//// 从Gin的Context中获取从jwt解析出来的用户角色id
//func GetUserAuthorityId(c *gin.Context) string {
//	if claims, exists := c.Get("claims"); !exists {
//		if cl, err := GetClaims(c); err != nil {
//			return ""
//		} else {
//			return cl.AuthorityId
//		}
//	} else {
//		waitUse := claims.(*systemReq.CustomClaims)
//		return waitUse.AuthorityId
//	}
//}

// 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(c *gin.Context) *user.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*user.CustomClaims)
		return waitUse
	}
}
