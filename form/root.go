package form

import (
	"github.com/graphql-go/graphql"
	"go.uber.org/fx"
)

type GraphGroup struct {
	fx.In
	Query    *graphql.Object `name:"rootQuery"`
	Mutation *graphql.Object `name:"rootMutation"`

	Queries   []*graphql.Object `group:"query"`
	Mutations []*graphql.Object `group:"mutation"`
}

type RootQueryOut struct {
	fx.Out
	Name  *graphql.Object `name:"rootQuery"`
	Group *graphql.Object `group:"query"`
}

type RootMutationOut struct {
	fx.Out
	Name  *graphql.Object `name:"rootMutation"`
	Group *graphql.Object `group:"mutation"`
}

func RootQuery() RootQueryOut {
	obj := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query", Description: "Root query", Fields: graphql.Fields{},
	})
	return RootQueryOut{Name: obj, Group: obj}
}

func RootMutation() RootMutationOut {
	obj := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation", Description: "Root mutation", Fields: graphql.Fields{},
	})
	return RootMutationOut{Name: obj, Group: obj}
}
