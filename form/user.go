package form

import (
	"github.com/graphql-go/graphql"
	"go.uber.org/fx"
)

type UserQueryOut struct {
	fx.Out
	Name  *graphql.Object `name:"userQuery"`
	Group *graphql.Object `group:"query"`
}

func UserQuery(root *graphql.Object) UserQueryOut {
	userType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
			},
		},
	)
	root.AddFieldConfig("user", &graphql.Field{
		Type:        userType,
		Description: "Get user by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return nil, nil
		},
	})
	return UserQueryOut{Name: userType, Group: userType}
}
