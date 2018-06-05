package controllers

import (
	auth "github.com/ne7ermore/gRBAC"
	"github.com/ne7ermore/gRBAC/services"
	. "github.com/yeqown/gweb/logger"
	"github.com/yeqown/gweb/utils"
	"sync"
)

/*
 * 新建用户
 */
type NewUserForm struct {
	Mobile string `schema:"mobile" valid:"Required;Mobile"`
}

var PoolNewUserForm = &sync.Pool{New: func() interface{} { return &NewUserForm{} }}

type NewUserResp struct {
	utils.CodeInfo
	User *services.User `json:"user"`
}

var PoolNewUserResp = &sync.Pool{New: func() interface{} { return &NewUserResp{} }}

func NewUser(req *NewUserForm) *NewUserResp {
	res := PoolNewUserResp.Get().(*NewUserResp)
	defer PoolNewUserResp.Put(res)
	AppL.Info("CreateUser before")
	user, err := auth.CreateUser(req.Mobile)
	AppL.Info("CreateUser after")
	if err != nil {
		utils.Response(res, utils.NewCodeInfo(utils.CodeSystemErr, err.Error()))
		AppL.Error(err.Error())
		return res
	}
	utils.Response(res, utils.NewCodeInfo(utils.CodeOk, ""))
	res.User = user
	return res
}

/*
 * 查询所有用户
 */
type QueryAllUsersForm struct {
	Limit int `schema:"limit" valid:"Min(1)"`
	Skip  int `shcema:"skip" valid:"Min(0)"`
	// Field string `shcema:"field" valid:"Required"`
}

var PoolQueryAllUsersForm = &sync.Pool{New: func() interface{} { return &QueryAllUsersForm{} }}

type QueryAllUsersResp struct {
	utils.CodeInfo
	Users []*services.User `json:"users"`
}

var PoolQueryAllUsersResp = &sync.Pool{New: func() interface{} { return &QueryAllUsersResp{} }}

func QueryUser(req *QueryAllUsersForm) *QueryAllUsersResp {
	res := PoolQueryAllUsersResp.Get().(*QueryAllUsersResp)
	defer PoolQueryAllUsersResp.Put(res)

	map_us, err := auth.GetAllUsers(req.Skip, req.Limit, CreateTimeDesc)
	if err != nil {
		utils.Response(res, utils.NewCodeInfo(utils.CodeSystemErr, err.Error()))
		return res
	}
	// us := make([]*services.User, 0, len(map_us))

	for _, val := range map_us {
		// AppL.Info(k)
		// us = append(us, val.(*services.User))
		println(val)

	}

	utils.Response(res, utils.NewCodeInfo(utils.CodeOk, ""))
	res.Users = map_us
	return res
}

/*
 * 为用户增加权限
 */
type AssignPerToUserForm struct {
	UserID string `schema:"user_id" valid:"Required"`
	RoleID string `schema:"role_id" valid:"Required"`
}

var PoolAssignPerToUserForm = &sync.Pool{New: func() interface{} { return &AssignPerToUserForm{} }}

type AssignPerToUserResp struct {
	utils.CodeInfo
	User *services.User `json:"user"`
}

var PoolAssignPerToUserResp = &sync.Pool{New: func() interface{} { return &AssignPerToUserResp{} }}

func AssignUserPermission(req *AssignPerToUserForm) *AssignPerToUserResp {
	res := PoolAssignPerToUserResp.Get().(*AssignPerToUserResp)
	defer PoolAssignPerToUserResp.Put(res)
	res.User = nil

	user, err := auth.AddRole(req.UserID, req.RoleID)
	if err != nil {
		utils.Response(res, utils.NewCodeInfo(utils.CodeSystemErr, err.Error()))
		AppL.Error(err.Error())
		return res
	}

	utils.Response(res, utils.NewCodeInfo(utils.CodeOk, ""))
	res.User = user
	return res
}

/*
 * 为用户删除权限
 */
type DelPerToUserForm struct {
	UserID string `schema:"user_id" valid:"Required"`
	RoleID string `schema:"role_id" valid:"Required"`
}

var PoolDelPerToUserForm = &sync.Pool{New: func() interface{} { return &DelPerToUserForm{} }}

type DelPerToUserResp struct {
	utils.CodeInfo
	User *services.User `json:"user"`
}

var PoolDelPerToUserResp = &sync.Pool{New: func() interface{} { return &DelPerToUserResp{} }}

func DelUserPermission(req *DelPerToUserForm) *DelPerToUserResp {
	res := PoolDelPerToUserResp.Get().(*DelPerToUserResp)
	defer PoolDelPerToUserResp.Put(res)

	user, err := auth.DelRole(req.UserID, req.RoleID)
	if err != nil {
		utils.Response(res, utils.NewCodeInfo(utils.CodeSystemErr, err.Error()))
		AppL.Error(err.Error())
		return res
	}

	utils.Response(res, utils.NewCodeInfo(utils.CodeOk, ""))
	res.User = user
	return res

}
