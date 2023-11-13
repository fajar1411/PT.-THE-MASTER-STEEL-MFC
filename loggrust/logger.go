package loggrust

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log = logrus.New()

func ConfigureLogger() {

	log.SetOutput(&lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    50,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})

	log.SetFormatter(&logrus.TextFormatter{})
	log.SetLevel(logrus.DebugLevel)
}

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		if err := next(c); err != nil {
			c.Error(err)
		}
		log.WithFields(logrus.Fields{
			"status":  res.Status,
			"method":  req.Method,
			"path":    req.URL.Path,
			"ip":      c.RealIP(),
			"latency": res.Header().Get("X-Response-Time"),
		}).Info("Request handled")
		return nil
	}
}
