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

## Local development

Go 1.18 or later is required. From the repository root, run:

```powershell
cd app
go test ./...
go run .
```

Open `http://localhost:3000/`. The app uses in-memory storage unless Redis is configured, so local Guestbook entries are not durable and are lost when the process stops. The module dependencies are pinned in `app/go.mod` and `app/go.sum`.

## Optional container and Kubernetes configuration

`Dockerfile`, `deployment.yml`, `service.yml`, and `hpa.yml` are retained as assignment configuration. Docker and Kubernetes are not required for local development and were not used for the verified local run. To complete cloud-specific assignment tasks, an authenticated IBM Container Registry namespace and real Kubernetes cluster are still required; local execution is not a substitute for their command output.

`hpa.yml` has `minReplicas: 1` because Kubernetes HPA does not support a configured minimum of zero. Autoscaling, rollout history, and rollback remain unverified until run against a real cluster.

## Submission status

See [FINAL_REPORT.md](./submission/FINAL_REPORT.md) and [ASSIGNMENT_ANSWER.md](./submission/ASSIGNMENT_ANSWER.md). The cloud-dependent answers remain blocked until the authenticated IBM registry and Kubernetes environment is available.
