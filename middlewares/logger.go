package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"time"

	"github.com/fsvxavier/unico/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var timeFormat = "02/Jan/2006:15:04:05 -0700"

// Logger is the logrus logger handler
func Logger(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// other handler can change c.Path so:
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		entry := map[string]interface{}{
			"http": map[string]interface{}{
				"latency":        latency, // time to process
				"request_method": c.Request.Method,
				"status_code":    statusCode,
				"path":           path,
			},
			"clientIP":   clientIP,
			"referer":    referer,
			"dataLength": dataLength,
			"userAgent":  clientUserAgent,
		}

		var logger utils.GenericLogger
		logger.Module = "server"
		logger.GetLogger()

		if len(c.Errors) > 0 {
			logger.LogIt("ERROR", c.Errors.ByType(gin.ErrorTypePrivate).String(), entry)
		} else {

			var bodyBytes []byte
			if c.Request.Body != nil {
				bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
			}
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

			msg := fmt.Sprintf("%s - %s [%s] \"%s %d\" %d %s \"%s\" (%dms) - body-send: %v", clientIP, time.Now().Format(timeFormat), c.Request.Method, path, statusCode, dataLength, referer, clientUserAgent, latency, string(bodyBytes))

			switch code := statusCode; {
			case code > 499:
				logger.LogIt("ERROR", msg, entry)
			case code > 399:
				logger.LogIt("WARN", msg, entry)
			default:
				if path != "/health" {
					logger.LogIt("INFO", msg, entry)
				}
			}
		}
	}
}
