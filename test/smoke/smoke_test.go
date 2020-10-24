// +build smoketests

package smoke

import (
	"flag"
	apicontext "github.com/brpaz/godog-api-context"
	"github.com/cucumber/godog"
	"log"
	"os"
	"testing"
)

var opts = godog.Options{
	Output:        os.Stdout,
	Format:        "pretty",
	StopOnFailure: true,
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)
}

func TestMain(m *testing.M) {

	flag.Parse()

	opts.Paths = flag.Args()

	url := os.Getenv("APP_URL")

	if url == "" {
		log.Fatal("APP_URL environment is not set")
	}

	apiContext := apicontext.New(url)

	status := godog.TestSuite{
		Name:                "Smoke Tests",
		ScenarioInitializer: apiContext.InitializeScenario,
		Options:             &opts,
	}.Run()

	os.Exit(status)
}
