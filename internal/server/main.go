package internal

import (
	"fmt"
	"net/http"
	"github.com/rs/zerolog/log"
	"github.com/gin-gonic/gin"
	"github.com/coy102/go-starter/endpoint"
	setting "github.com/coy102/go-starter/internal/setting"
)


func init() {
	setting.Setup()
}

// Run web server
func Run() {
	gin.SetMode(setting.ServerSetting.RunMode)

	router := endpoint.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HTTPPort)
	maxHeaderBytes := 1 << 20

	// Initialize gin instance
	server := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Info().Msgf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
