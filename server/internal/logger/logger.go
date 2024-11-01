package logger

import (
	"log"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap is ready logger to use(based on sugared zap logger)
var Zap *zap.SugaredLogger

func init() {
	Zap = newLogger()
}

var logDirPath = "./logs"

// newLogger creates new zep sugared logger
func newLogger() *zap.SugaredLogger {
	// creating log dir path(if not exist)
	if err := os.Mkdir(logDirPath, os.ModePerm); err != nil && !os.IsExist(err) {
		log.Fatalf("cannot create logger path: %s", err)
	}

	// creating log file(if not exist)
	logFile, err := os.OpenFile(filepath.Join(logDirPath, "server.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("cannot open logger file: %s", err)
	}

	// creating encoder config
	var cfg zapcore.EncoderConfig
	if os.Getenv("APP_ENV") == "development" {
		cfg = zap.NewDevelopmentEncoderConfig()
	} else {
		cfg = zap.NewProductionEncoderConfig()
	}
	// time layout 2006-01-02T15:04:05.000Z0700
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewJSONEncoder(cfg)

	//colorized output
	cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(cfg)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(logFile), zapcore.DebugLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
	)

	return zap.New(core).Sugar()
}
