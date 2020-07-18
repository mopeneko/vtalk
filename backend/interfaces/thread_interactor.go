package interfaces

type ThreadInteractor struct {
	threadRepository ThreadRepository
}

func NewThreadInteractor(repo ThreadRepository) *ThreadInteractor {
	return &ThreadInteractor{repo}
}

func (interactor *ThreadInteractor) Create(title, name, email, content string) error {
	return interactor.threadRepository.Create(title, name, email, content)
}
