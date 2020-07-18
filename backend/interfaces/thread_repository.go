package interfaces

type ThreadRepository interface {
	Create(string, string, string, string) error
}
