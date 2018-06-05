package services

import (
	. "github.com/yeqown/gweb/logger"
	. "github.com/yeqown/gweb/utils"

	"errors"
	auth "github.com/ne7ermore/gRBAC"
)

type Auth struct {
	Mobile  string `json:"mobile"`
	Ticket  string `json:"ticket"`
	ResDesc string `json:"res_desc"`
	Action  string `json:"action"`
}

type AuthArgs struct {
	Mobile  string `json:"mobile"`
	Ticket  string `json:"ticket"`
	ResDesc string `json:"res_desc"`
	Action  string `json:"action"`
}

func (a *Auth) IsPermitted(args *AuthArgs, reply *bool) (err error) {
	a.Mobile = args.Mobile
	a.Ticket = args.Ticket
	a.ResDesc = args.ResDesc
	a.Action = args.Action

	AppL.Infof("Called with: `%s %s %s %s`",
		args.Mobile,
		args.Ticket,
		args.ResDesc,
		args.Action,
	)

	// perm_id from res_desc and action
	ci, perm_id := getPermID(a.ResDesc, a.Action)
	if ci.Code != CodeOk {
		AppL.Errorf("get perm id err: %s\n", ci.Message)
		err = errors.New(ci.Message)
		return
	}

	ci, user_id := getUserID(a.Mobile)
	if ci.Code != CodeOk {
		AppL.Errorf("get user id err: %s\n", ci.Message)
		// 默认按照未找到处理，当作一般用户来鉴权

		r, _err := auth.GetRoleByName(UserRoleName)
		err = _err

		if err != nil {
			AppL.Errorf("get role id err: %s\n", ci.Message)
			return
		}

		role_id := r.Id.Hex()
		AppL.Infof("call IsRolePermitted with user id: %s, perm_id: %s", role_id, perm_id)

		*reply, err = auth.IsRolePermitted(role_id, perm_id)
		return
	}

	AppL.Infof("call IsPrmitted with user id: %s, perm_id: %s", user_id, perm_id)

	// chech permit
	*reply, err = auth.IsPrmitted(user_id, perm_id)
	if err != nil {
		AppL.Error(err.Error())
		return
	}
	AppL.Infof("IsPrmitted: %v", *reply)
	return
}

// 根据主服务来查询用户信息
func getUserID(mobile string) (*CodeInfo, string) {
	u, err := auth.GetUserByUid(mobile)
	if err != nil {
		return NewCodeInfo(CodeSystemErr, err.Error()), ""
	}
	return NewCodeInfo(CodeOk, ""), u.Id.Hex()
}

// 有可能权限未找到～
func getPermID(resDesc, action string) (*CodeInfo, string) {
	perm, err := auth.GetPermByDesc(
		Fstring("%s:%s:*", resDesc, action),
	)
	if err != nil {
		return NewCodeInfo(CodeSystemErr, err.Error()), ""
	}
	return NewCodeInfo(CodeOk, ""), perm.Id.Hex()
}
