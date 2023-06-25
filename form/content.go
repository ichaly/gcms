package form

import (
	"context"
	"github.com/graph-gophers/dataloader/v7"
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/data"
	"github.com/ichaly/gcms/util"
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
		db.Model(&data.Content{}).Where("created_by in ?", keys).Find(&contents)
		values := make(map[uint64]*data.Content)
		util.Reduce(contents, func(values map[uint64]*data.Content, c *data.Content) map[uint64]*data.Content {
			values[*c.CreatedBy] = c
			return values
		}, values)

		results := make([]*dataloader.Result[*data.Content], len(keys))
		for i, k := range keys {
			if v, ok := values[k]; ok {
				results[i] = &dataloader.Result[*data.Content]{Data: v}
			}
		}
		return results
	}
	cache := &dataloader.NoCache[uint64, *data.Content]{}
	loader := dataloader.NewBatchedLoader(batchFunc, dataloader.WithCache[uint64, *data.Content](cache))

	contentType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Content",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(*data.Content).ID, nil
					},
				},
				"title": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(*data.Content).Title, nil
					},
				},
			},
		},
	)
	user.AddFieldConfig("content", &graphql.Field{
		Type:        contentType,
		Description: "Get content by creator id",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			uid := p.Source.(*data.User).ID
			thunk := loader.Load(p.Context, uid)
			return func() (interface{}, error) {
				return thunk()
			}, nil
		},
	})
	return ContentQueryOut{Name: contentType, Group: contentType}
}
