package main

import (
	"flag"
	"github.com/brpaz/go-api-sample/test/testutil"
	"log"
	"os"
	"os/exec"
)

var outputFormat string

var verbose bool

func runTests() error {

	var args = []string{
		"--tags",
		"integrationdb",
		"-p",
		"1",
		"./...",
	}

	if verbose {
		args = append(args, "-v")
	}

	cmd := exec.Command("gotestsum", args...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}

// TestMain is the entry point to the application database integration tests.
// It is responsible to setup the database for tests by creating the test database and run the migrations on the new schema and launch the tests.
func main() {

	flag.StringVar(&outputFormat, "format", "testname", "Specifies the output test format. (See gotestsum available options)")
	flag.BoolVar(&verbose, "v", false, "Verbose mode")
	flag.Parse()

	if os.Getenv("APP_ENV") != "test" {
		log.Fatal("You can only run this command with APP_ENV=test")
	}

	log.Println("Setup testing database")

	_, err := testutil.SetupDB()

	if err != nil {
		log.Fatal(err)
	}

	if err := runTests(); err != nil {
		os.Exit(-1)
	}
}
