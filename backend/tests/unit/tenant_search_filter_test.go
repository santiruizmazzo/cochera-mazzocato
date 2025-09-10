package unit

import (
	"cochera/application"
	"net/url"
	"testing"
)

func TestTenantFilterCreationFromQueryParams(t *testing.T) {
	queryParams := &url.Values{}
	queryParams.Set("name", "Troy")
	queryParams.Set("lastName", "Bolton")

	filter, err := application.NewTenantFilterFromQueryParams(queryParams)

	if err != nil {
		t.Fatal("Tenant filter creation should've succeded: ", err)
	}

	if filter.Name != "Troy" {
		t.Fatal("Expected Troy, got ", filter.Name)
	}

	if filter.LastName != "Bolton" {
		t.Fatal("Expected Bolton, got ", filter.LastName)
	}

}
