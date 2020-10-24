package context

import (
	"github.com/brpaz/go-api-sample/internal/todo"
	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"gorm.io/gorm"
)

// ApiContext The main struct
type DbContext struct {
	db *gorm.DB
}

// NewDBContext creates a new instance of the the DB context
func NewDBContext(db *gorm.DB) *DbContext {
	return &DbContext{
		db: db,
	}
}

// InitializeScenario
func (ctx *DbContext) InitializeScenario(s *godog.ScenarioContext) {
	s.BeforeScenario(ctx.truncateDB)

	s.Step(`^I have todos:$`, ctx.IHaveTodos)
}

func (ctx *DbContext) truncateDB(s *godog.Scenario) {
	_ = ctx.db.Exec("TRUNCATE TABLE todos")
}

// IHaveTodos Setups todos on the database
func (ctx *DbContext) IHaveTodos(data *messages.PickleStepArgument_PickleTable) error {
	for _, row := range data.Rows {
		for _, cell := range row.Cells {
			repo := todo.NewPgRepository(ctx.db)
			_, err := repo.Create(todo.NewTodo(cell.Value))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
