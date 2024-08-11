package main

import (
	"github.com/chanudinho/go-api-example/pkg/configs"
	"github.com/chanudinho/go-api-example/pkg/repository/mysql"
	"github.com/chanudinho/go-api-example/pkg/server"
)

func main() {
	configs := configs.LoadEnv()

	database := mysql.ConnectDatabase(configs)

	server.StartServer(database)
}
