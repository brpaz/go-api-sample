package di

import (
	"github.com/brpaz/go-api-sample/internal/config"
	dic "github.com/sarulabs/di"
)

type Container dic.Container

// Build dependency injection container
func BuildContainer(config config.Config) Container {
	builder, err := dic.NewBuilder()
	if err != nil {
		panic(err)
	}

	for _, definition := range getServiceDefinitions(config) {
		if err := builder.AddDefinition(definition); err != nil {
			panic(err)
		}
	}

	return builder.Build()
}
