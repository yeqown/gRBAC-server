package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yeqown/gRBAC-server/services"
	"github.com/yeqown/server-common/code"
)

// Token middleware need each request has
// token that taken with the request
func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		if c.Request.Method == http.MethodGet {
			token = c.Query("token")
		} else {
			token = c.PostForm("token")
		}

		// empty
		if token != services.GetToken() {
			c.JSON(http.StatusOK, code.NewCodeInfo(code.CodeSystemErr, "Token invalid"))
			c.Abort()
			return
		}

		// continue called
		c.Next()
	}
}
