package main

import (
	"log"
	"os"

	"github.com/gustavolimam/pismo-teste/database"
	"github.com/gustavolimam/pismo-teste/web"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	database.ConnectDB(os.Getenv("DB_CS"))

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	web.RegisterRoutes(e)

	log.Println("Conexão HTTP aberta pela porta ", os.Getenv("PORT"))
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
