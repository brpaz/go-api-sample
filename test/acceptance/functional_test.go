// +build acceptance

package main

import (
	"flag"
	apiContext "github.com/brpaz/godog-api-context"
	"os"
	"testing"

	"github.com/DATA-DOG/godog"
)

var opt = godog.Options{
	Output: os.Stdout,
	Format: "progress", // can define default values,
	Strict: false,
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

func TestMain(m *testing.M) {
	flag.Parse()

	opt.Paths = flag.Args()

	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		apiContext.NewAPIContext(s, os.Getenv("APP_BASE_URL"))
	}, opt)

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}