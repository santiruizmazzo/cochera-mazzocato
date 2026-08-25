package unit

import (
	"cochera/application"
	"regexp"
	"testing"
)

func TestCurrentVersionMatchesSemver(t *testing.T) {
	semver := regexp.MustCompile(`^\d+\.\d+\.\d+$`)

	if !semver.MatchString(application.CurrentVersion()) {
		t.Fatalf("CurrentVersion() debe respetar el formato semver X.Y.Z, obtenido: %q", application.CurrentVersion())
	}
}
