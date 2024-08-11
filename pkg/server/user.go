package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/chanudinho/go-api-example/pkg/user/controller"
)

func addUserRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	userCtrl := controller.New(db)
	user := rg.Group("/user")

	user.POST("", userCtrl.CreateUser)
	user.POST("/auth", userCtrl.Login)
}
