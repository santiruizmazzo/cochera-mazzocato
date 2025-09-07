package endtoend

import (
	"cochera/application"
	"cochera/tests/utils"
	"net/http"
	"testing"
)

func TestHealthStatus_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	response, err := http.Get(testAPI.GetHealthStatusRoute())
	if err != nil {
		t.Fatalf("Failed sending GET request to %s: %v", testAPI.GetHealthStatusRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusOK, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "status", "operational", t)

	utils.AssertResponseContains(responseMap, "version", application.CurrentVersion(), t)
}
