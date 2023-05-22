package exception

import (
  "net/http"
  "syamsf/learning-gin/src/helper"
  "syamsf/learning-gin/src/model/web"

  "github.com/gin-gonic/gin"
)

func ErrorHandler(ctx *gin.Context) {
  ctx.Next()

  lastError := ctx.Errors.Last()
  if lastError != nil {
    err := lastError.Err

    if notFoundError(ctx, err) {
      return
    }

    internalServerError(ctx, err)
  }
}

func notFoundError(ctx *gin.Context,  err interface{}) bool {
  exception, ok := err.(NotFoundError)

  if ok {
    webResponse := web.WebResponse{
      Code: http.StatusNotFound,
      Status: "Not Found",
      Data: exception.Message,
    }

    helper.WriteResponseBody(webResponse, ctx)
    ctx.Abort()
    return true
  }

  return false
}

func internalServerError(ctx *gin.Context, err interface{}) {
  exception, ok := err.(error)

  if ok {
    webResponse := web.WebResponse{
      Code: http.StatusInternalServerError,
      Status: "Internal Server Error",
      Data: exception.Error(),
    }

    ctx.Abort()
    helper.WriteResponseBody(webResponse, ctx)
  }
}
