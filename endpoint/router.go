package endpoint

import (
	"os"
	"regexp"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/coy102/go-starter/endpoint/v1"
	_ "github.com/coy102/go-starter/docs" // Generated docs
)

var (
	rxURL = regexp.MustCompile(`^/regexp\d*`)
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	
	// Validate gin logger for debugging mode
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// log Setup output
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)
	
	r := gin.New()
	r.Use(gin.Recovery())

	// Add a logger middleware
	subLog := zerolog.New(os.Stdout).With().
		Logger()

	r.Use(logger.SetLogger(logger.Config{
		Logger:         &subLog,
		UTC:            true,
		SkipPathRegexp: rxURL,
	}))


	apiv1 := r.Group("/api/v1")
	apiv1.GET("/ping", v1.Ping)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}


