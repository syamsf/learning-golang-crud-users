package helper

import (
	"github.com/gin-gonic/gin"
	"syamsf/learning-gin/src/model/web"
)

func WriteResponseBody(webResponse web.WebResponse, ctx *gin.Context) {
	ctx.JSON(webResponse.Code, gin.H{
		"status": webResponse.Status,
		"code": webResponse.Code,
		"data": webResponse.Data,
	})
}