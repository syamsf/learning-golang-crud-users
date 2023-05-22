package controller

import (
	"errors"
	"net/http"
	"strconv"
	"syamsf/learning-gin/src/helper"
	"syamsf/learning-gin/src/model/web"
	"syamsf/learning-gin/src/modules/user/service"
	"syamsf/learning-gin/src/exception"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserControllerImpl(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) FindAll(ctx *gin.Context) {
	data := controller.UserService.FindAll()

	webResponse := web.WebResponse{
		Status: "OK",
		Code: http.StatusOK,
		Data: data,
	}

	helper.WriteResponseBody(webResponse, ctx)
}

func (controller *UserControllerImpl) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := controller.UserService.FindById(id)
	
	webResponse := web.WebResponse{
		Status: "OK",
		Code: http.StatusOK,
		Data: data,
	}


	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithError(http.StatusNotFound, exception.NewNotFoundError(err.Error()))
			return
		}

		customError := errors.New(err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, customError)
		return
	}

	helper.WriteResponseBody(webResponse, ctx)
}

func (controller *UserControllerImpl) Create(ctx *gin.Context) {
	data, err := controller.UserService.Create(ctx)

	webResponse := web.WebResponse{
		Status: "OK",
		Code: http.StatusOK,
		Data: data,
	}

	
	if err != nil {
		customError := errors.New(err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, customError)

		return
	}

	helper.WriteResponseBody(webResponse, ctx)
}

func (controller *UserControllerImpl) Update(ctx *gin.Context) {
	data, err := controller.UserService.Update(ctx)
	
	webResponse := web.WebResponse{
		Status: "OK",
		Code: http.StatusOK,
		Data: data,
	}

	
	if err != nil {
		customError := errors.New(err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, customError)
		
		return
	}

	helper.WriteResponseBody(webResponse, ctx)
}

func (controller *UserControllerImpl) Delete(ctx *gin.Context) {
	data, err := controller.UserService.Delete(ctx)
	
	webResponse := web.WebResponse{
		Status: "OK",
		Code: http.StatusOK,
		Data: data,
	}
	
	if err != nil {
		customError := errors.New(err.Error())
		ctx.AbortWithError(http.StatusInternalServerError, customError)
		
		return
	}

	helper.WriteResponseBody(webResponse, ctx)
}