package laniscope

import "gopkg.in/yaml.v3"

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
	Features []Laniscope `yaml:"features,omitempty"`
	raw      yaml.Node
}
