// Package laniscope implements Laniakea's behavior-independent source model.
//
// Each laniscope.yaml file is one complete Kubernetes-shaped source
// document.
// Kind packages own typed specs, backend packages own provider
// configuration, and an explicit Scheme connects them.
// Directory containment is represented by a separate Tree;
// neither backends nor children are recursive Laniscopes inside a
// document's spec.
//
package laniscope
