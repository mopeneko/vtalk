package infrastructure

import (
	"errors"
	"github.com/labstack/echo/v4"
	"os"
)

var router *echo.Echo

func init() {
	router = echo.New()

	/*
		handler, err := initDBHandler()
		if err != nil {
			log.Fatalf("failed to initialize DB handler: %+v\n", err)
		}
	*/
}

func RunRouter() {
	router.Logger.Fatal(router.Start(":1323"))
}

func initDBHandler() (*FirestoreHandler, error) {
	projectID := os.Getenv("VTALK_PROJECT_ID")
	if len(projectID) == 0 {
		return nil, errors.New("environment variable 'VTALK_PROJECT_ID' is not set")
	}
	return NewFirestoreHandler(projectID)
}
