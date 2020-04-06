package controller

import (
	 _ "github.com/coy102/go-starter/docs" // Generate docs
	"github.com/coy102/go-starter/controller/v1"
	"github.com/coy102/go-starter/controller/v2"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	apiv2 := r.Group("/api/v2")

	apiv1.GET("/ping", v1.Ping)
	apiv2.GET("/ping", v2.Ping())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}


