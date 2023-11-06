package main

import (
	"github.com/chanudinho/go-api-example/pkg/configs"
	"github.com/chanudinho/go-api-example/pkg/repository/mysql"
	"github.com/chanudinho/go-api-example/pkg/server"
)

func main() {
	cfg := configs.Load()
	mysql.ConnectDatabase(cfg)
	r := server.InitRouter()

	err := r.Run(":9090")
	if err != nil {
		return
	}
}
