package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	auth "github.com/ne7ermore/gRBAC"

	"github.com/yeqown/gRBAC-server/logger"
	"github.com/yeqown/server-common/code"
)

/*
 * 鉴权
 */
type isPermittedForm struct {
	UID     string `form:"uid"`
	ResName string `form:"res_name"`
	Action  string `form:"action"`
}

type isPermittedResp struct {
	code.CodeInfo
	Permitted bool `json:"permitted"`
}

// Auth handler
func Auth(c *gin.Context) {
	var (
		form = new(isPermittedForm)
		resp = new(isPermittedResp)
	)

	resp.Permitted = false
	if err := c.ShouldBind(form); err != nil {
		ResponseError(c, err)
		return
	}

	// get user by UserID
	u, err := auth.GetUserByUid(form.UID)
	if err != nil {
		code.FillCodeInfo(resp, code.NewCodeInfo(code.CodeSystemErr, err.Error()))
		logger.Logger.Errorf("get user with err: %s\n", err.Error())
		Response(c, resp)
		return
	}
	logger.Logger.Infof("get user with mongoid: %s\n", u.Id.Hex())

	// get perm by "res:action:*"
	p, err := auth.GetPermByDesc(
		fmt.Sprintf("%s:%s:%s", form.ResName, form.Action, "*"),
	)
	if err != nil {
		code.FillCodeInfo(resp, code.NewCodeInfo(code.CodeSystemErr, err.Error()))
		logger.Logger.Errorf("get perm with err: %s\n", err.Error())
		Response(c, resp)
		return
	}
	logger.Logger.Infof("get permission with mongoid: %s\n", p.Id.Hex())

	permitted, err := auth.IsPrmitted(u.Id.Hex(), p.Id.Hex())
	if err != nil {
		code.FillCodeInfo(resp, code.NewCodeInfo(code.CodeSystemErr, err.Error()))
		logger.Logger.Errorf(err.Error())
		Response(c, resp)
		return
	}

	code.FillCodeInfo(resp, code.NewCodeInfo(code.CodeOk, ""))
	resp.Permitted = permitted
	Response(c, resp)
	return
}
