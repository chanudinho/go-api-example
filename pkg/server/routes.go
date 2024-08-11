package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var router = gin.Default()

func StartServer(db *gorm.DB) {
	getAllRoutes(db)
	router.Run()
}

func getAllRoutes(db *gorm.DB) {
	v1 := router.Group("v1")
	addUserRoutes(v1, db)
	addTaskRoutes(v1, db)
}
