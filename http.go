package main

import (
	"github.com/yeqown/gRBAC-server/middleware"
	"os"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yeqown/gRBAC-server/controllers"
	"github.com/yeqown/gRBAC-server/logger"
	"github.com/yeqown/server-common/framework/ginic"
)

// StartHTTP ...
func StartHTTP(port int) {
	r := gin.New()

	// import middleware
	r.Use(ginic.LogRequest(logger.Logger, false))
	r.Use(ginic.Recovery(os.Stdout))

	// no need token to request
	r.POST("/admin/verify", controllers.Verify)

	// Token middleware
	r.Use(middleware.Token())
	r.POST("/auth", controllers.Auth)

	permGroup := r.Group("/perm")
	{
		permGroup.POST("/new", controllers.NewPermission)
		permGroup.POST("/edit", controllers.EditPermission)
		permGroup.GET("/list", controllers.QueryPermission)
	}

	roleGroup := r.Group("/role")
	{
		roleGroup.POST("/new", controllers.NewRole)
		roleGroup.GET("/list", controllers.QueryRole)
		roleGroup.POST("/assign_perm", controllers.AssignRolePermission)
		roleGroup.POST("/revoke_perm", controllers.DelRolePermission)
	}

	userGroup := r.Group("/user")
	{
		userGroup.POST("/new", controllers.NewUser)
		userGroup.GET("/list", controllers.QueryUser)
		userGroup.POST("/assign_role", controllers.AssignUserPermission)
		userGroup.POST("/revoke_role", controllers.DelUserPermission)
	}

	r.Run(fmt.Sprintf(":%d", port))
}
