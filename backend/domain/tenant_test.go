package domain

import (
	"fmt"
	"testing"
)

func TestTenantCreatedFromJSON(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	const expectedDni uint32 = 43295798
	const expectedName string = "Santiago"
	const expectedLastName string = "Ruiz Mazzocato"
	const expectedAddress string = "Roseti 745"
	const expectedPhone string = "543442407277"
	const expectedEmail string = "santimazzo98@gmail.com"
	const expectedEntryMonth string = "05-2025"

	jsonString := fmt.Sprintf("{'dni':%v,'name':%s,'last_name':%s,'address':%s,'phone':%s,'email':%s,'entry_month':%s}", expectedDni, expectedName, expectedLastName, expectedAddress, expectedPhone, expectedEmail, expectedEntryMonth)
	tenant, err := NewTenantFromJSON([]byte(jsonString))
	if err != nil {
		t.Error("Error creando inquilino desde JSON: ", err)
	}

	if tenant.dni != expectedDni {
		t.Errorf("Esperado DNI %v, obtenido %v", expectedDni, tenant.dni)
	}

	if tenant.name != expectedName {
		t.Errorf("Esperado nombre %s, obtenido %s", expectedName, tenant.name)
	}

	if tenant.lastName != expectedLastName {
		t.Errorf("Esperado apellido %s, obtenido %s", expectedLastName, tenant.lastName)
	}

	if tenant.address != expectedAddress {
		t.Errorf("Esperado domicilio %s, obtenido %s", expectedAddress, tenant.address)
	}

	if tenant.phone != expectedPhone {
		t.Errorf("Esperado telefono %s, obtenido %s", expectedPhone, tenant.phone)
	}

	if tenant.email != expectedEmail {
		t.Errorf("Esperado email %s, obtenido %s", expectedEmail, tenant.email)
	}

	expectedMonthOfYear, err := NewMonthOfYearFromString(expectedEntryMonth)
	if tenant.entryMonth != expectedMonthOfYear || err != nil {
		t.Errorf("Esperado mes de entrada %s, obtenido %s", expectedEntryMonth, tenant.entryMonth)
	}
}
