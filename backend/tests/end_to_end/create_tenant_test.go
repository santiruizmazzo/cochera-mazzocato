package endtoend

import (
	"cochera/domain"
	"cochera/tests/utils"

	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestCreateTenantWithMissingAttributes_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testApi.ResetDB()

	jsonData, _ := json.Marshal(map[string]string{
		"name":        "Carl",
		"last_name":   "Johnson",
		"entry_month": "01-2025",
	})

	response, err := http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
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

	// setupExistingTenant(t)
	existingTenant := domain.NewTenantBuilder().Build()
	jsonData, _ := json.Marshal(existingTenant)

	_, err := http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
	}

	newTenant := domain.NewTenantBuilder().WithEmail("another@email.com").Build()
	jsonData, _ = json.Marshal(newTenant)

	response, err := http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
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

	testApi.ResetDB()
	jsonData, _ := json.Marshal(map[string]any{
		"dni":         "hola",
		"name":        "Toni",
		"last_name":   "Cipriani",
		"entry_month": "02-2023",
	})

	response, err := http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
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

	testApi.ResetDB()
	jsonData, _ := json.Marshal(map[string]any{
		"dni":         -10,
		"name":        "Victor",
		"last_name":   "Vance",
		"entry_month": "02-2023",
	})

	response, err := http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
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

	testApi.ResetDB()
	newTenant := domain.NewTenantBuilder().WithDNI(0).Build()
	jsonTenant, _ := json.Marshal(newTenant)

	response, err := http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonTenant))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
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

	testApi.ResetDB()
	newTenant := domain.NewTenantBuilder().WithPhone("85718852").Build()
	jsonTenant, _ := json.Marshal(newTenant)

	response, err := http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonTenant))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
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

	testApi.ResetDB()
	newTenant := domain.NewTenantBuilder().WithPhone("+hola, que tal").Build()
	jsonTenant, _ := json.Marshal(newTenant)

	response, err := http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonTenant))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
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

	testApi.ResetDB()
	newTenant := domain.NewTenantBuilder().WithPhone("+5434424072773442407277").Build()
	jsonTenant, _ := json.Marshal(newTenant)

	response, err := http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonTenant))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
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

	testApi.ResetDB()
	newTenant := domain.NewTenantBuilder().WithPhone("+000000000").Build()
	jsonTenant, _ := json.Marshal(newTenant)

	response, err := http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonTenant))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
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

	testApi.ResetDB()
	newTenant := domain.NewTenantBuilder().WithEmail("hi, there").Build()
	jsonTenant, _ := json.Marshal(newTenant)

	response, err := http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonTenant))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
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

	testApi.ResetDB()
	newTenant := domain.NewTenantBuilder().WithEmail("sonny@forelliiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii.com").Build()
	jsonTenant, _ := json.Marshal(newTenant)

	response, err := http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonTenant))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
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

	existingTenant := domain.NewTenantBuilder().Build()
	jsonData, _ := json.Marshal(existingTenant)

	_, err := http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
	}

	newTenant := domain.NewTenantBuilder().WithDNI(1).Build()
	jsonTenant, _ := json.Marshal(newTenant)

	response, err := http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonTenant))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
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

	testApi.ResetDB()
	jsonData, _ := json.Marshal(map[string]any{
		"dni":         15151515,
		"name":        "Trevor",
		"last_name":   "Phillips",
		"email":       "a@b.com",
		"entry_month": "03-20255555",
	})

	response, err := http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
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

func TestCreateTenantWithReallyLargeName_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testApi.ResetDB()
	newTenant := domain.NewTenantBuilder().WithName("Trevorrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr").Build()
	jsonTenant, _ := json.Marshal(newTenant)

	response, err := http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonTenant))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "name must be 50 characters long at max", t)
}

func TestCreateTenantWithReallyLargeLastName_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testApi.ResetDB()
	newTenant := domain.NewTenantBuilder().WithLastName("Phillipssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss").Build()
	jsonTenant, _ := json.Marshal(newTenant)

	response, err := http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonTenant))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "last name must be 50 characters long at max", t)
}

func TestCreateTenantWithReallyLargeAddress_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testApi.ResetDB()
	newTenant := domain.NewTenantBuilder().WithAddress("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa").Build()
	jsonTenant, _ := json.Marshal(newTenant)

	response, err := http.Post(testApi.GetTenantsRoute(), "application/json", bytes.NewBuffer(jsonTenant))
	if err != nil {
		t.Fatalf("Failed sending POST request to %s: %v", testApi.GetTenantsRoute(), err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "address must be 100 characters long at max", t)
}
