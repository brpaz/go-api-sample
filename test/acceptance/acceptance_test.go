// +build acceptance

package main

import (
	"flag"
	"fmt"
	"github.com/brpaz/go-api-sample/internal/app"
	testContext "github.com/brpaz/go-api-sample/test/acceptance/context"
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

	logger, _ := zap.NewDevelopment()

	cfg, err := config.Load()

	if err != nil {
		log.Fatal("Failed to load application config:" + err.Error())
	}

	cfg.Env = config.EnvTest // force test env.
	cfg.DB.Database = fmt.Sprintf("%s_test", cfg.DB.Database)

	return app.New(cfg, logger)
}

func TestMain(m *testing.M) {

	flag.Parse()

	opts.Paths = flag.Args()

	db, err := testutil.SetupDB()

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

	status := godog.TestSuite{
		Name: "Acceptance Tests",
		ScenarioInitializer: func(s *godog.ScenarioContext) {
			apicontext.New(url).WithDebug(true).InitializeScenario(s)
			testContext.NewDBContext(db).InitializeScenario(s)
		},
		Options: &opts,
	}.Run()

	os.Exit(status)
}
