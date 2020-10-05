package todo

type CreateUseCase struct {
	repo Repository
}

func NewCreateUseCase(repo Repository) CreateUseCase {
	return CreateUseCase{
		repo: repo,
	}
}

func (uc CreateUseCase) Execute(todo CreateTodo) (Todo, error) {
	return Todo{}, nil
}
