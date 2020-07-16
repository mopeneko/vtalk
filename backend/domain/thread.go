package domain

import "time"

type Thread struct {
	Title     string    `firestore:"title"`
	Posts     []Post    `firestore:"posts"`
	UpdatedAt time.Time `firestore:"updated_at"`
}
