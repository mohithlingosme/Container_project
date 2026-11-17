# Guestbook Final Project — Execution Report

## Implementation completed locally

- Replaced the NGINX-only container target with the official IBM Skills Network Guestbook Go server and its public assets under `app/`.
- Updated the page title and heading to Guestbook v2; preserved the full v1 HTML snapshot at `submission/evidence/app` and the v2 snapshot at `submission/evidence/up-app`.
- Updated the Deployment and Service to use the app's HTTP port 3000 and a consistent `app: guestbook` selector. The Deployment retains the actual registry namespace as a placeholder rather than guessing it.
- Retained an HPA manifest with CPU target 60%, minimum 1, and maximum 4 replicas.
- Pinned Go module versions in the Docker build, including `gorilla/mux` 1.8.0, because 1.8.1 requires Go 1.20 and is incompatible with the required Go 1.18 builder.
- Added the upstream Apache-2.0 license.

## Local verification

- Built the Go application with the installed Go 1.26.4 toolchain in a temporary build directory.
- Ran the app and verified HTTP 200 on `/` and `/hello`; submitted a smoke-test entry through `/rpush/guestbook/smoke-test` and read it back through `/lrange/guestbook`.
- Confirmed the source and manifests describe port 3000. The Dockerfile's Go 1.18 image build was not executed.

These are local source/runtime checks only. They are not Docker, registry, or Kubernetes evidence.

## External execution blockers

- Docker CLI/engine: not installed.
- IBM Cloud CLI and Container Registry plugin: not installed; no authenticated registry namespace was available.
- kubectl: not installed; no kubeconfig or current cluster context exists.
- Consequently, image build/push, digest, manifest apply, HPA observation, CPU-driven scaling, rollout revision, rollback, and ReplicaSet output remain unavailable.

## 15-mark readiness

| Task | Status | Evidence / limitation |
|---|---|---|
| 01 Dockerfile | INCOMPLETE | Source is captured in `evidence/Dockerfile`; Docker image build not run. |
| 02 Registry v1 image | BLOCKED | No IBM CLI, registry authentication, namespace, or `cr images` output. |
| 03 Initial v1 HTML | VERIFIED | Complete v1 source in `evidence/app`. |
| 04 HPA output | BLOCKED | No cluster or live `kubectl get hpa` output. |
| 05 Autoscaling | BLOCKED | No cluster, metrics, or observed replica increase. |
| 06 Updated image push | BLOCKED | No Docker/IBM CLI, registry access, push output, or digest. |
| 07 Deployment apply | BLOCKED | No kubectl/cluster or real image reference. |
| 08 Updated v2 HTML | VERIFIED | Complete v2 source in `evidence/up-app`; local Go smoke test passed. |
| 09 CPU revision history | BLOCKED | No Deployment or real rollout history. |
| 10 Rollback ReplicaSets | BLOCKED | No rollback or `kubectl get rs` output. |

The point allocation across tasks was not provided, so a numeric score out of 15 cannot be calculated reliably. A completed `answer.txt` is intentionally not generated from explanatory substitutes for the missing command outputs.

## GitHub status

The work is on branch `feature/guestbook-final-project`. Commit and push status is reported separately after repository operations complete.
