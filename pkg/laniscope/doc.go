// Package laniscope loads and validates Laniakea source declarations.
//
// A Laniscope is a generic infrastructure boundary described by a
// laniscope.yaml file:
//
//	apiVersion: laniscope/v1alpha1
//	kind: Kubernetes
//	metadata: # uniquely identifies and organizes the object
//	  name: local
//	spec: # defines desired behavior and configuration
//	  backend:
//	    kind: k3d
//	    spec:
//	      image: rancher/k3s:v1.27.1+k3s1
//	  features:
//	    - kind: GitOps
//	      metadata:
//	        name: laniops
//	      spec:
//	        backend:
//	          kind: flux2
//
// The kind owns the user-facing spec contract.
// Backend kind selects the implementation backend, and backend.spec
// carries backend-only settings.
// Embedded features are full Laniscope objects under spec.features.
//
// YAML kind and backend names are normalized to lowercase internal IDs when
// resolving built-in capabilities.
// For example, Kubernetes plus k3d resolves to kubernetes/k3d.
//
package laniscope
