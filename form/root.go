package form

import "github.com/graphql-go/graphql"

func RootQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{Name: "RootQuery", Description: "RootQuery", Fields: graphql.Fields{}})
}

func RootMutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{Name: "RootMutation", Description: "RootMutation", Fields: graphql.Fields{}})
}
