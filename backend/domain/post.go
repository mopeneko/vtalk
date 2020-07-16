package domain

type Post struct {
	Name      string `firestore:"name"`
	Email     string `firestore:"email"`
	CreatedAt int    `firestore:"created_at"`
	Content   string `firestore:"content"`
}
