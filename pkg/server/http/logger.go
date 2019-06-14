package http

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"time"
)

//Logger is a middlerware that logs all requests
func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			start := time.Now()
			var err error
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			reqSize := req.Header.Get(echo.HeaderContentLength)
			if reqSize == "" {
				reqSize = "0"
			}

			fields := logrus.Fields{
				"path":          req.RequestURI,
				"method":        req.Method,
				"status":        res.Status,
				"request_size":  reqSize,
				"response_size": res.Size,
				"request_id":    res.Header().Get(echo.HeaderXRequestID),
				"duration":      stop.Sub(start).String(),
				"error":         err,
			}

			if err == nil {
				fields["error"] = ""
			}

			logrus.WithFields(fields).Info("request")

			return err
		}
	}
}
