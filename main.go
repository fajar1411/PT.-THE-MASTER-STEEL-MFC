package main

import (
	"fmt"
	"master/config"
	fr "master/faktoryandroute"
	"master/loggrust"
	"master/migrasi"
	"master/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	loggrust.ConfigureLogger()
	migrasi.MigrateDB(db)

	e := echo.New()
	e.Use(loggrust.LoggerMiddleware)
	fr.FaktoryAndRoute(e, db)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%d", cfg.SERVER_PORT)))
}
