package main

import (
	"syamsf/learning-gin/src/config"
	"syamsf/learning-gin/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	db := config.DB()

	routes.Api(engine, db)

	engine.Run()
}