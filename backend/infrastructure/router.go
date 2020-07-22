package infrastructure

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/mopeneko/vtalk/backend/interfaces"
	"log"
	"os"
)

var router *echo.Echo

func init() {
	router = echo.New()

	handler, err := initDBHandler()
	if err != nil {
		log.Fatalf("failed to initialize DB handler: %+v\n", err)
	}

	threadRepository := NewThreadRepository(*handler)
	threadController := interfaces.NewThreadController(threadRepository)
	router.POST("/threads", generateHandlerFunc(threadController.Create))

	postRepository := NewPostRepository(*handler)
	postController := interfaces.NewPostController(postRepository)
	router.POST("/threads/:id/posts", generateHandlerFunc(postController.Create))
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

func generateHandlerFunc(controllerFunc func(interfaces.Context) error) func(echo.Context) error {
	return func(c echo.Context) error {
		return controllerFunc(c)
	}
}
