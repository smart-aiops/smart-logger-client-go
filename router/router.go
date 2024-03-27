package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"smart/logger/client/demo/router/api"
	"smart/logger/client/demo/router/middleware"
)

// InitRouter ...
func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.AcclogSetUp())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	apiGroupV1 := r.Group("/api/v1")
	apiGroupV1.GET("/hello", api.Hello)
	return r
}
