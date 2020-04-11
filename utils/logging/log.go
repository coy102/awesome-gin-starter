package logging

import (
	"os"
	"regexp"
	"github.com/gin-gonic/gin"
	"github.com/coy102/go-starter/utils/file"
	"github.com/rs/zerolog"
	ginlogger "github.com/gin-contrib/logger"
)

var (
	rxURL = regexp.MustCompile(`^/regexp\d*`)
	// F pointer os file
	F *os.File
	runlogger *zerolog.Logger
)

// Logger struct
type Logger struct {
    logger *zerolog.Logger
}

// LoggerFileSetup write logs to .log file connected with gin router 
func LoggerFileSetup(isDebug bool, r *gin.Engine) zerolog.Logger {
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err = file.MustOpen(fileName, filePath)

    logLevel := zerolog.InfoLevel
    if isDebug {
        logLevel = zerolog.DebugLevel
    }
    zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(F).
		With().
		Timestamp().
		Logger()	

	if err != nil {
		logger.Fatal().
			Err(err).
			Msgf("logging.Setup err: %v", err)
		
	}

	
	r.Use(ginlogger.SetLogger(ginlogger.Config{
		Logger:         &logger,
		UTC:            true,
		SkipPathRegexp: rxURL,
	}))

	runlogger = &logger
	return logger
}

// Debug Log
func Debug(msg string) {
	runlogger.Debug().
			Msgf(msg)
}

// Info Log
func Info(msg string) {
	runlogger.Info().
			Msgf(msg)
}

// Error Log
func Error(msg string) {
	runlogger.Error().
			Msgf(msg)
}

// Fatal Log
func Fatal(msg string) {
	runlogger.Fatal().
			Msgf(msg)
}