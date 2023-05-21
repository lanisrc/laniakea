// Package k3d realizes a Kubernetes scope into a local cluster using k3d.
package k3d

import "errors"

// Kind define the kind name
const Kind = "k3d"

// Spec specify the k3d configuration.
type Spec struct {
	Image   string `yaml:"image,omitempty"`
	Servers int    `yaml:"servers,omitempty"`
	Agents  int    `yaml:"agents,omitempty"`
}

// Default supplies values for fields omitted from the source.
func (spec *Spec) Default() {
	spec.Servers = 1
}

// Validate checks the k3d configuration.
func (spec *Spec) Validate() error {
	if spec.Servers < 1 {
		return errors.New("servers must be at least 1")
	}
	if spec.Agents < 0 {
		return errors.New("agents must not be negative")
	}
	return nil
}
