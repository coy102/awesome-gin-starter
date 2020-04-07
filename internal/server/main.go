package internal

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/coy102/go-starter/controller"
	"github.com/coy102/go-starter/pkg/handlers"
	setting "github.com/coy102/go-starter/internal/setting"
)

func init() {
	// logging.Setup()
	setting.Setup()
	handler.InitEnv()
}

// Run web server
func Run() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := controller.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HTTPPort)
	maxHeaderBytes := 1 << 20

	// Initialize gin instance
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
