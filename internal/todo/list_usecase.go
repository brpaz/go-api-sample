package todo

type ListUseCase interface {
	Execute() ([]Todo, error)
}

type listUseCase struct {
	repo Repository
}

func NewListUseCase(repo Repository) *listUseCase {
	return &listUseCase{
		repo: repo,
	}
}

// Execute lists existing todos
func (uc *listUseCase) Execute() ([]Todo, error) {
	return uc.repo.FindAll()
}
