package apis

import "github.com/graphql-go/graphql"

func NewSchema() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "Schema",
		Fields:      graphql.Fields{},
		Description: "this is root for graphql schema",
	})
}
