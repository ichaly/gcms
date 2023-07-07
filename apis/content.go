package apis

import (
	"context"
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/boot"
	"github.com/ichaly/gcms/core"
	"github.com/ichaly/gcms/data"
	"gorm.io/gorm"
)

type ContentApi struct {
	db     *gorm.DB
	loader *boot.Loader[uint64, *data.Content]
}

func NewContentApi(d *gorm.DB, e *boot.Engine) core.Schema {
	my := &ContentApi{db: d}
	my.loader = boot.NewBatchedLoader(my.batchFunc)
	_ = e.AddTo(my.GetContests, boot.Query, "contents", "")
	_ = e.AddTo(my.GetContestsForUser, "User", "contents", "")
	return my
}

func (my *ContentApi) batchFunc(_ context.Context, keys []uint64) []*boot.Result[*data.Content] {
	var contents []*data.Content
	err := my.db.Model(&data.Content{}).Where("created_by in ?", keys).Find(&contents).Error
	values := make(map[uint64]*data.Content)
	for _, c := range contents {
		values[*c.CreatedBy] = c
	}
	results := make([]*boot.Result[*data.Content], len(keys))
	for i, k := range keys {
		r := &boot.Result[*data.Content]{
			Error: err,
		}
		if v, ok := values[k]; ok {
			r.Data = v
		}
		results[i] = r
	}
	return results
}

func (my *ContentApi) GetContests(p graphql.ResolveParams) ([]*data.Content, error) {
	return nil, nil
}

func (my *ContentApi) GetContestsForUser(p graphql.ResolveParams) (*data.Content, error) {
	uid := p.Source.(*data.User).ID
	thunk := my.loader.Load(p.Context, uint64(uid))
	return thunk()
}
