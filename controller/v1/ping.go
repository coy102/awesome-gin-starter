package v1

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// @Summary Get multiple article tags
// @Produce  json
// @Success 200 {string} app.Response
// @Failure 500 {string} app.Response
// @Router /api/v1/ping [get]
// @tags Tags
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}