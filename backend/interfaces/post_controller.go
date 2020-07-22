package interfaces

import (
	"github.com/mopeneko/vtalk/backend/domain"
	"github.com/mopeneko/vtalk/backend/usecase"
	"net/http"
)

type PostController struct {
	interactor *usecase.PostInteractor
}

func NewPostController(repo PostRepository) *PostController {
	interactor := usecase.NewPostInteractor(repo)
	return &PostController{interactor}
}

type PostCreateRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Content string `json:"content"`
}

type PostCreateResponse struct {
	Message string `json:"message"`
}

func (controller *PostController) Create(c Context) error {
	req := new(PostCreateRequest)
	res := new(PostCreateResponse)
	c.Bind(req)
	threadID := c.Param("id")
	post := &domain.Post{
		Name:    req.Name,
		Email:   req.Email,
		Content: req.Content,
	}

	err := controller.interactor.Create(threadID, post)
	if err != nil {
		res.Message = "内部エラーが発生しました。"
		return c.JSON(http.StatusInternalServerError, res)
	}

	return c.JSON(http.StatusOK, res)
}
