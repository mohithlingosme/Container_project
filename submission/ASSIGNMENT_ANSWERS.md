# Assignment Answer Status

This file is a transparent status record, not a completed answer sheet. Missing IBM Cloud and Kubernetes command output is not fabricated. The requested `answer.txt` cannot be completed until the blocked commands are run in the authenticated course environment.

## Question 1 — Dockerfile

The current Dockerfile source is captured in `submission/evidence/Dockerfile`. It packages the official Go Guestbook on port 3000. The Docker image build is **BLOCKED** because Docker and its engine are unavailable.

## Question 2 — IBM Container Registry images

**BLOCKED.** `ibmcloud cr images` was not run: IBM Cloud CLI, registry authentication, and the actual namespace are unavailable. No registry output is claimed.

## Question 3 — Initial Guestbook v1 HTML

The complete source snapshot with `<title>Guestbook - v1</title>` and `<h1>Guestbook - v1</h1>` is `submission/evidence/app`.

## Question 4 — HPA output

**BLOCKED.** No `kubectl get hpa` output is available. `hpa.yml` configures CPU utilization at 60%, with 1–4 replicas; this manifest is not evidence of a live HPA.

## Question 5 — Autoscaling

**BLOCKED.** No cluster or metrics were available to generate load or observe replicas. No scaling result is claimed.

## Question 6 — Updated image push

**BLOCKED.** Docker, IBM Cloud CLI, registry authentication, and the namespace are unavailable. No push output or digest exists.

## Question 7 — Updated Deployment apply

**BLOCKED.** kubectl, cluster access, and a real registry image reference are unavailable. The image field in `deployment.yml` remains an explicit placeholder.

## Question 8 — Updated Guestbook v2 HTML

The complete source with `<title>Guestbook - v2</title>` and `<h1>Guestbook - v2</h1>` is `submission/evidence/up-app`. The official Go application was compiled and locally smoke-tested; Kubernetes serving is unverified.

## Question 9 — CPU revision history

**BLOCKED.** No Deployment or rollout history is available. No revision number is claimed.

## Question 10 — ReplicaSets after rollback

**BLOCKED.** No rollback was executed and no `kubectl get rs` output is available. No ReplicaSet names, counts, or rollback result are claimed.
