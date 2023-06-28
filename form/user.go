package form

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/gcms/data"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type UserQueryOut struct {
	fx.Out
	Name  *graphql.Object `name:"userQuery"`
	Group *graphql.Object `group:"query"`
}

func UserQuery(root *graphql.Object, db *gorm.DB) UserQueryOut {
	userType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(*data.User).ID, nil
					},
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
	root.AddFieldConfig("users", &graphql.Field{
		Type:        graphql.NewList(userType),
		Description: "List user",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			var res []*data.User
			err := db.Model(&data.User{}).Find(&res).Error
			return res, err
		},
	})
	return UserQueryOut{Name: userType, Group: userType}
}
