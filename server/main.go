package server

import (
	"log"
	"os"
	"github.com/gin-gonic/gin"
	handler "github.com/coy102/go-starter/internal/handlers"
)

// Run web server
func Run() {
	handler.InitEnv()

	HOST, PORT := os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")

	// Initialize gin instance
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default() 

	r.GET("/ping", handler.Ping())

	log.Println("Running @ http://"+ HOST + ":" + PORT)

	r.Run(HOST + ":" + PORT)
}
