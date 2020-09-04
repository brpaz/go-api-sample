// +build acceptance

package main

import (
	"flag"
	"fmt"
	"github.com/brpaz/go-api-sample/internal/app"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/brpaz/go-api-sample/internal/config"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	apicontext "github.com/brpaz/godog-api-context"
	"github.com/cucumber/godog"
)

var opts = godog.Options{
	Output: os.Stdout,
	Format: "pretty",
	StopOnFailure: true,
}

var url string

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)
	flag.StringVar(&url, "app.url", "", "The URL under test")
}

func createApp() *app.App {

	if err := godotenv.Load("../../.env", "../../.env.test"); err != nil {
		log.Println("Failed to load .env files:" + err.Error())
	}

	logger := zap.NewNop()

	cfg, err := config.Load()

	if err != nil {
		log.Fatal("Failed to load application config:" + err.Error())
	}

	return app.New(cfg, logger)
}

func TestMain(m *testing.M) {

	flag.Parse()

	opts.Paths = flag.Args()

	// If the URL is not defined, we start a new instance of the App server for the tests.
	// If the URL is passed as an argument, we call that url directly.
	if url == "" {
		app := createApp()

		ts := httptest.NewServer(app)
		defer ts.Close()

		url = ts.URL
	}

	apiContext := apicontext.New(url)

	status := godog.TestSuite{
		Name:                "Acceptance Tests",
		ScenarioInitializer: apiContext.InitializeScenario,
		Options:             &opts,
	}.Run()

	os.Exit(status)
}
