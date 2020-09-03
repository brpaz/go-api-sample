package app

type App struct {
}

// New Creates a new instance of the application
func (*App) New() *App {
	return &App{}
}

func (*App) Start() error {
	return nil
}
