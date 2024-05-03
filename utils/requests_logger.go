package utils

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupHTTPRequestsLogger(app *echo.Echo, logFilePath string) error {
	logFile, err := os.OpenFile(
		logFilePath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		return err
	}

	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logFile,
		Format: "${time_rfc3339} ${status} ${method} ${uri} ${latency_human}\n",
	}))

	return nil
}
