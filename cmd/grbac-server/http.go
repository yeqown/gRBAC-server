package main

import (
	"fmt"
	"os"

	_authd "github.com/yeqown/gRBAC-server/internal-modules/auth/delivery"
	_permd "github.com/yeqown/gRBAC-server/internal-modules/permission/delivery"
	_roled "github.com/yeqown/gRBAC-server/internal-modules/role/delivery"
	_userd "github.com/yeqown/gRBAC-server/internal-modules/user/delivery"
	"github.com/yeqown/gRBAC-server/pkg/logger"
	"github.com/yeqown/gRBAC-server/pkg/mw"

	"github.com/gin-gonic/gin"
	"github.com/yeqown/infrastructure/framework/ginic"
)

// StartHTTP ...
func StartHTTP(port int) {
	r := gin.New()

	// import middleware
	r.Use(ginic.LogRequest(logger.Logger, false))
	r.Use(ginic.Recovery(os.Stdout))

	// no need token to request
	r.POST("/admin/verify", _authd.Verify)

	// Token middleware
	r.Use(mw.Token())
	r.POST("/auth", _authd.Auth)

	permGroup := r.Group("/perm")
	{
		permGroup.POST("/new", _permd.NewPermission)
		permGroup.POST("/edit", _permd.EditPermission)
		permGroup.GET("/list", _permd.QueryPermission)
	}

	roleGroup := r.Group("/role")
	{
		roleGroup.POST("/new", _roled.NewRole)
		roleGroup.GET("/list", _roled.QueryRole)
		roleGroup.POST("/assign_perm", _roled.AssignRolePermission)
		roleGroup.POST("/revoke_perm", _roled.DelRolePermission)
	}

	userGroup := r.Group("/user")
	{
		userGroup.POST("/new", _userd.NewUser)
		userGroup.GET("/list", _userd.QueryUser)
		userGroup.POST("/assign_role", _userd.AssignUserPermission)
		userGroup.POST("/revoke_role", _userd.DelUserPermission)
	}

	r.Run(fmt.Sprintf(":%d", port))
}
