package app

import (
	"net/http"
	e "github.com/coy102/go-starter/utils/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, e.Success
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, e.Error
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, e.InvalidParam
	}

	return http.StatusOK, e.Success
}
