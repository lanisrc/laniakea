package laniscope

import (
	"errors"
	"fmt"
	"strings"
	"gopkg.in/yaml.v3"
)

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

// Validate checks behavior-independent envelope rules.
// Supported kinds and versions are checked by a Scheme.
func (s Laniscope) Validate() error {
	gv := strings.Split(s.APIVersion, "/")
	if len(gv) != 2 {
		return errors.New("apiVersion must use group/version form")
	}
	if s.Kind == "" {
		return errors.New("kind is required")
	}

	if s.Metadata.Name == "" {
		return errors.New("metadata.name is required")
	}

	if s.Spec.Kind == 0 || s.Spec.Tag == "!!null" {
		return nil
	}
	if s.Spec.Kind != yaml.MappingNode {
		return errors.New("spec must be a mapping")
	}
	// A yaml.Node bypasses the decoder's normal duplicate-key check.
	// Decode a temporary view to verify the complete raw spec
	// without interpreting fields owned by a particular kind.
	var fields map[string]any
	if err := s.Spec.Decode(&fields); err != nil {
		return fmt.Errorf("spec: %w", err)
	}

	return nil
}
