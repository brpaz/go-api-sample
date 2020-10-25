---
title: 'Dependency Injection'
position: 6
category: 'Application Structure'
fullscreen: false
---

The Dependency Injection container powered by the [sarulabs/di](https://github.com/sarulabs/di) package, provides a centralized way to manage the depdendencies between our structs.

You could do this wiring by hand, and many people do that with Go, but after working in some bigger projects, I found out that it starts to get a little messy, and a feel of being reinventing the wheel so for this project I decided to experiment with this package.

The service definitions are placed in `app/di/services.go`.

Ex:

```go
const (
	ServiceDB                = "db"
	ServiceLogger            = "logger"
	ServiceTodoRepository    = "todo.repository"
)

func getServiceDefinitions(config config.Config) []dic.Definition {
	var defs = []dic.Definition{
		{
			Name: ServiceDB,
			Build: func(ctn dic.Container) (interface{}, error) {
				logger := ctn.Get(ServiceLogger).(*zap.Logger)

				return db.GetConnection(config, logger)
			},
		},
		{
			Name: ServiceTodoRepository,
			Build: func(ctn dic.Container) (interface{}, error) {
				dbConn := ctn.Get(ServiceDB).(*gorm.DB)
				return todo.NewPgRepository(dbConn), nil
			},
		}
	}

	return defs
}
```

## Add a new service

To add a new service to the application, the only thing you need to do is to add it to the `defs` slice shown previously.

You can get the built service at runtime, using:

```go
di.Get("myservice").(*MyServiceStruct)
```

<alert>
The only package that should know about the DIC is the `app` package. For all the rest, the dependencies should be injected explicitly or passed as function arguments to avoid unecessary coupling with the DI Container.
</alert>
