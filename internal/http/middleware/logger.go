package middleware

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLogger is an example of echo middleware that logs requests using logger "zap"
func ZapLogger(log *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			fields := []zapcore.Field{
				zap.Int("status", res.Status),
				zap.String("latency", time.Since(start).String()),
				zap.String("correlation_id", id),
				zap.String("host", req.Host),
				zap.String("remote_ip", c.RealIP()),
				zap.String("user_agent", req.UserAgent()),
			}

			label := fmt.Sprintf("%s %s", req.Method, req.RequestURI)
			n := res.Status
			switch {
			case n >= 500:
				log.Error(label, fields...)
			case n >= 400:
				log.Warn(label, fields...)
			case n >= 300:
				log.Info(label, fields...)
			default:
				log.Info(label, fields...)
			}

			return nil
		}
	}
}
