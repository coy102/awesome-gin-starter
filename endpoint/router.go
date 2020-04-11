package endpoint

import (
	"regexp"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/coy102/go-starter/endpoint/v1"
	_ "github.com/coy102/go-starter/docs" // Generated docs
	logging "github.com/coy102/go-starter/utils/logging"

)

var rxURL = regexp.MustCompile(`^/regexp\d*`)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	// Logger middleware
	logging.LoggerFileSetup(gin.IsDebugging(), r)


	apiv1 := r.Group("/api/v1")
	apiv1.GET("/ping", v1.Ping)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}


