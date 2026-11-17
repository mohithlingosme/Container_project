# IBM Kubernetes Guestbook

This project uses the official IBM Skills Network Guestbook v1 Go application, with the required v1/v2 page title changes. The Go HTTP server listens on port 3000 and serves the frontend from `app/public/`. It stores entries in memory unless a Redis service is configured, so the default setup does not provide durable or shared storage across pods.

The upstream application is licensed under Apache-2.0; see [LICENSE](./LICENSE). Its source is from [ibm-developer-skills-network/guestbook](https://github.com/ibm-developer-skills-network/guestbook).

## Files

- `app/main.go` and `app/public/` — Guestbook server and frontend.
- `Dockerfile` — two-stage Go build for the assignment's port 3000 application.
- `deployment.yml` — Kubernetes Deployment; the IBM Container Registry namespace is intentionally a placeholder until discovered from the authenticated account.
- `service.yml` — NodePort Service targeting port 3000.
- `hpa.yml` — CPU autoscaling configuration (1–4 replicas, 60% target).
- `submission/evidence/` — source snapshots and actual execution blockers/results.

## Run locally

From `app/`, initialize the Go module and start the application:

```powershell
go mod init guestbook
go mod tidy
go run main.go
```

Open `http://localhost:3000/`. The project was locally compiled with Go 1.26.4 and smoke-tested on port 3000. The Dockerfile pins its Go dependencies to versions compatible with the specified Go 1.18 builder; the container build has not been run because Docker is unavailable.

## Build and deploy

The registry host, namespace, IBM Cloud authentication, and Kubernetes cluster must be obtained from the intended course environment. Do not guess or commit credentials. Once those are available, build and push the v1/v2 image tags, replace the image placeholder in `deployment.yml`, then apply the Deployment, Service, and HPA to the authenticated cluster. Capture command output from that actual environment; local tests are not a substitute for cluster evidence.

`hpa.yml` has `minReplicas: 1` because Kubernetes HPA does not support a configured minimum of zero. Autoscaling, rollout history, and rollback remain unverified until run against a real cluster.

## Submission status

See [FINAL_REPORT.md](./submission/FINAL_REPORT.md) and [ASSIGNMENT_ANSWER.md](./submission/ASSIGNMENT_ANSWER.md). The cloud-dependent answers remain blocked until the authenticated IBM registry and Kubernetes environment is available.
