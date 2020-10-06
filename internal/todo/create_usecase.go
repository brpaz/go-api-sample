package todo

type CreateUseCase interface {
	Execute(request CreateTodo) (Todo, error)
}

type createUseCase struct {
	repo Repository
}

func NewCreateUseCase(repo Repository) createUseCase {
	return createUseCase{
		repo: repo,
	}
}

// Execute executes the creation of a new todo
func (uc createUseCase) Execute(request CreateTodo) (Todo, error) {

	todoEntity := NewTodo(request.Description)
	createdTodo, err := uc.repo.Create(todoEntity)

	if err != nil {
		return Todo{}, err
	}

	return createdTodo, nil
}
