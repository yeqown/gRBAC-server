package controllers

import (
	"fmt"
	auth "github.com/ne7ermore/gRBAC"
	"github.com/ne7ermore/gRBAC/services"
	. "github.com/yeqown/gweb/logger"
	"github.com/yeqown/gweb/utils"
	"sync"
)

/*
 * 新建角色
 */
type NewRoleForm struct {
	RoleName string `schema:"role_name" valid:"Required"`
}

var PoolNewRoleForm = &sync.Pool{New: func() interface{} { return &NewRoleForm{} }}

type NewRoleResp struct {
	utils.CodeInfo
	Role *services.Role `json:"role"`
}

var PoolNewRoleResp = &sync.Pool{New: func() interface{} { return &NewRoleResp{} }}

func NewRole(req *NewRoleForm) *NewRoleResp {
	res := PoolNewRoleResp.Get().(*NewRoleResp)
	defer PoolNewRoleResp.Put(res)

	role, err := auth.CreateRole(req.RoleName)
	if err != nil {
		utils.Response(res, utils.NewCodeInfo(utils.CodeSystemErr, err.Error()))
		AppL.Error(err.Error())
		return res
	}

	utils.Response(res, utils.NewCodeInfo(utils.CodeOk, ""))
	res.Role = role
	return res
}

/*
 * 查询所有角色
 */
type QueryAllRolesForm struct {
	Limit int `schema:"limit" valid:"Required;Min(1)"`
	Skip  int `shcema:"skip" valid:"Min(0)"`
	// Field string `shcema:"field" valid:"Required"`
}

var PoolQueryAllRolesForm = &sync.Pool{New: func() interface{} { return &QueryAllRolesForm{} }}

type QueryAllRolesResp struct {
	utils.CodeInfo
	Roles []*services.Role `json:"roles"`
}

var PoolQueryAllRolesResp = &sync.Pool{New: func() interface{} { return &QueryAllRolesResp{} }}

func QueryRole(req *QueryAllRolesForm) *QueryAllRolesResp {
	res := PoolQueryAllRolesResp.Get().(*QueryAllRolesResp)
	defer PoolQueryAllRolesResp.Put(res)
	res.Roles = nil

	roles, err := auth.GetAllRoles(req.Skip, req.Limit, CreateTimeDesc)
	if err != nil {
		utils.Response(res, utils.NewCodeInfo(utils.CodeSystemErr, err.Error()))
		return res
	}
	fmt.Println(roles)

	utils.Response(res, utils.NewCodeInfo(utils.CodeOk, ""))
	res.Roles = roles
	return res
}

/*
 * 为角色增加权限
 */
type AssignPerToRoleForm struct {
	RoleID       string `schema:"role_id" valid:"Required"`
	PermissionID string `schema:"permission_id" valid:"Required"`
}

var PoolAssignPerToRoleForm = &sync.Pool{New: func() interface{} { return &AssignPerToRoleForm{} }}

type AssignPerToRoleResp struct {
	utils.CodeInfo
	Role *services.Role `json:"role"`
}

var PoolAssignPerToRoleResp = &sync.Pool{New: func() interface{} { return &AssignPerToRoleResp{} }}

func AssignRolePermission(req *AssignPerToRoleForm) *AssignPerToRoleResp {
	res := PoolAssignPerToRoleResp.Get().(*AssignPerToRoleResp)
	defer PoolAssignPerToRoleResp.Put(res)

	role, err := auth.Assign(req.RoleID, req.PermissionID)

	if err != nil {
		utils.Response(res, utils.NewCodeInfo(utils.CodeSystemErr, err.Error()))
		AppL.Error(err.Error())
		return res
	}

	utils.Response(res, utils.NewCodeInfo(utils.CodeOk, ""))
	res.Role = role
	return res
}

/*
 * 为角色删除权限
 */
type DelPerToRoleForm struct {
	RoleID       string `schema:"role_id" valid:"Required"`
	PermissionID string `schema:"permission_id" valid:"Required"`
}

var PoolDelPerToRoleForm = &sync.Pool{New: func() interface{} { return &DelPerToRoleForm{} }}

type DelPerToRoleResp struct {
	utils.CodeInfo
	Role *services.Role `json:"role"`
}

var PoolDelPerToRoleResp = &sync.Pool{New: func() interface{} { return &DelPerToRoleResp{} }}

func DelRolePermission(req *DelPerToRoleForm) *DelPerToRoleResp {
	res := PoolDelPerToRoleResp.Get().(*DelPerToRoleResp)
	defer PoolDelPerToRoleResp.Put(res)

	role, err := auth.Revoke(req.RoleID, req.PermissionID)
	if err != nil {
		utils.Response(res, utils.NewCodeInfo(utils.CodeSystemErr, err.Error()))
		AppL.Error(err.Error())
		return res
	}

	utils.Response(res, utils.NewCodeInfo(utils.CodeOk, ""))
	res.Role = role
	return res

}
