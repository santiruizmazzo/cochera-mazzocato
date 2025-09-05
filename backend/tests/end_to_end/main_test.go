package endtoend

import (
	testingapi "cochera/tests/utils/testing_api"
	"log"
	"os"
	"testing"
)

var testApi *testingapi.TestingAPI
var err error

func TestMain(m *testing.M) {
	code := 1
	defer func() {
		os.Exit(code)
	}()

	testApi, err = testingapi.NewTestingAPI()
	if err != nil {
		log.Println("Could not create testing API: ", err)
		return
	}

	defer testApi.Stop()
	testApi.Run()

	code = m.Run()
}
