package middleware

import (
	"doovvvblog/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Logger() gin.HandlerFunc {
	filePath := utils.AppConfig.Server.FilePath + "\\log\\log.log"
	logFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic("日志文件打开失败")
	}

	logger := logrus.New()
	logger.Out = logFile
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.SetOutput(&lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     7,
		Compress:   false,
	})
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		latencyTime := time.Since(startTime)
		entry := logger.WithFields(logrus.Fields{
			"status":  c.Writer.Status(),
			"method":  c.Request.Method,
			"uri":     c.Request.RequestURI,
			"ip":      c.ClientIP(),
			"latency": latencyTime.String(),
			"agent":   c.Request.UserAgent(),
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if c.Writer.Status() != http.StatusOK {
			entry.Warn()
		}
		entry.Info("Http Request")

	}
}
