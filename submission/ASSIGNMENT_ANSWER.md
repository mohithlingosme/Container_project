# IBM Kubernetes Guestbook — Assignment Answers

These answers use checked-in project source and actual local test results. Cloud command output is included only when it was actually captured. This environment has no Docker engine, IBM Cloud CLI, `kubectl`, kubeconfig, or connected course cluster; cloud-dependent answers are therefore explicitly marked **BLOCKED** rather than fabricated.

## Question 1 — Dockerfile

```dockerfile
FROM golang:1.18 AS builder

WORKDIR /app

COPY app/main.go .

RUN go mod init guestbook
RUN go mod edit '-require=github.com/codegangsta/negroni@v1.0.0' '-require=github.com/gorilla/mux@v1.8.0' '-require=github.com/xyproto/simpleredis/v2@v2.9.0'
RUN go mod tidy
RUN go build -o main main.go

FROM ubuntu:18.04

COPY --from=builder /app/main /app/guestbook
COPY app/public/index.html /app/public/index.html
COPY app/public/script.js /app/public/script.js
COPY app/public/style.css /app/public/style.css
COPY app/public/jquery.min.js /app/public/jquery.min.js

WORKDIR /app

CMD ["./guestbook"]

EXPOSE 3000
```

The Go app compiled locally with Go 1.26.4 and passed a local HTTP/API smoke test. A Docker image build was **not run** because Docker and its engine are unavailable. Source copy: [`evidence/Dockerfile`](./evidence/Dockerfile).

## Question 2 — IBM Container Registry images (v1)

**BLOCKED — actual `ibmcloud cr images` output is unavailable.** IBM Cloud CLI/plugin, authenticated registry access, and the account's namespace are not present in this environment. No image listing or v1 tag is claimed. See [`evidence/task-02-ibm-cr-images-v1.txt`](./evidence/task-02-ibm-cr-images-v1.txt).

## Question 3 — Initial Guestbook HTML (v1)

Complete source snapshot:

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta content="text/html; charset=utf-8" http-equiv="Content-Type">
    <meta charset="utf-8">
    <meta content="width=device-width" name="viewport">
    <link href="style.css" rel="stylesheet">
    <title>Guestbook - v1</title>
  </head>
  <body>
    <div id="header">
      <h1>Guestbook - v1</h1>
    </div>

    <div id="guestbook-entries">
      <link href="https://afeld.github.io/emoji-css/emoji.css" rel="stylesheet">
      <p>Waiting for database connection... <i class='em em-boat'></i></p>

    </div>

    <div>
      <form id="guestbook-form">
        <input autocomplete="off" id="guestbook-entry-content" type="text">
        <a href="#" id="guestbook-submit">Submit</a>
      </form>
    </div>

    <div>
      <p><h2 id="guestbook-host-address"></h2></p>
      <p><a href="env">/env</a>
      <a href="info">/info</a></p>
    </div>
    <script src="jquery.min.js"></script>
    <script src="script.js"></script>
  </body>
</html>
```

This is the preserved v1 source snapshot at [`evidence/app`](./evidence/app), not proof that a v1 image was deployed.

## Question 4 — Horizontal Pod Autoscaler

**BLOCKED — actual `kubectl get hpa` table is unavailable.** The checked-in [`hpa.yml`](../hpa.yml) configures the `guestbook` Deployment target, CPU utilization target 60%, minimum 1 replica, and maximum 4 replicas. This manifest is configuration, not live HPA output. A zero minimum is not configured because Kubernetes HPA does not permit it. See [`evidence/task-04-hpa-zero-replicas.txt`](./evidence/task-04-hpa-zero-replicas.txt).

## Question 5 — Autoscaling results

**BLOCKED — no cluster load test or replica increase was observed.** The cluster, metrics server, and `kubectl` are unavailable. No HPA metrics, replica counts, or successful autoscaling result can be supplied. The local Go app/API smoke test is not a Kubernetes autoscaling test. See [`evidence/task-05-autoscaling-replicas.txt`](./evidence/task-05-autoscaling-replicas.txt).

## Question 6 — Updated image push and digest

**BLOCKED — no Docker push output or image digest exists.** Docker, IBM Cloud CLI, registry authentication, and the actual registry namespace are unavailable. See [`evidence/task-06-docker-push-digest.txt`](./evidence/task-06-docker-push-digest.txt).

## Question 7 — Updated Deployment apply

**BLOCKED — no `kubectl apply` output exists.** `kubectl` and cluster access are unavailable. [`deployment.yml`](../deployment.yml) still has the explicit image placeholder `us.icr.io/REPLACE_WITH_NAMESPACE/guestbook:v2`; no registry namespace was guessed. See [`evidence/task-07-deployment-apply.txt`](./evidence/task-07-deployment-apply.txt).

## Question 8 — Updated Guestbook HTML (v2)

Complete source snapshot:

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta content="text/html; charset=utf-8" http-equiv="Content-Type">
    <meta charset="utf-8">
    <meta content="width=device-width" name="viewport">
    <link href="style.css" rel="stylesheet">
    <title>Guestbook - v2</title>
  </head>
  <body>
    <div id="header">
      <h1>Guestbook - v2</h1>
    </div>

    <div id="guestbook-entries">
      <link href="https://afeld.github.io/emoji-css/emoji.css" rel="stylesheet">
      <p>Waiting for database connection... <i class='em em-boat'></i></p>

    </div>

    <div>
      <form id="guestbook-form">
        <input autocomplete="off" id="guestbook-entry-content" type="text">
        <a href="#" id="guestbook-submit">Submit</a>
      </form>
    </div>

    <div>
      <p><h2 id="guestbook-host-address"></h2></p>
      <p><a href="env">/env</a>
      <a href="info">/info</a></p>
    </div>
    <script src="jquery.min.js"></script>
    <script src="script.js"></script>
  </body>
</html>
```

This source is preserved at [`evidence/up-app`](./evidence/up-app). The locally compiled app returned HTTP 200 on `/` and `/hello`; a message submitted to `/rpush/guestbook/smoke-test` was returned by `/lrange/guestbook`. This verifies local behavior only, not a container or cluster deployment. Full local results: [`evidence/task-08-guestbook-v2-html.txt`](./evidence/task-08-guestbook-v2-html.txt).

## Question 9 — Rollout history and CPU revision

**BLOCKED — no Deployment rollout history or actual revision is available.** `kubectl` and a cluster context are absent; no CPU revision was created or inspected. The local Deployment manifest's CPU request/limit values are not rollout evidence. See [`evidence/task-09-rollout-history.txt`](./evidence/task-09-rollout-history.txt).

## Question 10 — ReplicaSets after rollback

**BLOCKED — no rollback was executed and no `kubectl get rs` output is available.** No ReplicaSet names, counts, or rollback result are claimed. See [`evidence/task-10-replicasets-after-rollback.txt`](./evidence/task-10-replicasets-after-rollback.txt).

## Completion status

This document provides the verified source answers for Questions 1, 3, and 8, plus the available local verification. It is **not a complete cloud-execution submission**: Questions 2, 4–7, and 9–10 require real output from the authenticated IBM Cloud registry and Kubernetes course environment.
