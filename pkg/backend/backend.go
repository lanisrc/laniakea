package backend

import "gopkg.in/yaml.v3"

type Backend struct {
	Kind string    `yaml:"kind"`
	Spec yaml.Node `yaml:"spec,omitempty"`
}
