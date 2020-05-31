# Understanding Kubernetes source code

## Overview
* This is a doc to explain the structure of Kubernetes source code.
* To better view Kubernetes source code, clone the code to local directory, and use VS Code + Go extension + Go Outliner extension. Go Outliner extension shows all types, functions, variables and constants in a package. This is especially useful because Go code in a packaged can be scattered in many files and are not ordered by types. 

## /api/openapi-spec
* This directory contains /swagger.json, which is the spec of all Kubernetes API.
* This spec is used to generate or define the most fundamental data structures (struct) in Go language source code.

## /staging/src/k8s.io
* This is where the source codes are located.
* Children directories are synced to other Git repos labeled with "k8s-staging" (e.g. https://github.com/kubernetes/kubelet). Those repos are then built and imported back as dependent packages (located under /pkg). This seems to be necessary to better manage the dependency tree, and make vendoring easier.

### ./apimachinery
* This is a library that provide meta-structures and utilities for Kubernetes API types.
* For example, it defines how versioning works (./pkg/version), what labels are (./pkg/labels).
* It is very low on the dependency tree.

### ./api
* This directory contains API definitions by API groups. Each child directory contains API definitions of a group. Most of API structs are in /types.go file in each group's root directory.
* More on [Paths for API and API groups](https://www.oreilly.com/library/view/managing-kubernetes/9781492033905/ch04.html).
* ./core: The most essential API object types, such as Container, Pod, Volumn, ConfigMap.
* ./app: Higher-level API object types based on core group, such as Deployment, ReplicaSet.
