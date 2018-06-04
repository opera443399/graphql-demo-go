package main

import (
	"github.com/opera443399/cicd-backend/utils"
)

func main() {
	utils.NewGQLDaemon(utils.SvcGateway, Schema).Listen()
}
