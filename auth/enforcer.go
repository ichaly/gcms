package auth

import (
	"github.com/casbin/casbin/v2"
	adapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func NewEnforcer(d *gorm.DB) (*casbin.Enforcer, error) {
	a, err := adapter.NewAdapterByDBUseTableName(d, "sys", "casbin")
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer("./conf/casbin.conf", a)
	if err != nil {
		return nil, err
	}
	err = e.LoadPolicy()
	if err != nil {
		return nil, err
	}
	registerFunction(e, d)
	return e, nil
}

func registerFunction(e *casbin.Enforcer, d *gorm.DB) {
	e.AddFunction("permit", func(args ...interface{}) (interface{}, error) {
		sub, obj, act := args[0].(string), args[1].(string), args[2].(string)
		//判断时候有相应的策略，如果没有则放行
		policy := e.GetFilteredPolicy(1, obj, act)
		if len(policy) == 0 {
			return true, nil
		}
		//是否是超级管理员
		return e.HasRoleForUser(sub, "superadmin")
	})
}
