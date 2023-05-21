package service

import (
	"syamsf/learning-gin/src/modules/user/model"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	FindAll() []model.User
	FindById(id int) model.User
	Create(ctx *gin.Context) (*model.User, error)
	Update(ctx *gin.Context) (*model.User, error)
	Delete(ctx *gin.Context) (*model.User, error)
}