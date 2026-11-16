# Guestbook Final Project — Execution Report

## Assignment overview

The GitHub repository was empty when inspected. A responsive, static Guestbook and Kubernetes deployment configuration have been added. Entries are stored in browser `localStorage`; they are not shared across browsers or pods.

## Project architecture

NGINX serves a single-page HTML/CSS/JavaScript application on port 80. The Dockerfile includes an HTTP health check. Kubernetes resources include a one-replica Deployment with readiness/liveness probes, resource requests/limits, rolling-update settings, a NodePort Service, and an autoscaling/v2 HPA.

## Environment and execution

- The working environment is Windows PowerShell.
- Git and GitHub CLI are installed; Git identity is configured.
- The target GitHub repository was confirmed public and empty; GitHub CLI authentication was active.
- Docker, IBM Cloud CLI, and kubectl were not installed (`Get-Command` reported each as `NOT INSTALLED`).
- The local HTML page rendered in the browser; a submitted message was shown and remained after page reload. This verifies browser-side behavior only, not a container or cluster deployment.
- Consequently there was no Docker daemon check, IBM Cloud target/registry authentication, active Kubernetes context, cluster connectivity, or live cluster output to report.
- No image build/push, registry digest, manifest apply, autoscaling, rollout, or rollback is claimed.

## Docker image details

`Dockerfile` builds on `nginx:alpine`, copies `index.html` into the NGINX document root, exposes port 80, and serves in the foreground. Image tags v1 and v2 have not been built or pushed. The registry hostname, account namespace, and image reference remain unknown; `deployment.yml` therefore contains a clearly marked placeholder.

## Kubernetes deployment details

The Deployment uses label-based selectors, HTTP probes, CPU request 50m / limit 250m, memory request 32Mi / limit 128Mi, and a rolling update strategy. The Service is NodePort. Manifests have not been submitted to a cluster.

## Autoscaling results

`hpa.yml` uses CPU utilization at 60%, minimum 1 replica, and maximum 4 replicas. This is a valid starting configuration, not a capacity-verified course-cluster limit. The HPA and metrics server were not available to test. No zero-replica HPA observation or CPU-triggered replica increase is recorded. Kubernetes HPA does not allow a configured minimum of zero; a possible transient live zero-replica reading is not guaranteed and must not be presented without cluster evidence.

## Version 2, rolling update, and rollback

The current `index.html` has the requested exact `<title>Guestbook – v2</title>` and an H1 of `Guestbook – v2`, using U+2013 EN DASH. The v2 image, rollout, CPU-related revision, and genuine rollback were not run because Docker, IBM Cloud access, and Kubernetes access are unavailable.

## Ten-task grading checklist

| Task | Result | Evidence |
|---|---|---|
| 01 Dockerfile | Source complete; image build blocked | `submission/evidence/task-01-dockerfile.txt` |
| 02 IBM CR images v1 | Blocked: IBM CLI/auth unavailable | `submission/evidence/task-02-ibm-cr-images-v1.txt` |
| 03 Default HTML | Initial source snapshot complete | `submission/evidence/task-03-default-index-html.txt` |
| 04 HPA zero replicas | Blocked: no cluster; zero minimum invalid | `submission/evidence/task-04-hpa-zero-replicas.txt` |
| 05 Autoscaling replicas | Blocked: no cluster/load observation | `submission/evidence/task-05-autoscaling-replicas.txt` |
| 06 Docker push digest | Blocked: no Docker/IBM CLI | `submission/evidence/task-06-docker-push-digest.txt` |
| 07 Deployment apply | Blocked: no kubectl/cluster; image placeholder | `submission/evidence/task-07-deployment-apply.txt` |
| 08 Guestbook v2 HTML | Source complete; live verification blocked | `submission/evidence/task-08-guestbook-v2-html.txt` |
| 09 Rollout history | Blocked: no cluster/revision | `submission/evidence/task-09-rollout-history.txt` |
| 10 ReplicaSets after rollback | Blocked: no cluster/rollback | `submission/evidence/task-10-replicasets-after-rollback.txt` |

Evidence files are under `submission/evidence/`. Each records what is available and marks unavailable execution as blocked; no output, digest, timestamp, or replica count has been invented.

## Outstanding blockers and next steps

Install/start Docker Desktop, install/authenticate IBM Cloud CLI with the Container Registry plugin, select the intended IBM Cloud account/region, create or select a registry namespace, and obtain access to a running Kubernetes cluster with metrics-server. Then build and push v1/v2, replace the image placeholder, apply the manifests, capture actual command output, generate CPU load, and execute/verify the revision and rollback sequence. Do not commit credentials or local kubeconfig.

## Score readiness

**Verified score readiness: not determinable from the supplied information.** The point allocation across the ten tasks was not present in the empty repository, so assigning a numeric score out of 15 would be speculative. Source/configuration is prepared for tasks 01, 03, and 08, but no cloud-dependent marks are verified; the 70% passing threshold is not currently evidenced.
