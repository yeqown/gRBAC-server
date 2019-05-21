package delivery

import (
	"github.com/gin-gonic/gin"
	auth "github.com/ne7ermore/gRBAC"
	"github.com/ne7ermore/gRBAC/services"
	"github.com/yeqown/gRBAC-server/internal-modules/common"
	"github.com/yeqown/gRBAC-server/pkg/logger"
	"github.com/yeqown/infrastructure/types/codes"
)

/*
 * 新建用户
 */
type newUserForm struct {
	Mobile string `form:"mobile" valid:"Required;Mobile"`
}

type newUserResp struct {
	codes.Proto
	User *services.User `json:"user"`
}

// NewUser ...
func NewUser(c *gin.Context) {
	var (
		form = new(newUserForm)
		resp = new(newUserResp)
	)

	if err := c.ShouldBind(form); err != nil {
		common.ResponseError(c, err)
		return
	}

	user, err := auth.CreateUser(form.Mobile)
	if err != nil {
		codes.Fill(resp, codes.New(codes.CodeSystemErr, err.Error()))
		logger.Logger.Error(err.Error())
		common.Response(c, resp)
		return
	}
	codes.Fill(resp, codes.New(codes.CodeOK, ""))
	resp.User = user
	common.Response(c, resp)
	return
}

/*
 * 查询所有用户
 */
type queryAllUsersForm struct {
	Limit int `form:"limit" valid:"Min(1)"`
	Skip  int `shcema:"skip" valid:"Min(0)"`
	// Field string `shcema:"field" valid:"Required"`
}

type queryAllUsersResp struct {
	codes.Proto
	Users []*services.User `json:"users"`
}

// QueryUser ...
func QueryUser(c *gin.Context) {

	var (
		form = new(queryAllUsersForm)
		resp = new(queryAllUsersResp)
	)

	if err := c.ShouldBind(form); err != nil {
		common.ResponseError(c, err)
		return
	}

	mapUsers, err := auth.GetAllUsers(form.Skip, form.Limit, common.CreateTimeDesc)
	if err != nil {
		codes.Fill(resp, codes.New(codes.CodeSystemErr, err.Error()))
		common.Response(c, resp)
		return
	}
	// us := make([]*services.User, 0, len(mapUsers))

	// for _, val := range mapUsers {
	// 	// logger.Logger.Info(k)
	// 	// us = append(us, val.(*services.User))
	// 	println(val)
	// }

	codes.Fill(resp, codes.New(codes.CodeOK, ""))
	resp.Users = mapUsers
	common.Response(c, resp)
	return
}

/*
 * 为用户增加权限
 */
type assignPerToUserForm struct {
	UserID string `form:"user_id" valid:"Required"`
	RoleID string `form:"role_id" valid:"Required"`
}

type assignPerToUserResp struct {
	codes.Proto
	User *services.User `json:"user"`
}

// AssignUserPermission ...
func AssignUserPermission(c *gin.Context) {
	var (
		form = new(assignPerToUserForm)
		resp = new(assignPerToUserResp)
	)

	if err := c.ShouldBind(form); err != nil {
		common.ResponseError(c, err)
		return
	}

	resp.User = nil

	user, err := auth.AddRole(form.UserID, form.RoleID)
	if err != nil {
		codes.Fill(resp, codes.New(codes.CodeSystemErr, err.Error()))
		logger.Logger.Error(err.Error())
		common.Response(c, resp)
		return
	}

	codes.Fill(resp, codes.New(codes.CodeOK, ""))
	resp.User = user
	common.Response(c, resp)
	return
}

/*
 * 为用户删除权限
 */
type delPerToUserForm struct {
	UserID string `form:"user_id" valid:"Required"`
	RoleID string `form:"role_id" valid:"Required"`
}

type delPerToUserResp struct {
	codes.Proto
	User *services.User `json:"user"`
}

// DelUserPermission ...
func DelUserPermission(c *gin.Context) {
	var (
		form = new(delPerToUserForm)
		resp = new(delPerToUserResp)
	)

	if err := c.ShouldBind(form); err != nil {
		common.ResponseError(c, err)
		return
	}

	user, err := auth.DelRole(form.UserID, form.RoleID)
	if err != nil {
		codes.Fill(resp, codes.New(codes.CodeSystemErr, err.Error()))
		logger.Logger.Error(err.Error())
		common.Response(c, resp)
		return
	}

	codes.Fill(resp, codes.New(codes.CodeOK, ""))
	resp.User = user
	common.Response(c, resp)
	return
}
