package infrastructure

import "github.com/labstack/echo/v4"

var router *echo.Echo

func init() {
	router = echo.New()
}

func RunRouter() {
	router.Logger.Fatal(router.Start(":1323"))
}
