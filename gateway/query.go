package main

import (
	"github.com/graphql-go/graphql"
	"github.com/opera443399/cicd-backend/gateway/types"
)

func myself(p graphql.ResolveParams) (interface{}, error) {
	var user = types.Myself{}
	user.Phone = "我的号码是111"
	user.Token = "我的口令是zzz"
	return user, nil
}
