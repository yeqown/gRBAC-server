package delivery

import (
	"github.com/yeqown/gRBAC-server/internal-modules/common"
	"github.com/yeqown/gRBAC-server/pkg/logger"

	"github.com/gin-gonic/gin"
	auth "github.com/ne7ermore/gRBAC"
	"github.com/ne7ermore/gRBAC/services"
	"github.com/yeqown/infrastructure/types/codes"
)

/*
 * 增加权限
 */
type newPermissionForm struct {
	Desc string `form:"permission_desc"`
	Name string `form:"permission_name"`
}

type newPermissionResp struct {
	codes.Proto
	Permission *services.Permission `json:"Permission"`
}

// NewPermission ...
func NewPermission(c *gin.Context) {
	var (
		form = new(newPermissionForm)
		resp = new(newPermissionResp)
	)

	if err := c.ShouldBind(form); err != nil {
		common.ResponseError(c, err)
		return
	}

	permission, err := auth.CreatePermisson(form.Name, form.Desc)
	if err != nil {
		codes.Fill(resp, codes.New(codes.CodeSystemErr, err.Error()))
		logger.Logger.Error(err.Error())
		common.Response(c, resp)
		return
	}

	codes.Fill(resp, codes.New(codes.CodeOK, ""))
	resp.Permission = permission
	common.Response(c, resp)
	return
}

/*
 * 获取权限
 */

type queryPermissionForm struct {
	Limit int `form:"limit;default=10"`
	Skip  int `form:"skip;default=0"`
}

type queryPermissionResp struct {
	codes.Proto
	Permissions []*services.Permission `json:"permissions"`
}

// QueryPermission query all permissions
func QueryPermission(c *gin.Context) {

	var (
		form = new(queryPermissionForm)
		resp = new(queryPermissionResp)
	)

	if err := c.ShouldBind(form); err != nil {
		common.ResponseError(c, err)
		return
	}

	mapPerms, err := auth.GetAllPerms(form.Skip, form.Limit, common.CreateTimeDesc)
	if err != nil {
		codes.Fill(resp, codes.New(codes.CodeSystemErr, err.Error()))
		common.Response(c, resp)
		return
	}

	// for _, perm := range mapPerms {
	// 	println(perm)
	// 	// logger.Logger.Info(k, )
	// 	// logger.Logger.Info(perm.(*services.Permission))
	// }

	codes.Fill(resp, codes.New(codes.CodeOK, ""))
	resp.Permissions = mapPerms
	common.Response(c, resp)
	return
}

/*
 * 编辑权限
 */

type editPermissionForm struct {
	PermissionID string `form:"permission_id"`
	Desc         string `form:"desc"`
	Name         string `form:"name"`
}

type editPermissionResp struct {
	codes.Proto
	Permission *services.Permission `json:"permission"`
}

// EditPermission ...
func EditPermission(c *gin.Context) {
	var (
		form = new(editPermissionForm)
		resp = new(editPermissionResp)
	)

	if err := c.ShouldBind(form); err != nil {
		common.ResponseError(c, err)
		return
	}

	permission, err := auth.UpdatePerm(form.PermissionID, form.Desc, form.Name)
	if err != nil {
		codes.Fill(resp, codes.New(codes.CodeSystemErr, err.Error()))
		logger.Logger.Error(err.Error())
		common.Response(c, resp)
		return
	}

	codes.Fill(resp, codes.New(codes.CodeOK, ""))
	resp.Permission = permission
	common.Response(c, resp)
	return
}
