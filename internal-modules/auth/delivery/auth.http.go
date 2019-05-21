package delivery

import (
	"fmt"

	"github.com/gin-gonic/gin"
	auth "github.com/ne7ermore/gRBAC"

	"github.com/yeqown/gRBAC-server/internal-modules/common"
	"github.com/yeqown/gRBAC-server/pkg/logger"
	"github.com/yeqown/gRBAC-server/pkg/secret"
	"github.com/yeqown/infrastructure/types/codes"
)

/*
 * 鉴权
 */
type isPermittedForm struct {
	UID     string `form:"uid" binding:"required"`
	ResName string `form:"res_name" binding:"required"`
	Action  string `form:"action" binding:"required"`
}

type isPermittedResp struct {
	codes.Proto
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
		common.ResponseError(c, err)
		return
	}

	// get user by UserID
	u, err := auth.GetUserByUid(form.UID)
	if err != nil {
		codes.Fill(resp, codes.New(codes.CodeSystemErr, err.Error()))
		logger.Logger.Errorf("get user with err: %s\n", err.Error())
		common.Response(c, resp)
		return
	}
	logger.Logger.Infof("get user with mongoid: %s\n", u.Id.Hex())

	// get perm by "res:action:*"
	p, err := auth.GetPermByDesc(
		fmt.Sprintf("%s:%s:%s", form.ResName, form.Action, "*"),
	)
	if err != nil {
		codes.Fill(resp, codes.New(codes.CodeSystemErr, err.Error()))
		logger.Logger.Errorf("get perm with err: %s\n", err.Error())
		common.Response(c, resp)
		return
	}
	logger.Logger.Infof("get permission with mongoid: %s\n", p.Id.Hex())

	permitted, err := auth.IsPrmitted(u.Id.Hex(), p.Id.Hex())
	if err != nil {
		codes.Fill(resp, codes.New(codes.CodeSystemErr, err.Error()))
		logger.Logger.Errorf(err.Error())
		common.Response(c, resp)
		return
	}

	codes.Fill(resp, codes.New(codes.CodeOK, ""))
	resp.Permitted = permitted
	common.Response(c, resp)
	return
}

type verifyForm struct {
	Secret string `form:"secret" binding:"required"`
}
type verifyResp struct {
	codes.Proto
	Verified bool   `json:"verified"`
	Token    string `json:"token,omitempty"`
}

// Verify ...
func Verify(c *gin.Context) {
	var (
		form = new(verifyForm)
		resp = new(verifyResp)
	)
	// resp.Verified = false
	if err := c.ShouldBind(form); err != nil {
		common.ResponseError(c, err)
		return
	}

	logger.Logger.Infof("get %s, want %s", form.Secret, secret.GetSecret())

	if form.Secret != secret.GetSecret() {
		codes.Fill(resp, codes.New(codes.CodeSystemErr, "wrong token input"))
		common.Response(c, resp)
		return
	}

	resp.Verified = true
	resp.Token = secret.GetSecret()
	codes.Fill(resp, codes.Get(codes.CodeOK))
	common.Response(c, resp)
	return
}
