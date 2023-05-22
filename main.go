package main

import (
	"syamsf/learning-gin/src/config"
	"syamsf/learning-gin/src/routes"
	"syamsf/learning-gin/src/exception"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	db := config.DB()
	
	engine.Use(exception.ErrorHandler)
	routes.Api(engine, db)

	engine.Run()
}