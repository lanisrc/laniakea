---
title: "Design"
date: 2023-05-19T23:20:03-07:00
draft: true
---

This document describes the core ideas behind `Laniakea`.

## The `Laniscope` Tree

At its heart, `Laniakea` is built around a recursive directory tree.
Each directory defines a `Laniscope`.

```text
<laniscope>/
  laniscope.yaml
  <child-laniscope>/
    laniscope.yaml
```

A `Laniscope` is an infrastructure boundary described by a
"`laniscope.yaml`" file.
The concept is intentionally flexible.
It is not tied to a particular organizational or technical layer, such
as a tenant, team, or cluster.

Depending on the deployment, a `Laniscope` may represent:

* a service mesh or federation of clusters
* a Kubernetes cluster managed by `Laniakea`
* a tenant in a multi-tenant platform
* a hub, lab, collaboration, project, instrument, or workspace
* a delegated repository
* another infrastructure boundary that may be introduced in the future

The key idea is simple:
`Laniakea` models infrastructure as a hierarchy of `Laniscope`s.
Rather than baking assumptions about organizational structure into the
system, each scope can opt into the capabilities and policies it
needs.
This allows the same abstraction to operate at many different scales,
from a single workspace to an entire federation of clusters.

## GitOps Native

`Laniakea` is designed to be GitOps native.

Administrators and users work directly with the source `Laniscope`
tree.
`Laniakea` validates that tree and renders it into standard GitOps
resources.

```text
source Laniscope tree -> validate -> render -> FluxCD
```

The rendered output remains ordinary YAML, Kustomize configurations,
and FluxCD resources.
Users should always be able to inspect the generated artifacts and
understand exactly what will be applied to a cluster.

`Laniakea` exists to reduce the amount of boilerplate that people need
to write and maintain.
It should make common infrastructure patterns easier to express
without hiding the underlying Kubernetes and GitOps state.

The source `Laniscope` tree is the primary interface presented to
users.
The rendered manifests are an implementation detail, but they remain
fully visible and editable.
`Laniakea` does not replace GitOps tools such as FluxCD;
it provides a structured way to organize and generate the resources
that those tools manage.
