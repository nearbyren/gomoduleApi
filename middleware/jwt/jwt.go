package jwt

import (
	"github.com/nearbyren/gomodule/consts"
	"github.com/gin-gonic/gin"
	"github.com/nearbyren/gomodule/pkg/util"
	"net/http"
	"time"
)

//校验token
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = consts.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = consts.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = consts.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = consts.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != consts.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  consts.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}
		//让程序继续走下去
		c.Next()
	}
}
