package tests

import (
	"cochera/tests/utils"

	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func setupExistingTenant(t *testing.T) {
	jsonData, _ := json.Marshal(map[string]any{
		"dni":         17888423,
		"name":        "Trevor",
		"last_name":   "Phillips",
		"entry_month": "09-2024",
		"email":       "trevor@phillips.com",
	})

	_, err := http.Post(testApi.GetTenantCreationRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantCreationRoute(), err)
	}
}

func TestCreateTenantWithMissingAttributes_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	jsonData, _ := json.Marshal(map[string]string{
		"name":        "Carl",
		"last_name":   "Johnson",
		"entry_month": "01-2025",
	})

	response, err := http.Post(testApi.GetTenantCreationRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantCreationRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertResponseContains(responseMap, "detail", "dni is required", t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)
}

func TestCreateTenantWithDuplicateDNI_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	setupExistingTenant(t)

	jsonData, _ := json.Marshal(map[string]any{
		"dni":         17888423,
		"name":        "Johnny",
		"last_name":   "Klebitz",
		"entry_month": "11-2025",
	})

	response, err := http.Post(testApi.GetTenantCreationRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantCreationRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusConflict, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "dni already exists", t)
}

func TestCreateTenantWithStringDNI_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	jsonData, _ := json.Marshal(map[string]any{
		"dni":         "hola",
		"name":        "Toni",
		"last_name":   "Cipriani",
		"entry_month": "02-2023",
	})

	response, err := http.Post(testApi.GetTenantCreationRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantCreationRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "dni must be a positive integer", t)
}

func TestCreateTenantWithNegativeDNI_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	jsonData, _ := json.Marshal(map[string]any{
		"dni":         -10,
		"name":        "Victor",
		"last_name":   "Vance",
		"entry_month": "02-2023",
	})

	response, err := http.Post(testApi.GetTenantCreationRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantCreationRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "dni must be a positive integer", t)
}

func TestCreateTenantWithZeroDNI_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	jsonData, _ := json.Marshal(map[string]any{
		"dni":         0,
		"name":        "Victor",
		"last_name":   "Vance",
		"entry_month": "02-2023",
	})

	response, err := http.Post(testApi.GetTenantCreationRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantCreationRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "dni must be a positive integer", t)
}

func TestCreateTenantWithPhoneWithoutPlusSign_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	jsonData, _ := json.Marshal(map[string]any{
		"dni":         999999,
		"name":        "Lance",
		"last_name":   "Vance",
		"phone":       "543442407277",
		"email":       "lance@vance.com",
		"entry_month": "02-2023",
	})

	response, err := http.Post(testApi.GetTenantCreationRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantCreationRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "phone must start with + sign", t)
}

func TestCreateTenantWithPhoneWithoutNumbers_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	jsonData, _ := json.Marshal(map[string]any{
		"dni":         1212121,
		"name":        "Phil",
		"last_name":   "Cassidy",
		"phone":       "+hola, que tal?",
		"email":       "phil@cassidy.gun",
		"entry_month": "02-2023",
	})

	response, err := http.Post(testApi.GetTenantCreationRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantCreationRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "phone must contain numbers only", t)
}

func TestCreateTenantWithPhoneWithTooManyNumbers_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	jsonData, _ := json.Marshal(map[string]any{
		"dni":         1212121,
		"name":        "Phil",
		"last_name":   "Cassidy",
		"phone":       "+5434424072773442407277",
		"email":       "phil@cassidy.gun",
		"entry_month": "02-2023",
	})

	response, err := http.Post(testApi.GetTenantCreationRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantCreationRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "phone must have 15 digits max", t)
}

func TestCreateTenantWithPhoneFullOfZeros_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	jsonData, _ := json.Marshal(map[string]any{
		"dni":         1212121,
		"name":        "Phil",
		"last_name":   "Cassidy",
		"phone":       "+000000000000",
		"email":       "phil@cassidy.gun",
		"entry_month": "02-2023",
	})

	response, err := http.Post(testApi.GetTenantCreationRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantCreationRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "phone cannot be full of zeroes", t)
}

func TestCreateTenantWithInvalidEmail_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	jsonData, _ := json.Marshal(map[string]any{
		"dni":         8888888,
		"name":        "Sonny",
		"last_name":   "Forelli",
		"email":       "hi, there",
		"entry_month": "03-2025",
	})

	response, err := http.Post(testApi.GetTenantCreationRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantCreationRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "email must follow standard format", t)
}

func TestCreateTenantWithVeryLargeEmail_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	jsonData, _ := json.Marshal(map[string]any{
		"dni":         8888888,
		"name":        "Sonny",
		"last_name":   "Forelli",
		"email":       "sonny@forelliiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii.com",
		"entry_month": "03-2025",
	})

	response, err := http.Post(testApi.GetTenantCreationRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantCreationRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "email must be 100 characters long at max", t)
}

func TestCreateTenantWithDuplicateEmail_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	jsonData, _ := json.Marshal(map[string]any{
		"dni":         8888888,
		"name":        "Trevor",
		"last_name":   "Phillips",
		"email":       "trevor@phillips.com",
		"entry_month": "03-2025",
	})

	response, err := http.Post(testApi.GetTenantCreationRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantCreationRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusConflict, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "email already in use", t)
}

func TestCreateTenantWithInvalidFormatEntryMonth_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	jsonData, _ := json.Marshal(map[string]any{
		"dni":         15151515,
		"name":        "Trevor",
		"last_name":   "Phillips",
		"email":       "a@b.com",
		"entry_month": "03-20255555",
	})

	response, err := http.Post(testApi.GetTenantCreationRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantCreationRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "entry month must have this format: MM-YYYY", t)
}
