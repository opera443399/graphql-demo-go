package types

import (
	"github.com/graphql-go/graphql"
)

//Myself 个人信息
type Myself struct {
	Phone string `json:"basic"`
	Token string `json:"token"`
}

//GQLMyself GraphQL Myself Struct
var GQLMyself = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Myself",
		Fields: graphql.Fields{
			"phone": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "手机号码",
			},
			"token": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "token",
			},
		},
	},
)
