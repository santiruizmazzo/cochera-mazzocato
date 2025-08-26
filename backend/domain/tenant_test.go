package domain

import (
	"fmt"
	"testing"
)

func TestTenantCreatedFromJSON(t *testing.T) {
	const expectedDni uint32 = 43295798
	const expectedName string = "Santiago"
	const expectedLastName string = "Ruiz Mazzocato"
	const expectedAddress string = "Roseti 745"
	const expectedPhone string = "543442407277"
	const expectedEmail string = "santimazzo98@gmail.com"
	const expectedEntryMonth string = "05-2025"

	jsonString := fmt.Sprintf(`{"dni":%v,"name":"%s","last_name":"%s","address":"%s","phone":"%s","email":"%s","entry_month":"%s"}`, expectedDni, expectedName, expectedLastName, expectedAddress, expectedPhone, expectedEmail, expectedEntryMonth)
	tenant, err := NewTenantFromJSON([]byte(jsonString))
	if err != nil {
		t.Fatal("Error creando inquilino desde JSON: ", err)
	}

	if tenant.DNI != expectedDni {
		t.Fatalf("Esperado DNI %v, obtenido %v", expectedDni, tenant.DNI)
	}

	if tenant.Name != expectedName {
		t.Fatalf("Esperado nombre %s, obtenido %s", expectedName, tenant.Name)
	}

	if tenant.LastName != expectedLastName {
		t.Fatalf("Esperado apellido %s, obtenido %s", expectedLastName, tenant.LastName)
	}

	if tenant.Address != expectedAddress {
		t.Fatalf("Esperado domicilio %s, obtenido %s", expectedAddress, tenant.Address)
	}

	if tenant.Phone != expectedPhone {
		t.Fatalf("Esperado telefono %s, obtenido %s", expectedPhone, tenant.Phone)
	}

	if tenant.Email != expectedEmail {
		t.Fatalf("Esperado email %s, obtenido %s", expectedEmail, tenant.Email)
	}

	expectedMonthOfYear, err := NewMonthOfYearFromString(expectedEntryMonth)
	if tenant.EntryMonth != expectedMonthOfYear || err != nil {
		t.Fatalf("Esperado mes de entrada %s, obtenido %s", expectedEntryMonth, tenant.EntryMonth)
	}
}
