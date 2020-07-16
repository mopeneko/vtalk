package domain

import (
	"time"
)

type Post struct {
	Name      string    `firestore:"name"`
	Email     string    `firestore:"email"`
	CreatedAt time.Time `firestore:"created_at"`
	Content   string    `firestore:"content"`
}
