package integration_test

import (
	"fmt"
	"os"
	"testing"
)

var (
	apiKey string
)

func TestMain(m *testing.M) {
	// Get API key
	apik, ok := os.LookupEnv("API_KEY")
	if !ok {
		fmt.Println("Can't run integration tests without an API key. Please request one at https://nvd.nist.gov/developers/request-an-api-key.")
		os.Exit(1)
	}
	apiKey = apik

	os.Exit(m.Run())
}
