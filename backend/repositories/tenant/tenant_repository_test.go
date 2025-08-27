package tenantrepo

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	code := m.Run()
	os.Exit(code)
}

func TestTenantRepositorySavesSuccessfully(t *testing.T) {

}
