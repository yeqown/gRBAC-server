package usecase

import (
	"fmt"

	auth "github.com/ne7ermore/gRBAC"
	"github.com/yeqown/gRBAC-server/pkg/logger"
	"github.com/yeqown/infrastructure/types/codes"
)

// Auth for rpc calling
type Auth struct{}

// Args ...
type Args struct {
	UID string `json:"uid"`
	// Ticket  string `json:"ticket"`
	ResDesc string `json:"res_desc"`
	Action  string `json:"action"`
}

// IsPermitted check the permitted or not
func (a Auth) IsPermitted(args *Args, reply *bool) error {
	logger.Logger.Infof("Called with: `%s %s %s`",
		args.UID,
		// args.Ticket,
		args.ResDesc,
		args.Action,
	)

	// permID from res_desc and action
	ci, permID := getPermID(args.ResDesc, args.Action)
	if ci.ErrCode != codes.CodeOK {
		logger.Logger.Errorf("get perm id err: %s\n", ci.ErrMessage)
		return fmt.Errorf("could not get permission id, %v", ci.ErrMessage)
	}

	ci, userID := getUserID(args.UID)
	if ci.ErrCode != codes.CodeOK {
		logger.Logger.Errorf("get user id err: %s\n", ci.ErrMessage)
		r, err := auth.GetRoleByName("default")
		if err != nil {
			logger.Logger.Errorf("get role id err: %s\n", ci.ErrMessage)
			return err
		}

		roleID := r.Id.Hex()
		logger.Logger.Infof("call IsRolePermitted with user id: %s, permID: %s", roleID, permID)

		if *reply, err = auth.IsRolePermitted(roleID, permID); err != nil {
			return err
		}
		return nil
	}

	logger.Logger.Infof("call IsPrmitted with user id: %s, permID: %s", userID, permID)
	if b, err := auth.IsPrmitted(userID, permID); err != nil {
		logger.Logger.Error(err.Error())
		return err
	} else {
		*reply = b
	}
	logger.Logger.Infof("IsPrmitted: %v", *reply)
	return nil
}

// get user id by param
func getUserID(uid string) (*codes.Proto, string) {
	u, err := auth.GetUserByUid(uid)
	if err != nil {
		return codes.New(codes.CodeSystemErr, err.Error()), ""
	}
	return codes.New(codes.CodeOK, ""), u.Id.Hex()
}

// 有可能权限未找到～
func getPermID(resDesc, action string) (*codes.Proto, string) {
	perm, err := auth.GetPermByDesc(
		fmt.Sprintf("%s:%s:*", resDesc, action),
	)
	if err != nil {
		return codes.New(codes.CodeSystemErr, err.Error()), ""
	}
	return codes.New(codes.CodeOK, ""), perm.Id.Hex()
}
