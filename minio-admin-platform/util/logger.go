package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)
var (
	Logger *logrus.Logger
)

func init(){
	Logger = InitLogger()
}

func InitLogger() *logrus.Logger{
	log := logrus.New()
	pathMap := lfshook.PathMap{
		logrus.InfoLevel: "logs/info.log",
		logrus.ErrorLevel: "logs/error.log",
		logrus.WarnLevel: "logs/warn.log",
	}
	log.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{}))
	return log
}

func LoggerHandler(log *logrus.Logger) gin.HandlerFunc{
	return func(c *gin.Context){
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknow"
		}
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}
		entry := logrus.NewEntry(log).WithFields(logrus.Fields{
			//"hostname":   hostname,
			//"statusCode": statusCode,
			//"latency":    latency, // time to process
			//"clientIP":   clientIP,
			//"method":     c.Request.Method,
			//"path":       path,
			//"referer":    referer,
			//"dataLength": dataLength,
			//"userAgent":  clientUserAgent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			//msg := fmt.Sprintf("%s - %s [%s] \"%s %s\" %d %d \"%s\" \"%s\" (%dms)", clientIP, hostname, time.Now().Format("02/Jan/2006:15:04:05 -0700"), c.Request.Method, path, statusCode, dataLength, referer, clientUserAgent, latency)
			msg := fmt.Sprintf("[%s] %s - %s  \"%s %s\" %d %d \"%s\" \"%s\" (%dms)", time.Now().Format("02/Jan/2006:15:04:05 -0700"),clientIP, hostname, c.Request.Method, path, statusCode, dataLength, referer, clientUserAgent, latency)
			if statusCode > 499 {
				entry.Error(msg)
			} else if statusCode > 399 {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}
	}
}
