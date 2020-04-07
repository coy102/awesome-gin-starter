package app

import (
	"github.com/gin-gonic/gin"
)

// Gin Struct type
type Gin struct {
	C *gin.Context
}

// Response Struct type obj
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: httpCode,
		Msg:  GetMsg(errCode),
		Data: data,
	})
	return
}
