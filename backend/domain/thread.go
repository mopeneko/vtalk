package domain

type Thread struct {
	Title     string `firestore:"title"`
	Posts     []Post `firestore:"posts"`
	UpdatedAt int    `firestore:"updated_at"`
}
