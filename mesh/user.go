package mesh

import (
	"context"
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/boot"
	"github.com/ichaly/gcms/core"
	"github.com/ichaly/gcms/data"
	"gorm.io/gorm"
	"time"
)

const User = "User"

type UserSchema struct {
	db     *gorm.DB
	loader *boot.Loader[uint64, []*data.Content]
}

func NewUserSchema(d *gorm.DB, e *boot.Engine) core.Schema {
	my := &UserSchema{db: d}
	my.loader = boot.NewBatchedLoader(my.batchFunc)
	e.NewQuery(my.GetUsers)
	e.NewBuilder(my.GetAge).To(User).Description("年龄")
	e.NewBuilder(my.GetContents).To(User).Description("用户内容")
	return my
}

func (my *UserSchema) GetUsers(p graphql.ResolveParams) ([]*data.User, error) {
	var res []*data.User
	err := my.db.Model(&data.User{}).Find(&res).Error
	return res, err
}

func (my *UserSchema) GetAge(p graphql.ResolveParams) (int, error) {
	user := p.Source.(*data.User)
	if user.Birthday.IsZero() {
		return 0, nil
	}
	return time.Now().Year() - user.Birthday.Year(), nil
}

func (my *UserSchema) GetContents(p graphql.ResolveParams) func() ([]*data.Content, error) {
	uid := uint64(p.Source.(*data.User).ID)
	return my.loader.Load(p.Context, uid)
}

func (my *UserSchema) batchFunc(_ context.Context, keys []uint64) []*boot.Result[[]*data.Content] {
	var res []*data.Content
	err := my.db.Model(&data.Content{}).Where("created_by in ?", keys).Find(&res).Error
	values := make(map[uint64][]*data.Content)
	for _, c := range res {
		values[*c.CreatedBy] = append(values[*c.CreatedBy], c)
	}
	results := make([]*boot.Result[[]*data.Content], len(keys))
	for i, k := range keys {
		r := &boot.Result[[]*data.Content]{
			Error: err,
		}
		if v, ok := values[k]; ok {
			r.Data = v
		}
		results[i] = r
	}
	return results
}
