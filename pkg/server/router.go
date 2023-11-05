package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/def", func(ctx *gin.Context) {
			fmt.Println("hello word")
		})
	}

	return r
}
