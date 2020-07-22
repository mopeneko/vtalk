package infrastructure

import (
	"github.com/mopeneko/vtalk/backend/domain"
	"log"
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
	docRef := repo.conn.Collection("threads").Doc(threadID)

	_, err := docRef.Get(ctx)
	if err != nil {
		log.Printf("Failed to create: Post %+v at Thread %s | %+v\n", post, threadID, err)
		return err
	}

	_, _, err = docRef.Collection("posts").Add(ctx, post)
	if err != nil {
		log.Printf("Failed to create: Post %+v at Thread %s | %+v\n", post, threadID, err)
		return err
	}

	log.Printf("Created: Post %+v at Thread %s\n", post, threadID)
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
