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

func (s *Spec) UnmarshalYAML(value *yaml.Node) error {
	var decoded struct {
		Backend  backend.Backend `yaml:"backend,omitempty"`
		Features []Laniscope     `yaml:"features,omitempty"`
	}
	if err := value.Decode(&decoded); err != nil {
		return err
	}
	s.Backend  = decoded.Backend
	s.Features = decoded.Features
	s.raw      = *value
	return nil
}

func (s Spec) Node() yaml.Node {
	return s.raw
}
