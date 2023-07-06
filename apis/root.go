package apis

import "github.com/graphql-go/graphql"

func NewRoot(schema *graphql.Object) *graphql.Object {
	root := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Root",
		Fields:      graphql.Fields{},
		Description: "this is root for graphql schema",
	})
	schema.AddFieldConfig("root", &graphql.Field{
		Type: root, Description: "Root Query",
	})
	return root
}
