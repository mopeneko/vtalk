package interfaces

import "github.com/mopeneko/vtalk/backend/domain"

type PostInteractor struct {
	postRepository PostRepository
}

func NewPostInteractor(repo PostRepository) *PostInteractor {
	return &PostInteractor{repo}
}

func (interactor *PostInteractor) Create(threadID string, post *domain.Post) error {
	return interactor.postRepository.Create(threadID, post)
}
