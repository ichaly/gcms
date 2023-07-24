package user

import (
	"context"
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/boot"
	"github.com/ichaly/gcms/core"
	"github.com/ichaly/gcms/data"
	"gorm.io/gorm"
)

type contents struct {
	db     *gorm.DB
	loader *boot.Loader[uint64, []*data.Content]
}

func NewContents(db *gorm.DB) core.Schema {
	my := &contents{db: db}
	my.loader = boot.NewBatchedLoader(my.batchContents)
	return my
}

func (*contents) Name() string {
	return "contents"
}

func (*contents) Host() interface{} {
	return User
}

func (*contents) Description() string {
	return "用户作品"
}

func (my *contents) Resolve(p graphql.ResolveParams) func() ([]*data.Content, error) {
	uid := uint64(p.Source.(*data.User).ID)
	return my.loader.Load(p.Context, uid)
}

func (my *contents) batchContents(_ context.Context, keys []uint64) []*boot.Result[[]*data.Content] {
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
