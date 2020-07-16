package infrastructure

import (
	"github.com/mopeneko/vtalk/backend/domain"
	"unicode/utf8"
)

type PostRepository struct {
	FirestoreHandler
}

func NewPostRepository(handler FirestoreHandler) *PostRepository {
	return &PostRepository{handler}
}

func (repo *PostRepository) Create(threadID string, post *domain.Post) error {
	post.Name = correctName(post.Name)
	_, _, err := repo.conn.
		Collection("threads").
		Doc(threadID).
		Collection("posts").
		Add(ctx, post)
	return err
}

func correctName(name string) string {
	if len(name) == 0 {
		return "名無しさん"
	}

	if utf8.RuneCountInString(name) > 20 {
		return (string)(([]rune)(name)[:20])
	}

	return name
}
