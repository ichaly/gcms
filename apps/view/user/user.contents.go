package user

import (
	"context"
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/apps/data"
	"github.com/ichaly/gcms/core"
	"gorm.io/gorm"
)

type contents struct {
	db     *gorm.DB
	loader *core.Loader[uint64, []*data.Content]
}

func NewUserContents(db *gorm.DB) core.Schema {
	my := &contents{db: db}
	my.loader = core.NewBatchedLoader(my.batchContents)
	return my
}

func (*contents) Name() string {
	return "contents"
}

func (*contents) Description() string {
	return "用户作品"
}

func (*contents) Host() interface{} {
	return User
}

func (*contents) Type() interface{} {
	return []*data.Content{}
}

func (my *contents) Resolve(p graphql.ResolveParams) (interface{}, error) {
	uid := uint64(p.Source.(*data.User).ID)
	thunk := my.loader.Load(p.Context, uid)
	return func() (interface{}, error) {
		return thunk()
	}, nil
}

func (my *contents) batchContents(ctx context.Context, keys []uint64) []*core.Result[[]*data.Content] {
	var res []*data.Content
	err := my.db.WithContext(ctx).Model(&data.Content{}).Where("created_by in ?", keys).Find(&res).Error
	values := make(map[uint64][]*data.Content)
	for _, c := range res {
		values[*c.CreatedBy] = append(values[*c.CreatedBy], c)
	}
	results := make([]*core.Result[[]*data.Content], len(keys))
	for i, k := range keys {
		r := &core.Result[[]*data.Content]{
			Error: err,
		}
		if v, ok := values[k]; ok {
			r.Data = v
		}
		results[i] = r
	}
	return results
}
