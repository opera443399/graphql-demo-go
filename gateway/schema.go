package main

import (
	"github.com/graphql-go/graphql"
	"github.com/opera443399/cicd-backend/gateway/types"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootQuery",
	Description: "查询信息，类似GET",
	Fields: graphql.Fields{
		"myself": &graphql.Field{
			Type:        types.GQLMyself,
			Description: "查询用户信息",
			Resolve:     myself,
		},
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootMutation",
	Description: "执行操作，类似POST",
	Fields: graphql.Fields{
		"login": &graphql.Field{
			Type:        types.GQLMyself,
			Description: "测试 login 操作",
			Args: graphql.FieldConfigArgument{
				"phone": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "手机号码",
				},
				"code": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "口令",
				},
			},
			Resolve: login,
		},
	},
})

//Schema graphql schema 入口
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	},
)
