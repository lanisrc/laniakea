package laniscope

import "gopkg.in/yaml.v3"

const APIVersion = "laniscope/v1alpha1"

// TypeMeta follows the Kubernetes apiVersion/kind convention.
type TypeMeta struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
}

// ObjectMeta is the source-owned subset of Kubernetes object metadata.
type ObjectMeta struct {
	Name        string            `yaml:"name"`
	Namespace   string            `yaml:"namespace,omitempty"`
	Labels      map[string]string `yaml:"labels,omitempty"`
	Annotations map[string]string `yaml:"annotations,omitempty"`
}

// Laniscope is the behavior-independent container.
type Laniscope struct {
	         TypeMeta   `yaml:",inline"`
	Metadata ObjectMeta `yaml:"metadata"`
	Spec     yaml.Node  `yaml:"spec,omitempty"`
}
