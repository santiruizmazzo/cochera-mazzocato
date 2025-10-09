package endtoend

import (
	"cochera/domain"
	"cochera/tests/utils"

	"net/http"
	"testing"
)

func TestCreateTenantWithMissingAttributes_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()

	expectedTenant := map[string]string{
		"name":        "Carl",
		"last_name":   "Johnson",
		"entry_month": "01-2025",
	}

	response, err := testAPI.CreateTenant(expectedTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertResponseContains(responseMap, "detail", "el DNI es obligatorio", t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)
}

func TestCreateTenantWithDuplicateDNI_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()
	existingTenant := domain.NewTenantBuilder().Build()

	response, err := testAPI.CreateTenant(existingTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	newTenant := domain.NewTenantBuilder().WithEmail("another@email.com").Build()

	response, err = testAPI.CreateTenant(newTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusConflict, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "el DNI ya existe", t)
}

func TestCreateTenantWithStringDNI_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()

	newTenant := map[string]any{
		"dni":         "hola",
		"name":        "Toni",
		"last_name":   "Cipriani",
		"entry_month": "02-2023",
	}

	response, err := testAPI.CreateTenant(newTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "el DNI debe ser un número entero", t)
}

func TestCreateTenantWithNegativeDNI_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()

	newTenant := map[string]any{
		"dni":         -10,
		"name":        "Victor",
		"last_name":   "Vance",
		"entry_month": "02-2023",
	}

	response, err := testAPI.CreateTenant(newTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "el DNI debe estar entre 1 y 4294967295", t)
}

func TestCreateTenantWithZeroDNI_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()
	newTenant := domain.NewTenantBuilder().WithDNI(0).Build()

	response, err := testAPI.CreateTenant(newTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "el DNI debe estar entre 1 y 4294967295", t)
}

func TestCreateTenantWithPhoneWithoutPlusSign_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()
	newTenant := map[string]any{
		"id":          1,
		"dni":         12345678,
		"name":        "Huang",
		"last_name":   "Lee",
		"address":     "22 Beer Heights St.",
		"phone":       "85718852",
		"email":       "huang@lee.com",
		"entry_month": "01-2025",
	}

	response, err := testAPI.CreateTenant(newTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "el teléfono debe comenzar con un símbolo de +", t)
}

func TestCreateTenantWithPhoneWithoutNumbers_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()
	newTenant := map[string]any{
		"id":          1,
		"dni":         12345678,
		"name":        "Huang",
		"last_name":   "Lee",
		"address":     "22 Beer Heights St.",
		"phone":       "+hola, que tal",
		"email":       "huang@lee.com",
		"entry_month": "01-2025",
	}

	response, err := testAPI.CreateTenant(newTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "el teléfono solo puede tener números", t)
}

func TestCreateTenantWithPhoneWithTooManyNumbers_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()
	newTenant := map[string]any{
		"id":          1,
		"dni":         12345678,
		"name":        "Huang",
		"last_name":   "Lee",
		"address":     "22 Beer Heights St.",
		"phone":       "+5434424072773442407277",
		"email":       "huang@lee.com",
		"entry_month": "01-2025",
	}

	response, err := testAPI.CreateTenant(newTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "el teléfono debe tener 15 dígitos como máximo", t)
}

func TestCreateTenantWithPhoneFullOfZeros_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()
	newTenant := map[string]any{
		"id":          1,
		"dni":         12345678,
		"name":        "Huang",
		"last_name":   "Lee",
		"address":     "22 Beer Heights St.",
		"phone":       "+000000000",
		"email":       "huang@lee.com",
		"entry_month": "01-2025",
	}

	response, err := testAPI.CreateTenant(newTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "el teléfono no puede estar lleno de ceros", t)
}

func TestCreateTenantWithInvalidEmail_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()
	newTenant := domain.NewTenantBuilder().WithEmail("hi, there").Build()

	response, err := testAPI.CreateTenant(newTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "el email debe seguir el formato estándar", t)
}

func TestCreateTenantWithVeryLargeEmail_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()
	newTenant := domain.NewTenantBuilder().WithEmail("sonny@forelliiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii.com").Build()

	response, err := testAPI.CreateTenant(newTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "el email debe tener 100 caracteres como máximo", t)
}

func TestCreateTenantWithDuplicateEmail_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()
	existingTenant := domain.NewTenantBuilder().Build()

	response, err := testAPI.CreateTenant(existingTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	newTenant := domain.NewTenantBuilder().WithDNI(1).Build()
	response, err = testAPI.CreateTenant(newTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusConflict, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "el email ya está en uso", t)
}

func TestCreateTenantWithInvalidFormatEntryMonth_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()
	newTenant := map[string]any{
		"dni":         15151515,
		"name":        "Trevor",
		"last_name":   "Phillips",
		"email":       "a@b.com",
		"entry_month": "03-20255555",
	}

	response, err := testAPI.CreateTenant(newTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "imposible procesar este año", t)
}

func TestCreateTenantWithReallyLargeName_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()
	newTenant := domain.NewTenantBuilder().WithName("Trevorrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr").Build()

	response, err := testAPI.CreateTenant(newTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "el nombre/apellido debe tener 50 caracteres como máximo", t)
}

func TestCreateTenantWithReallyLargeLastName_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()
	newTenant := domain.NewTenantBuilder().WithLastName("Phillipssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss").Build()

	response, err := testAPI.CreateTenant(newTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "el nombre/apellido debe tener 50 caracteres como máximo", t)
}

func TestCreateTenantWithReallyLargeAddress_EndToEnd(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	testAPI.ClearTenants()
	newTenant := domain.NewTenantBuilder().WithAddress("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa").Build()

	response, err := testAPI.CreateTenant(newTenant)
	if err != nil {
		t.Fatalf("Failed creating tenant: %v", err)
	}

	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			t.Fatalf("Failed closing response body: %v", cerr)
		}
	}()

	responseMap := utils.CreateMapFromBody(response.Body, t)

	utils.AssertStatusCodeIs(http.StatusBadRequest, response.StatusCode, t)

	utils.AssertResponseContains(responseMap, "detail", "el domicilio debe tener 100 caracteres como máximo", t)
}
