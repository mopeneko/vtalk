package interfaces

import "github.com/mopeneko/vtalk/backend/domain"

type PostRepository interface {
	Create(string, *domain.Post) error
}
