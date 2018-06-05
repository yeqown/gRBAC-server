package controllers

import (
	auth "github.com/ne7ermore/gRBAC"
	"github.com/ne7ermore/gRBAC/services"
	. "github.com/yeqown/gweb/logger"
	"github.com/yeqown/gweb/utils"
	"sync"
)

/*
 * 增加权限
 */
type NewPermissionForm struct {
	PermissionDes  string `schema:"permission_desc" valid:"Required"`
	PermissionName string `schema:"permission_name" valid:"Required"`
}

var PoolNewPermissionForm = &sync.Pool{New: func() interface{} { return &NewPermissionForm{} }}

type NewPermissionResp struct {
	utils.CodeInfo
	Permission *services.Permission `json:"Permission"`
}

var PoolNewPermissionResp = &sync.Pool{New: func() interface{} { return &NewPermissionResp{} }}

func NewPermission(req *NewPermissionForm) *NewPermissionResp {
	res := PoolNewPermissionResp.Get().(*NewPermissionResp)
	PoolNewPermissionResp.Put(res)

	permission, err := auth.CreatePermisson(req.PermissionName, req.PermissionDes)
	if err != nil {
		utils.Response(res, utils.NewCodeInfo(utils.CodeSystemErr, err.Error()))
		AppL.Error(err.Error())
		return res
	}

	utils.Response(res, utils.NewCodeInfo(utils.CodeOk, ""))
	res.Permission = permission
	return res
}

/*
 * 获取权限
 */

type QueryPermissionForm struct {
	Limit int `schema:"limit" valid:"Required"`
	Skip  int `schema:"skip" valid:"Min(0)"`
	// Field string `schema:"field" valid:"Required"`
}

var PoolQueryPermissionForm = &sync.Pool{New: func() interface{} { return &QueryPermissionForm{} }}

type QueryPermissionResp struct {
	utils.CodeInfo
	Permissions []*services.Permission `json:"permissions"`
}

var PoolQueryPermissionResp = &sync.Pool{New: func() interface{} { return &QueryPermissionResp{} }}

func QueryPermission(req *QueryPermissionForm) *QueryPermissionResp {
	res := PoolQueryPermissionResp.Get().(*QueryPermissionResp)
	PoolQueryPermissionResp.Put(res)

	map_perms, err := auth.GetAllPerms(req.Skip, req.Limit, CreateTimeDesc)
	if err != nil {
		utils.Response(res, utils.NewCodeInfo(utils.CodeSystemErr, err.Error()))
		return res
	}

	for _, perm := range map_perms {
		println(perm)
		// AppL.Info(k, )
		// AppL.Info(perm.(*services.Permission))
	}

	utils.Response(res, utils.NewCodeInfo(utils.CodeOk, ""))
	res.Permissions = map_perms
	return res
}

/*
 * 编辑权限
 */

type EditPermissionForm struct {
	PermissionID string `schema:"permission_id" valid:"Required"`
	Desc         string `schema:"desc" valid:"Required;MinSize(1)"`
	Name         string `schema:"name" valid:"Required;MinSize(1)"`
}

var PoolEditPermissionForm = &sync.Pool{New: func() interface{} { return &EditPermissionForm{} }}

type EditPermissionResp struct {
	utils.CodeInfo
	Permission *services.Permission `permission`
}

var PoolEditPermissionResp = &sync.Pool{New: func() interface{} { return &EditPermissionResp{} }}

func EditPermission(req *EditPermissionForm) *EditPermissionResp {
	res := PoolEditPermissionResp.Get().(*EditPermissionResp)
	PoolEditPermissionResp.Put(res)

	permission, err := auth.UpdatePerm(req.PermissionID, req.Desc, req.Name)
	if err != nil {
		utils.Response(res, utils.NewCodeInfo(utils.CodeSystemErr, err.Error()))
		AppL.Error(err.Error())
		return res
	}

	utils.Response(res, utils.NewCodeInfo(utils.CodeOk, ""))
	res.Permission = permission
	return res
}
