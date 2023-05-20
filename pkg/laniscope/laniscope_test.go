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

func TestSpecShape(t *testing.T) {
	const base = "apiVersion: " + APIVersion + "\nkind: Test\nmetadata:\n  name: test\n"
	tests := map[string]struct {
		spec  string
		valid bool
	}{
		"absent":     {valid: true},
		"null":       {spec: "spec:\n", valid: true},
		"mapping":    {spec: "spec:\n  key: value\n", valid: true},
		"scalar":     {spec: "spec: value\n"},
		"sequence":   {spec: "spec:\n  - value\n"},
		"duplicates": {spec: "spec:\n  key: one\n  key: two\n"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := Decode([]byte(base + test.spec))
			if (err == nil) != test.valid {
				t.Fatalf("Decode() error = %v, valid = %v", err, test.valid)
			}
		})
	}
}
