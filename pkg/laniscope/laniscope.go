package laniscope

import (
	"gopkg.in/yaml.v3"

	"laniakea/pkg/backend"
)

type Laniscope struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata,omitempty"`
	Spec       Spec     `yaml:"spec,omitempty"`
}

type Metadata struct {
	Name string `yaml:"name"`
}

type Spec struct {
	Features []Laniscope     `yaml:"features,omitempty"`
	Backend  backend.Backend `yaml:"backend,omitempty"`
	raw      yaml.Node
}
