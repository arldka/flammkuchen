<div align="center">

# Flammkuchen

[![Go Report Card](https://goreportcard.com/badge/github.com/arldka/flammkuchen)](https://goreportcard.com/report/github.com/arldka/flammkuchen)
![Latest Release](https://img.shields.io/github/v/release/arldka/flammkuchen?include_prereleases)
![GitHub License](https://img.shields.io/github/license/arldka/flammkuchen)
![Docker Image Size (tag)](https://img.shields.io/docker/image-size/arldka/flammkuchen/latest)

Flammkuchen is an open source FluxCD UI that aims to replace the main use-cases that were previously covered [Flamingo](https://github.com/flux-subsystem-argo/flamingo) which is now inactive / no longer maintained.
The project aims to be simple enough with a limited set of core-features as to be easily maintainable / reusable for other usecases.

The main features are:
- HelmRelease & Kustomization searchable overview with current status that *should support most recent versions of FluxCD (no matter the version of CRDs installed in your cluster.)*
- Detailed view of HelmReleases and Kustomization showing the deployed objects with their statuses (including any custom CRDs)
- **IN-PROGESS** ability to link back to GKE / Datadog service catalog to have a quick way to access more detailed statuses or logs from another platform

</div>

> [!WARNING]
> The active development focusses on maintaining functionality and keeping CVE-free dependencies and is not currently tested
> While the UI aims to be read-only (for security purposes), it cannot be considered at this stage as safe and therefore should not be used in production or exposed publicly.

# How to deploy ?

## In-Cluster

1. Create a Deployment: Use the ghcr.io/arldka/flammkuchen Docker image and expose containerPort: 8080.
2. Service Account: Use a dedicated Kubernetes service account for cluster role bindings.
3. Cluster Role Binding: Ensure the service account has a clusterrolebinding on the view clusterrole to list all objects.
4. Read Access to Secrets: Since the view cluster role doesn't provide read access to secrets, Flammkuchen requires additional permissions:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: read-secrets-clusterrole
rules:
- apiGroups: [""]
  resources: ["secrets"]
  verbs: ["get", "list"]
```

# Development

## Repository Structure

```bash
./
|- cmd
|   |- server # net/http and k8sclient init, serves the handlers
|- components
|   |- *.templ # templ files that will generate HTML using go
|   |- *.go    # Generated go files by templ (DO NOT MANUALLY MODIFY)
|   |- objects #  Customizable templ objects to display custom info for different kinds of Kubernetes resources
|- handlers # stores the net/http handlers. use the services to retrieve data and internal functions
|- internal # Miscellaneous small reusable functions to be used in services, handlers and components
|- services # Inits the k8sclient and is generally the getter of Kubernetes objects
|- Dockerfile #  multi-stage build Dockerfile that can be used for local builds
|- Dockerfile.release # Dockerfile used by goreleaser
```

## Used Tools / Packages
* **Library**: [htmx](https://htmx.org/) used via CDN
* **CSS Framework**: [daisyUI](https://daisyui.com/)
* **Builds HTML in Go**: uses the [templ](https://github.com/a-h/templ) package
* **HTTP Server**: Built-in [net/http](https://pkg.go.dev/net/http) go package
* **Release Management** [GoReleaser](https://goreleaser.com/) and signed by [cosign](https://github.com/sigstore/cosign?tab=readme-ov-file)

## How to run locally

With a-h/templ installed in your environment:
```bash
 VERSION=$(git rev-parse HEAD) templ generate --watch --proxy="http://localhost:8080" --cmd="go run cmd/server/main.go"
```

This runs the app and authenticates with your current kubernetes context for testing. 
The app will live reload when go / templ files are modified. 

---

## Disclaimer

Flammkuchen is not an official product and is provided as-is to the community.
