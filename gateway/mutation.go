package main

import (
	"fmt"
	"strings"
	
	"github.com/graphql-go/graphql"
	"github.com/opera443399/cicd-backend/gateway/types"
)

func checkPhone(p graphql.ResolveParams, phone string) bool {
	phoneInWhiteList := "15011112222|16011112222"
	if phoneInWhiteList != "" {
		phoneList := strings.Split(phoneInWhiteList, "|")
		for _, v := range phoneList {
			if phone == v {
				return true
			}
		}
		return false
	}
	return true
}

func checkCode(p graphql.ResolveParams, code string) bool {
	codeInWhiteList := "150150|160160"
	if codeInWhiteList != "" {
		codeList := strings.Split(codeInWhiteList, "|")
		for _, v := range codeList {
			if code == v {
				return true
			}
		}
		return false
	}
	return true
}

func login(p graphql.ResolveParams) (interface{}, error) {
	phone := p.Args["phone"].(string)
	if allow := checkPhone(p, phone); allow != true {
		return false, fmt.Errorf("手机号码不可用: %s", phone)
	}
	code := p.Args["code"].(string)
	if allow := checkCode(p, code); allow != true {
		return false, fmt.Errorf("口令不可用: %s", code)
	}
	var user = types.Myself{}
	user.Phone = phone
	user.Token = "token_" + code

	return user, nil
}
