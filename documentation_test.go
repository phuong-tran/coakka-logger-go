package coakka_logger

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGoCompatibilityFloorAndCurrentToolchainCI(t *testing.T) {
	module, err := os.ReadFile("go.mod")
	if err != nil {
		t.Fatalf("read go.mod: %v", err)
	}
	if !strings.Contains(string(module), "\ngo 1.18\n") {
		t.Fatal("go.mod must retain the tested Go 1.18 compatibility floor")
	}

	workflow, err := os.ReadFile(filepath.Join(".github", "workflows", "go-ci.yml"))
	if err != nil {
		t.Fatalf("read go-ci workflow: %v", err)
	}
	text := string(workflow)
	for _, required := range []string{"\"1.18.10\"", "- stable", "Check public docs language"} {
		if !strings.Contains(text, required) {
			t.Fatalf("go-ci workflow is missing toolchain matrix entry %q", required)
		}
	}
	if strings.Contains(text, "go-version-file: go.mod") {
		t.Fatal("go-ci must not conflate the module compatibility floor with the current CI toolchain")
	}
}
