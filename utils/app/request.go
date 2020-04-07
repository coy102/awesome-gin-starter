package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/rs/zerolog/log"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		log.Info().
			Str("Error", err.Key).
			Msg(err.Message)
	}

	return
}
