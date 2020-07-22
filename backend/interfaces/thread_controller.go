package interfaces

import (
	"github.com/mopeneko/vtalk/backend/usecase"
	"net/http"
)

type ThreadController struct {
	interactor *usecase.ThreadInteractor
}

func NewThreadController(repo ThreadRepository) *ThreadController {
	interactor := usecase.NewThreadInteractor(repo)
	return &ThreadController{interactor}
}

type ThreadCreateRequest struct {
	Title   string `json:"title"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Content string `json:"content"`
}

type ThreadCreateResponse struct {
	Message string `json:"message"`
}

func (controller *ThreadController) Create(c Context) error {
	req := new(ThreadCreateRequest)
	res := new(ThreadCreateResponse)

	c.Bind(&req)

	err := controller.interactor.Create(req.Title, req.Name, req.Email, req.Content)
	if err != nil {
		res.Message = "内部エラーが発生しました。"
		return c.JSON(http.StatusInternalServerError, res)
	}
	return c.JSON(http.StatusOK, res)
}
