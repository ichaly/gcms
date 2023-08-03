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
	e, err := casbin.NewEnforcer("../conf/casbin.conf", a)
	if err != nil {
		return nil, err
	}
	err = e.LoadPolicy()
	if err != nil {
		return nil, err
	}
	return e, nil
}
