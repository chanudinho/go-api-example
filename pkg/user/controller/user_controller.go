package controller

import (
	"net/http"

	"github.com/chanudinho/go-api-example/pkg/user/domain"
	"github.com/chanudinho/go-api-example/pkg/user/repository"
	"github.com/chanudinho/go-api-example/pkg/user/service"
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

func (ctrl *Controller) CreateUser(gc *gin.Context) {
	var params domain.User
	err := gc.ShouldBindJSON(&params)
	if err != nil {
		gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, err := ctrl.Service.Create(gc, params)
	if err != nil {
		gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	gc.JSON(http.StatusOK, user)
}

func (ctrl *Controller) Login(gc *gin.Context) {
	var params domain.User
	err := gc.ShouldBindJSON(&params)
	if err != nil {
		gc.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	token, err := ctrl.Service.Login(gc, params)
	if err != nil {
		gc.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	gc.JSON(http.StatusOK, gin.H{"token": token})
}
