package v1

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/coy102/go-starter/utils/app"
	"github.com/coy102/go-starter/utils/e"
)

// @Summary Get multiple article tags
// @Produce json
// @Success 200 {string} app.Response
// @Failure 500 {string} app.Response
// @Router /v1/ping [get]
// @tags Ping
func Ping(c *gin.Context) {
	appG := app.Gin{C: c}

	appG.Response(http.StatusOK, e.Success, map[string]string{
		"message": "ping",
	})
}