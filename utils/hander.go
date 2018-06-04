package utils

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/graphql-go/handler"
	"github.com/opera443399/cicd-backend/log"
)

//SvcHandler service handler
type SvcHandler struct {
	Service ServicesCode
	Schema  graphql.Schema
	handler *handler.Handler
}

//NewGQLHandler new graphql handler
func NewGQLHandler(service ServicesCode, Schema graphql.Schema) *SvcHandler {
	return &SvcHandler{
		Service: service,
		Schema:  Schema,
		handler: handler.New(&handler.Config{
			Schema: &Schema,
			Pretty: true,
			GraphiQL: true,
		}),
	}
}

func (h SvcHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if tokens, ok := r.URL.Query()["token"]; ok {
		if len(tokens) <= 0 {
			GQLReturnError(ErrTokenInvalid, w)
			return
		}
		if tokens[0] != "token_test" {
			log.Error("token invalid: %s", tokens[0])
			GQLReturnError(ErrTokenInvalid, w)
			return
		}
	} else {
		log.Error("token not found")
		GQLReturnError(ErrTokenNotFound, w)
		return
	}
	h.handler.ServeHTTP(w, r)
}

//GQLReturnError retrun graphql error
func GQLReturnError(err error, w http.ResponseWriter) {
	//InfoRaw(string(debug.Stack()))
	result := &graphql.Result{
		Errors: gqlerrors.FormatErrors(err),
	}
	w.WriteHeader(http.StatusOK)
	buff, _ := json.MarshalIndent(result, "", "\t")
	w.Write(buff)
}
