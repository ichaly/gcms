package form

import (
	"context"
	"fmt"
	"github.com/graph-gophers/dataloader/v7"
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/data"
	"github.com/ichaly/gcms/util"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"sync"
)

type ContentQueryOut struct {
	fx.Out
	Name  *graphql.Object `name:"contentQuery"`
	Group *graphql.Object `group:"query"`
}

func ContentQuery(user *graphql.Object, db *gorm.DB) ContentQueryOut {
	batchFunc := func(_ context.Context, keys []uint64) []*dataloader.Result[*data.Content] {
		var contents []*data.Content
		db.Model(&data.Content{}).Where("id in ?", keys).Find(&contents)
		values := make(map[uint64]*data.Content)
		util.Reduce(contents, func(values map[uint64]*data.Content, c *data.Content) map[uint64]*data.Content {
			values[c.ID] = c
			return values
		}, values)

		var results []*dataloader.Result[*data.Content]
		//results := make([]*dataloader.Result[*data.Content], len(keys))
		for _, k := range keys {
			if v, ok := values[k]; ok {
				results = append(results, &dataloader.Result[*data.Content]{Data: v})
				//results[i] = &dataloader.Result[*data.Content]{Data: v}
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
		Description: "Get content by user id",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var err error
			var res *data.Content
			var wg sync.WaitGroup
			go func() {
				defer func() {
					if err := recover(); err != nil {
						fmt.Println(err)
					}
				}()
				wg.Add(1)
				defer wg.Done()
				uid := p.Source.(*data.User).ID
				fn := loader.Load(p.Context, uid)
				res, err = fn()
			}()
			wg.Wait()
			return res, err
		},
	})
	return ContentQueryOut{Name: contentType, Group: contentType}
}
