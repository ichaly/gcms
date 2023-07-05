package form

import (
	"context"
	"github.com/graph-gophers/dataloader/v7"
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/data"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type ContentQueryOut struct {
	fx.Out
	Name  *graphql.Object `name:"contentQuery"`
	Group *graphql.Object `group:"query"`
}

func ContentQuery(user *graphql.Object, db *gorm.DB) ContentQueryOut {
	batchFunc := func(_ context.Context, keys []uint64) []*dataloader.Result[*data.Content] {
		var contents []*data.Content
		err := db.Model(&data.Content{}).Where("created_by in ?", keys).Find(&contents).Error
		values := make(map[uint64]*data.Content)
		for _, c := range contents {
			values[*c.CreatedBy] = c
		}
		results := make([]*dataloader.Result[*data.Content], len(keys))
		for i, k := range keys {
			r := &dataloader.Result[*data.Content]{
				Error: err,
			}
			if v, ok := values[k]; ok {
				r.Data = v
			}
			results[i] = r
		}
		return results
	}
	cache := &dataloader.NoCache[uint64, *data.Content]{}
	loader := dataloader.NewBatchedLoader(batchFunc, dataloader.WithCache[uint64, *data.Content](cache))

	contentType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Content",
			Fields: graphql.Fields{
				"id":    &graphql.Field{Type: graphql.ID},
				"title": &graphql.Field{Type: graphql.String},
			},
		},
	)
	user.AddFieldConfig("content", &graphql.Field{
		Type:        contentType,
		Description: "Get content by creator id",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			uid := p.Source.(*data.User).ID
			thunk := loader.Load(p.Context, uint64(uid))
			return func() (interface{}, error) {
				return thunk()
			}, nil
		},
	})
	return ContentQueryOut{Name: contentType, Group: contentType}
}
