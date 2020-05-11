package util

import (
	"github.com/gin-gonic/gin"
	"minio-admin/middleware"
)

var (
	Engine *gin.Engine
)

func init(){
	Engine = InitGin(LoggerHandler(InitLogger()))
}

func InitGin(LoggerHandler gin.HandlerFunc) *gin.Engine {
	engine := gin.New()
	engine.Use(LoggerHandler)
	engine.Use(middleware.Cors())
	engine.Use(gin.Recovery())
	return engine
}
