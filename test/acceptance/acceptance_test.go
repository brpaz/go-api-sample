// +build acceptance

package main

import (
	"flag"
	"github.com/brpaz/go-api-sample/internal/app"
	"github.com/brpaz/go-api-sample/test/testutil"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/brpaz/go-api-sample/internal/config"
	apicontext "github.com/brpaz/godog-api-context"
	"github.com/cucumber/godog"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var opts = godog.Options{
	Output:        os.Stdout,
	Format:        "pretty",
	StopOnFailure: true,
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)
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

	cfg.Env = config.EnvTest // force test env.

	return app.New(cfg, logger)
}

func TestMain(m *testing.M) {

	flag.Parse()

	opts.Paths = flag.Args()

	_, err := testutil.SetupDB()

	if err != nil {
		log.Fatal(err)
	}

	// If the URL is not defined, we start a new instance of the App server for the tests.
	// If the URL is passed as an argument, we call that url directly.
	url := os.Getenv("APP_URL")
	if url == "" {
		appInstance := createApp()

		ts := httptest.NewServer(appInstance)
		defer ts.Close()

		log.Printf("Test application running on: %s \n", ts.URL)

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
