package di

import (
	"github.com/brpaz/go-api-sample/internal"
	"github.com/brpaz/go-api-sample/internal/config"
	"github.com/brpaz/go-api-sample/internal/db"
	"github.com/brpaz/go-api-sample/internal/http/healthcheck"
	"github.com/brpaz/go-api-sample/internal/todo"
	dic "github.com/sarulabs/di"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	ServiceDB = "db"
	ServiceLogger = "logger"
	ServiceTodoRepository = "todo.repository"
	ServiceTodoCreateHandler = "todo.handler.create"
	ServiceTodoListHandler = "todo.handler.list"
	ServiceTodoCreateUseCase = "todo.usecase.create"
	ServiceTodoListUseCase = "todo.usecase.list"
	ServiceHealcheckHandler = "healthcheck.handler"
)

func getServiceDefinitions(config config.Config) []dic.Definition {
	var defs = []dic.Definition{
		{
			Name:  ServiceDB,
			Scope: dic.App,
			Build: func(ctn dic.Container) (interface{}, error) {
				logger := ctn.Get(ServiceLogger).(*zap.Logger)
				return db.GetConnection(config, logger)
			},
		},
		{
			Name: ServiceTodoRepository,
			Build: func(ctn dic.Container) (interface{}, error) {
				dbConn := ctn.Get(ServiceDB).(*gorm.DB)
				return todo.NewPgRepository(dbConn),  nil
			},
		},
		{
			Name: ServiceTodoListUseCase,
			Build: func(ctn dic.Container) (interface{}, error) {
				repo := ctn.Get(ServiceTodoRepository).(*todo.PgRepository)
				return todo.NewListUseCase(repo), nil
			},
		},
		{

			Name: ServiceTodoCreateUseCase,
			Build: func(ctn dic.Container) (interface{}, error) {
				repo := ctn.Get(ServiceTodoRepository).(*todo.PgRepository)
				return todo.NewCreateUseCase(repo), nil
			},
		},
		{
			Name: ServiceTodoCreateHandler,
			Build: func(ctn dic.Container) (interface{}, error) {
				uc := ctn.Get(ServiceTodoCreateUseCase).(todo.CreateUseCase)
				return todo.NewCreateHandler(uc), nil
			},
		},
		{
			Name: 	ServiceHealcheckHandler,
			Build: func(ctn dic.Container) (interface{}, error) {
				dbConn := ctn.Get(ServiceDB).(*gorm.DB)
				return healthcheck.NewHandler(internal.AppName, internal.AppDescription, internal.BuildCommit,dbConn), nil
			},
		},
	}

	return defs
}
