package laniscope

import (
	"bytes"
	"path/filepath"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	// Load local-dev as baseline
	scope, err := Load(filepath.Join("..", "..", "demo", "local-dev", "laniscope.yaml"))
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if scope.APIVersion != APIVersion || scope.Kind != "Kubernetes" || scope.Metadata.Name != "local-dev" {
		t.Fatalf("Load() = %#v", scope)
	}

	// Change and save scope for round trip test
	scope.Metadata.Name = "round-trip"
	filename := filepath.Join(t.TempDir(), "laniscope.yaml")
	if err := scope.Save(filename); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	// Reload scope
	got, err := Load(filename)
	if err != nil {
		t.Fatalf("Load(saved file) error = %v", err)
	}
	if got.Metadata.Name != "round-trip" {
		t.Fatalf("metadata.name = %q, want round-trip", got.Metadata.Name)
	}

	// Compare YAML
	wantYAML, err := scope.Encode()
	if err != nil {
		t.Fatalf("Encode(original) error = %v", err)
	}
	gotYAML, err := got.Encode()
	if err != nil {
		t.Fatalf("Encode(saved) error = %v", err)
	}
	if !bytes.Equal(gotYAML, wantYAML) {
		t.Fatalf("round-trip YAML differs:\ngot:\n%s\nwant:\n%s", gotYAML, wantYAML)
	}
}
