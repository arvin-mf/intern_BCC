package middleware

import (
	"errors"
	"intern_BCC/model"
	sdk_jwt "intern_BCC/sdk/jwt"
	"intern_BCC/sdk/response"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if !strings.HasPrefix(authorization, "Bearer ") {
			c.Abort()
			msg := "wrong header value"
			response.FailOrError(c, http.StatusForbidden, msg, errors.New(msg))
			return
		}
		tokenJwt := authorization[7:]
		claims := model.UserClaims{}
		jwtKey := os.Getenv("secret_key")
		if err := sdk_jwt.DecodeToken(tokenJwt, &claims, jwtKey); err != nil {
			c.Abort()
			response.FailOrError(c, http.StatusUnauthorized, "unauthorized", err)
			return
		}
		c.Set("user", claims)
	}
}
