package auth

//https://blog.csdn.net/weixin_45566022/article/details/108880608
import (
	"fmt"
	"github.com/casbin/casbin/v2"
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Casbin struct {
	enforcer *casbin.Enforcer
}

func NewCasbin(d *gorm.DB) (*Casbin, error) {
	a, err := adapter.NewAdapterByDBUseTableName(d, "sys", "casbin")
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer("../conf/rbac_model.conf", a)
	if err != nil {
		return nil, err
	}
	err = e.LoadPolicy()
	if err != nil {
		return nil, err
	}
	return &Casbin{enforcer: e}, nil
}

func (my *Casbin) Name() string {
	return "Casbin"
}

func (my *Casbin) Init(r *gin.RouterGroup) {
	r.Group("api").Use(my.permission())
}

func (my *Casbin) permission() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取资源
		obj := c.Request.URL.RequestURI()
		//获取方法
		act := c.Request.Method
		//获取实体
		sub := "admin"
		//判断策略中是否存在
		ok, err := my.enforcer.Enforce(sub, obj, act)
		if err != nil {
			return
		}
		if ok {
			fmt.Println("通过权限")
			c.Next()
		} else {
			fmt.Println("没有通过权限")
			c.Abort()
		}
	}
}
