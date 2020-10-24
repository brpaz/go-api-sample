package main

import (
	"flag"
	"github.com/brpaz/go-api-sample/test/testutil"
	"log"
	"os"
	"os/exec"
)

func runTests() error {
	cmd := exec.Command("gotestsum", "--format", "testname", "--", "-v", "--tags", "integrationdb", "-p", "1", "./...")
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
