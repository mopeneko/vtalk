package usecase

type ThreadRepository interface {
	Create(string, string, string, string) error
}
