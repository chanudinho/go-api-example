package server

import (
	"github.com/chanudinho/go-api-example/pkg/middlewares"
	"github.com/chanudinho/go-api-example/pkg/task/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func addTaskRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	taskCtrl := controller.New(db)
	task := rg.Group("/task")

	task.POST("", middlewares.VerifyToken("technician"), taskCtrl.Create)
	task.GET("", middlewares.VerifyToken("manager"), taskCtrl.FindAllManagerTasks)
	task.GET("/my", middlewares.VerifyToken(""), taskCtrl.FindAllTechTasks) // put vai ser technician e delete manager
	task.PUT("/:id", middlewares.VerifyToken("technician"), taskCtrl.Update)
}
