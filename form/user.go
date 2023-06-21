package form

import (
	"github.com/graphql-go/graphql"
)

func UserQuery(root *graphql.Object) *graphql.Object {
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
	return userType
}
