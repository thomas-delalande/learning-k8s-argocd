This is a minimal project to configure Kubernetes and ArgoCD.

## Prerequisites

Running this as is requires the following dependencies:

- `argocd`
- `kubectl`
- `docker`
- `minikube`
- `go`

## Usage

```bash
./scripts/setup.sh

# To access the ArgoCD UI on localhost:8080
./scripts/port-forward-argocd-ui.sh

# Once synchronised
./scripts/port-forward-server.sh

# You can now access the server on http://localhost:8081/
# and the secret value on http://localhost:8081/secret
# and the database on http://localhost:8081/postgres?value=my-custom-value

# To remove resources and stop minikube
./scripts/teardown.sh
```

## Result

### ArgoCD UI

![ArgoCD UI](./docs/image-1-argo-cd-ui.png)

### Server Response

![Server Response](./docs/image-2-server-response.png)

