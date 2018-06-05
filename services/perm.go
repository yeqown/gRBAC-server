// 约定权限描述
// 生成内置的角色及权限

package services

import (
	auth "github.com/ne7ermore/gRBAC"
	"github.com/ne7ermore/gRBAC/services"
)

// Action admin
const (
	AdminLogin     = "登录管理后台" // 1 登录管理后台
	AdminDoRepay   = "回款"     // 4 回款
	ProductPublish = "发布产品"   // 2 发布产品
	ProductBuy     = "购买产品"   // 5 购买产品
	ProductUpdate  = "更新产品"   // 3 更新产品
	UserRecharge   = "充值"     // 6 充值
	UserWithdraw   = "提现"     // 7 提现

	AdminRoleName = "管理员"
	UserRoleName  = "普通用户"
)

var ActionsMap = map[string]string{
	AdminLogin:     "admin:login:*",     // 登录管理后台
	AdminDoRepay:   "admin:repay:*",     // 管理后台回款
	ProductPublish: "product:publish:*", // 新增产品
	ProductBuy:     "product:buy:*",     // 购买产品
	ProductUpdate:  "product:update:*",  // 更新产品
	UserRecharge:   "user:recharge:*",   // 充值
	UserWithdraw:   "user:withdraw:*",   // 提现
}

var (
	innerPerms map[string]*services.Permission
	innerRoles map[string]*services.Role
	adminPerms []string
	userPerms  []string
)

// 生成内置角色，管理员，一般用户
func initInnerRole() {
	if r, err := auth.CreateRole(AdminRoleName); err != nil {
		println(err.Error())
	} else {
		innerRoles[AdminRoleName] = r
		for _, perm := range adminPerms {
			if _, err := auth.Assign(r.Id.Hex(), innerPerms[perm].Id.Hex()); err != nil {
				println(err.Error())
				continue
			}
			println("assgin to role `管理员` with perm: ", perm)
		}
	}

	if r, err := auth.CreateRole(UserRoleName); err != nil {
		println(err.Error())
	} else {
		innerRoles[UserRoleName] = r
		for _, perm := range userPerms {
			if _, err := auth.Assign(r.Id.Hex(), innerPerms[perm].Id.Hex()); err != nil {
				println(err.Error())
				continue
			}
			println("assgin to role `普通用户` with perm: ", perm)
		}
	}
}

// 生成内置权限，参见ActionsMap
func initInnerPerm() {
	for k, val := range ActionsMap {
		perm, err := auth.CreatePermisson(k, val)
		if err != nil {
			println(err)
			continue
		}
		innerPerms[k] = perm
	}
}

// 生成内置用户，18302889215
func initInnerUser() {
	u, err := auth.CreateUser("18302889215")
	if err != nil {
		println(err.Error())
		return
	}

	if _, err := auth.AddRole(u.Id.Hex(), innerRoles[AdminRoleName].Id.Hex()); err != nil {
		println(err.Error())
	}
}

// 生成内置角色权限
func InitRPU() {
	innerPerms = make(map[string]*services.Permission, len(ActionsMap))
	innerRoles = make(map[string]*services.Role, 2)

	adminPerms = []string{
		AdminLogin,
		AdminDoRepay,
		ProductPublish,
		ProductBuy,
		ProductUpdate,
		UserRecharge,
		UserWithdraw,
	}

	userPerms = []string{
		ProductBuy,
		UserRecharge,
		UserWithdraw,
	}

	initInnerPerm()
	initInnerRole()
	initInnerUser()
}
