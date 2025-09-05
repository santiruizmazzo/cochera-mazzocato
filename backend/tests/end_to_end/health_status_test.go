package endtoend

import (
	"cochera/application/version"
	"cochera/tests/utils"
	"net/http"
	"testing"
)

func TestHealthStatus_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	response, err := http.Get(testApi.GetHealthStatusRoute())
	if err != nil {
		t.Fatalf("Failed sending GET request to %s: %v", testApi.GetHealthStatusRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusOK, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "status", "operational", t)

	utils.AssertResponseContains(responseMap, "version", version.Current(), t)
}
