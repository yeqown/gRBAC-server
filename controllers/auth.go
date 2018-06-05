package controllers

import (
	auth "github.com/ne7ermore/gRBAC"
	. "github.com/yeqown/gweb/logger"
	"github.com/yeqown/gweb/utils"
	"sync"
)

/*
 * 鉴权
 */
type IsPermittedForm struct {
	Mobile  string `schema:"mobile" valid:"Required;Length(11)"`
	ResName string `schema:"res_name" valid:"Required;MinSize(1)"`
	Action  string `schema:"action" valid:"Required;MinSize(1)"`
}

var PoolIsPermittedForm = &sync.Pool{New: func() interface{} { return &IsPermittedForm{} }}

type IsPermittedResp struct {
	utils.CodeInfo
	Permitted bool `json:"permitted"`
}

var PoolIsPermittedResp = &sync.Pool{New: func() interface{} { return &IsPermittedResp{} }}

func Auth(req *IsPermittedForm) *IsPermittedResp {
	res := PoolIsPermittedResp.Get().(*IsPermittedResp)
	defer PoolNewPermissionResp.Put(res)
	res.Permitted = false

	// get user by UserID
	u, err := auth.GetUserByUid(req.Mobile)
	if err != nil {
		utils.Response(res, utils.NewCodeInfo(utils.CodeSystemErr, err.Error()))
		AppL.Errorf("get user with err: %s\n", err.Error())
		return res
	}
	AppL.Infof("get user with mongoid: %s\n", u.Id.Hex())

	// get perm by "res:action:*"
	p, err := auth.GetPermByDesc(
		utils.Fstring("%s:%s:%s", req.ResName, req.Action, "*"),
	)
	if err != nil {
		utils.Response(res, utils.NewCodeInfo(utils.CodeSystemErr, err.Error()))
		AppL.Errorf("get perm with err: %s\n", err.Error())
		return res
	}
	AppL.Infof("get permission with mongoid: %s\n", p.Id.Hex())

	permitted, err := auth.IsPrmitted(u.Id.Hex(), p.Id.Hex())
	if err != nil {
		utils.Response(res, utils.NewCodeInfo(utils.CodeSystemErr, err.Error()))
		AppL.Errorf(err.Error())
		return res
	}

	utils.Response(res, utils.NewCodeInfo(utils.CodeOk, ""))
	res.Permitted = permitted
	return res
}
