package main

import (
	"github.com/chanudinho/go-api-example/pkg/server"
)

func main() {
	r := server.InitRouter()

	err := r.Run(":9090")
	if err != nil {
		return
	}
}
