package main

import (
  "syamsf/learning-gin/src/config"
  "syamsf/learning-gin/src/routes"
  "syamsf/learning-gin/src/exception"
  "github.com/gin-gonic/gin"
)

func main() {
  // TODO:
  // - Add wire DI
  // - Add migration
  // - Add unit test

  engine := gin.Default()
  db := config.DB()

  engine.Use(exception.ErrorHandler)
  routes.Api(engine, db)

  engine.Run()
}
