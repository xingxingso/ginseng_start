package middlewares

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger(notlogged ...string) gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknow"
	}

	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {
			if raw != "" {
				path = path + "?" + raw
			}

			statusCode := c.Writer.Status()

			errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

			entry := logrus.WithFields(logrus.Fields{
				"channel":      "sys",
				"clientIP":     c.ClientIP(),
				"context":      nil,
				"errorMessage": errorMessage,
				"hostname":     hostname,
				// "latency":      time.Now().Sub(start).String(),
				"latency": time.Now().Sub(start),
				// "level":        "info", // had
				"method": c.Request.Method,
				// "msg":        "", // had
				"path":       path,
				"referer":    c.Request.Referer(),
				"statusCode": statusCode,
				// "time":       "2020-11-02 16:12:09", // had textFormat not print
			})

			if len(c.Errors) > 0 || statusCode >= http.StatusInternalServerError {
				entry.Error(errorMessage)
			} else if statusCode >= http.StatusBadRequest {
				entry.Warn()
			} else {
				entry.Info()
			}
		}
	}
}
