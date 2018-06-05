package main

import (
	"flag"
	"github.com/yeqown/gweb"
	. "github.com/yeqown/gweb/logger"
	"github.com/yeqown/gweb/utils"
	"net/http"
	"net/rpc"

	ctr "auth-server/controllers"
	"auth-server/services"
)

var (
	conf     = flag.String("conf", "./configs/config.json", "-conf filename")
	init_rpu = flag.Bool("init_rpu", false, "-init_rpu true or false")
)

func main() {
	flag.Parse()
	if err := LoadConfig(*conf); err != nil {
		AppL.Fatal(err.Error())
	}

	if *init_rpu {
		services.InitRPU()
		return
	}
	// set file Handler
	gweb.SetFileHanlder("/admin/", "./view")

	registerRpcService(gweb.GetRpcServer())
	go gweb.StartRpcServer(server_ins.RpcC)

	gweb.SetEntryHook(authTokenValid)

	registerHttpHandler()
	gweb.StartHttpServer(server_ins.HttpC)

}

// 必须拥有这个token才能操作auth-server
// type HandleEntryFunc func(w http.ResponseWriter, req *http.Request, closed chan bool)
//
func authTokenValid(w http.ResponseWriter, req *http.Request) *utils.CodeInfo {
	_req_cpy := utils.CopyRequest(req)
	utils.ParseRequest(_req_cpy)

	token := _req_cpy.FormValue("token")

	if token != server_ins.Token {
		return utils.NewCodeInfo(utils.CodeParamRequired, "token非法")
	}

	// support cross origin
	w.Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
	return nil
}

func registerHttpHandler() {
	gweb.AddRoute(
		&gweb.Route{
			"/perm/new", http.MethodPost, ctr.NewPermission,
			ctr.PoolNewPermissionForm,
			ctr.PoolNewPermissionResp,
		})
	gweb.AddRoute(
		&gweb.Route{
			"/perm/edit", http.MethodPost, ctr.EditPermission,
			ctr.PoolEditPermissionForm,
			ctr.PoolEditPermissionResp,
		})
	gweb.AddRoute(
		&gweb.Route{
			"/perm/list", http.MethodGet, ctr.QueryPermission,
			ctr.PoolQueryPermissionForm,
			ctr.PoolQueryPermissionResp,
		})
	// 角色
	gweb.AddRoute(
		&gweb.Route{
			"/role/new", http.MethodPost, ctr.NewRole,
			ctr.PoolNewRoleForm,
			ctr.PoolNewRoleResp,
		})
	gweb.AddRoute(
		&gweb.Route{
			"/role/list", http.MethodGet, ctr.QueryRole,
			ctr.PoolQueryAllRolesForm,
			ctr.PoolQueryAllRolesResp,
		})
	gweb.AddRoute(
		&gweb.Route{
			"/role/assign_perm", http.MethodPost, ctr.AssignRolePermission,
			ctr.PoolAssignPerToRoleForm,
			ctr.PoolAssignPerToRoleResp,
		})
	gweb.AddRoute(
		&gweb.Route{
			"/role/revoke_perm", http.MethodPost, ctr.DelRolePermission,
			ctr.PoolDelPerToRoleForm,
			ctr.PoolDelPerToRoleResp,
		})
	// 用户
	gweb.AddRoute(
		&gweb.Route{
			"/user/new", http.MethodPost, ctr.NewUser,
			ctr.PoolNewUserForm,
			ctr.PoolNewUserResp,
		})
	gweb.AddRoute(
		&gweb.Route{
			"/user/list", http.MethodGet, ctr.QueryUser,
			ctr.PoolQueryAllUsersForm,
			ctr.PoolQueryAllUsersResp,
		})
	gweb.AddRoute(
		&gweb.Route{
			"/user/assign_role", http.MethodPost, ctr.AssignUserPermission,
			ctr.PoolAssignPerToUserForm,
			ctr.PoolAssignPerToUserResp,
		})
	gweb.AddRoute(
		&gweb.Route{
			"/user/revoke_role", http.MethodPost, ctr.DelUserPermission,
			ctr.PoolDelPerToUserForm,
			ctr.PoolDelPerToUserResp,
		})
	// 鉴权
	gweb.AddRoute(
		&gweb.Route{
			"/user/auth", http.MethodPost, ctr.Auth,
			ctr.PoolIsPermittedForm,
			ctr.PoolIsPermittedResp,
		})
}

// registerRpcService
func registerRpcService(s *rpc.Server) {
	authRpc := new(services.Auth)
	s.Register(authRpc)
}
