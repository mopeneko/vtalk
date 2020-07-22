package infrastructure

import (
	"cloud.google.com/go/firestore"
	"github.com/mopeneko/vtalk/backend/domain"
	"log"
)

type ThreadRepository struct {
	FirestoreHandler
}

func NewThreadRepository(handler FirestoreHandler) *ThreadRepository {
	return &ThreadRepository{handler}
}

func (repo *ThreadRepository) Create(title, name, email, content string) error {
	name = correctName(name)
	thread := &domain.Thread{
		Title: title,
		Posts: []domain.Post{
			{
				Name:      name,
				Email:     email,
				Content:   content,
				CreatedAt: (int)(firestore.ServerTimestamp),
			},
		},
		UpdatedAt: (int)(firestore.ServerTimestamp),
	}
	docRef, _, err := repo.conn.Collection("threads").Add(ctx, thread)
	if err != nil {
		log.Printf("Failed to create: Thread %s | %+v\n", title, err)
		return err
	}
	log.Printf("Created: Thread %s(ID:%s)\n", title, docRef.ID)
	return err
}
