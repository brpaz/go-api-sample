package di

import (
	"github.com/brpaz/go-api-sample/internal/config"
	dic "github.com/sarulabs/di"
	"go.uber.org/zap"
)

type Container dic.Container

// Build dependency injection container
func BuildContainer(config config.Config, logger *zap.Logger) Container {
	builder, err := dic.NewBuilder()

	if err != nil {
		panic(err)
	}

	if err := builder.Set("logger", logger); err != nil {
		panic(err)
	}

	for _, definition := range getServiceDefinitions(config) {
		if err := builder.AddDefinition(definition); err != nil {
			panic(err)
		}
	}

	return builder.Build()
}
