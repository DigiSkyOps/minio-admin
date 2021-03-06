package util

import (
	"github.com/gin-gonic/gin"
	"minio-admin/middleware"
)

var (
	Engine *gin.Engine
)

func init(){
	Engine = initGin(LoggerHandler(InitLogger()))
}

func initGin(LoggerHandler gin.HandlerFunc) *gin.Engine {
	engine := gin.New()
	engine.Use(LoggerHandler)
	engine.Use(middleware.Cors())
	engine.Use(gin.Recovery())
	return engine
}
