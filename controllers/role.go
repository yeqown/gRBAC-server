package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	auth "github.com/ne7ermore/gRBAC"
	"github.com/ne7ermore/gRBAC/services"
	"github.com/yeqown/gRBAC-server/logger"
	"github.com/yeqown/server-common/code"
)

/*
 * 新建角色
 */
type newRoleForm struct {
	RoleName string `form:"role_name" binding:"required"`
}

type newRoleResp struct {
	code.CodeInfo
	Role *services.Role `json:"role"`
}

// NewRole ...
func NewRole(c *gin.Context) {
	var (
		form = new(newRoleForm)
		resp = new(newRoleResp)
	)

	if err := c.ShouldBind(form); err != nil {
		ResponseError(c, err)
		return
	}

	role, err := auth.CreateRole(form.RoleName)
	if err != nil {
		code.FillCodeInfo(resp, code.NewCodeInfo(code.CodeSystemErr, err.Error()))
		logger.Logger.Error(err.Error())
		Response(c, resp)
		return
	}

	code.FillCodeInfo(resp, code.NewCodeInfo(code.CodeOk, ""))
	resp.Role = role
	Response(c, resp)
	return
}

/*
 * 查询所有角色
 */
type queryAllRolesForm struct {
	Limit int `form:"limit"`
	Skip  int `form:"skip"`
}

type queryAllRolesResp struct {
	code.CodeInfo
	Roles []*services.Role `json:"roles"`
}

// QueryRole ...
func QueryRole(c *gin.Context) {
	var (
		form = new(queryAllRolesForm)
		resp = new(queryAllRolesResp)
	)

	if err := c.ShouldBind(form); err != nil {
		ResponseError(c, err)
		return
	}

	resp.Roles = nil

	roles, err := auth.GetAllRoles(form.Skip, form.Limit, CreateTimeDesc)
	if err != nil {
		code.FillCodeInfo(resp, code.NewCodeInfo(code.CodeSystemErr, err.Error()))
		Response(c, resp)
		return
	}
	fmt.Println(roles)

	code.FillCodeInfo(resp, code.NewCodeInfo(code.CodeOk, ""))
	resp.Roles = roles
	Response(c, resp)
	return
}

/*
 * 为角色增加权限
 */
type assignPerToRoleForm struct {
	RoleID       string `form:"role_id" binding:"required"`
	PermissionID string `form:"permission_id" binding:"required"`
}

type assignPerToRoleResp struct {
	code.CodeInfo
	Role *services.Role `json:"role"`
}

// AssignRolePermission ....
func AssignRolePermission(c *gin.Context) {
	var (
		form = new(assignPerToRoleForm)
		resp = new(assignPerToRoleResp)
	)

	if err := c.ShouldBind(form); err != nil {
		ResponseError(c, err)
		return
	}

	role, err := auth.Assign(form.RoleID, form.PermissionID)

	if err != nil {
		code.FillCodeInfo(resp, code.NewCodeInfo(code.CodeSystemErr, err.Error()))
		logger.Logger.Error(err.Error())
		Response(c, resp)
		return
	}

	code.FillCodeInfo(resp, code.NewCodeInfo(code.CodeOk, ""))
	resp.Role = role
	Response(c, resp)
	return
}

/*
 * 为角色删除权限
 */
type delPerToRoleForm struct {
	RoleID       string `form:"role_id" binding:"required"`
	PermissionID string `form:"permission_id" binding:"required"`
}

type delPerToRoleResp struct {
	code.CodeInfo
	Role *services.Role `json:"role"`
}

// DelRolePermission ...
func DelRolePermission(c *gin.Context) {
	var (
		form = new(delPerToRoleForm)
		resp = new(delPerToRoleResp)
	)

	if err := c.ShouldBind(form); err != nil {
		ResponseError(c, err)
		return
	}

	role, err := auth.Revoke(form.RoleID, form.PermissionID)
	if err != nil {
		code.FillCodeInfo(resp, code.NewCodeInfo(code.CodeSystemErr, err.Error()))
		logger.Logger.Error(err.Error())
		Response(c, resp)
		return
	}

	code.FillCodeInfo(resp, code.NewCodeInfo(code.CodeOk, ""))
	resp.Role = role
	Response(c, resp)
	return
}
