package mw

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yeqown/gRBAC-server/pkg/secret"
	"github.com/yeqown/infrastructure/types/codes"
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
		if token != secret.GetSecret() {
			c.JSON(http.StatusOK, codes.New(codes.CodeSystemErr, "Token invalid"))
			c.Abort()
			return
		}

		// continue called
		c.Next()
	}
}
