package controller

import (
	"net/http"
	"strconv"

	"github.com/chanudinho/go-api-example/pkg/task/domain"
	"github.com/chanudinho/go-api-example/pkg/task/repository"
	"github.com/chanudinho/go-api-example/pkg/task/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	service.Service
}

func New(db *gorm.DB) *Controller {
	return &Controller{
		Service: service.New(repository.New(db)),
	}
}

func (ctrl *Controller) Create(gc *gin.Context) {
	var params domain.Task
	err := gc.ShouldBindJSON(&params)
	if err != nil {
		gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userID := gc.Value("userID").(uint)
	params.UserID = userID

	task, err := ctrl.Service.Create(gc, params)
	if err != nil {
		gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	gc.JSON(http.StatusOK, task)
}

func (ctrl *Controller) FindAllTechTasks(gc *gin.Context) {
	var query domain.FindAllRequest
	err := gc.ShouldBindQuery(&query)
	if err != nil {
		gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userID := gc.Value("userID").(uint)
	query.UserID = userID

	tasks, err := ctrl.Service.FindAll(gc, query)
	if err != nil {
		gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	gc.JSON(http.StatusOK, tasks)
}

func (ctrl *Controller) FindAllManagerTasks(gc *gin.Context) {
	var query domain.FindAllRequest
	err := gc.ShouldBindQuery(&query)
	if err != nil {
		gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	tasks, err := ctrl.Service.FindAll(gc, query)
	if err != nil {
		gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	gc.JSON(http.StatusOK, tasks)
}

func (ctrl *Controller) Update(gc *gin.Context) {
	taskID, err := strconv.ParseUint(gc.Param("id"), 10, 64)
	if err != nil {
		gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var params domain.UpdateRequest
	err = gc.ShouldBindJSON(&params)
	if err != nil {
		gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	params.ID = uint(taskID)
	userID := gc.Value("userID").(uint)
	params.UserID = userID

	task, err := ctrl.Service.Update(gc, params)
	if err != nil {
		gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	gc.JSON(http.StatusOK, task)
}
