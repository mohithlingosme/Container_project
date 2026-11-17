# IBM Kubernetes Guestbook — Assignment Answers

These answers use the local Go development workflow and actual local test results. This environment has no Docker engine, IBM Cloud CLI, `kubectl`, kubeconfig, or connected course cluster. IBM registry and Kubernetes outputs cannot be produced by local development, so cloud-specific questions are explicitly marked **NOT RUN LOCALLY** rather than fabricated. This local work does not substitute for cloud-execution marks in the course rubric.

## Question 1 — Dockerfile

```dockerfile
FROM golang:1.18 AS builder

WORKDIR /app

COPY app/go.mod app/go.sum ./
RUN go mod download

COPY app/main.go .
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

`app/go.mod` and `app/go.sum` pin the Go dependencies. The app builds and tests locally with Go 1.26.4; Docker image build was **not run** because Docker and its engine are unavailable. Source copy: [`evidence/Dockerfile`](./evidence/Dockerfile).

## Question 2 — IBM Container Registry images (v1)

**NOT RUN LOCALLY — actual `ibmcloud cr images` output is unavailable.** IBM Cloud CLI/plugin, authenticated registry access, and the account's namespace are not present in this environment. Local development does not create an IBM registry image or listing. No image listing or v1 tag is claimed. See [`evidence/task-02-ibm-cr-images-v1.txt`](./evidence/task-02-ibm-cr-images-v1.txt).

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

**NOT RUN LOCALLY — actual `kubectl get hpa` table is unavailable.** The checked-in [`hpa.yml`](../hpa.yml) configures the `guestbook` Deployment target, CPU utilization target 60%, minimum 1 replica, and maximum 4 replicas. This manifest is configuration, not live HPA output. Local Go execution does not run an HPA. A zero minimum is not configured because Kubernetes HPA does not permit it. See [`evidence/task-04-hpa-zero-replicas.txt`](./evidence/task-04-hpa-zero-replicas.txt).

## Question 5 — Autoscaling results

**NOT RUN LOCALLY — no cluster load test or replica increase was observed.** The cluster, metrics server, and `kubectl` are unavailable. No HPA metrics, replica counts, or successful autoscaling result can be supplied. The local Go app/API test is not a Kubernetes autoscaling test. See [`evidence/task-05-autoscaling-replicas.txt`](./evidence/task-05-autoscaling-replicas.txt).

## Question 6 — Updated image push and digest

**NOT RUN LOCALLY — no Docker push output or image digest exists.** Docker, IBM Cloud CLI, registry authentication, and the actual registry namespace are unavailable. Local execution runs the Go process directly and does not create or push an image. See [`evidence/task-06-docker-push-digest.txt`](./evidence/task-06-docker-push-digest.txt).

## Question 7 — Updated Deployment apply

**NOT RUN LOCALLY — no `kubectl apply` output exists.** `kubectl` and cluster access are unavailable. [`deployment.yml`](../deployment.yml) still has the explicit image placeholder `us.icr.io/REPLACE_WITH_NAMESPACE/guestbook:v2`; no registry namespace was guessed. For local development, run `cd app`, `go test ./...`, then `go run .` and open `http://localhost:3000/`. See the local results under Question 8 and [`README.md`](../README.md).

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

This source is preserved at [`evidence/up-app`](./evidence/up-app). Local verification with Go 1.26.4:

```text
go test ./...: PASS
go vet ./...: PASS
GET /: HTTP 200; title and heading are Guestbook - v2
GET /hello: HTTP 200
GET /rpush/guestbook/local-dev-test: ["local-dev-test"]
GET /lrange/guestbook: ["local-dev-test"]
```

This verifies the local application only, not a container or cluster deployment. Detailed results: [`evidence/task-08-guestbook-v2-html.txt`](./evidence/task-08-guestbook-v2-html.txt).

## Question 9 — Rollout history and CPU revision

**NOT RUN LOCALLY — no Deployment rollout history or actual revision is available.** `kubectl` and a cluster context are absent; no CPU revision was created or inspected. The local Deployment manifest's CPU request/limit values are not rollout evidence. See [`evidence/task-09-rollout-history.txt`](./evidence/task-09-rollout-history.txt).

## Question 10 — ReplicaSets after rollback

**NOT RUN LOCALLY — no rollback was executed and no `kubectl get rs` output is available.** No ReplicaSet names, counts, or rollback result are claimed. See [`evidence/task-10-replicasets-after-rollback.txt`](./evidence/task-10-replicasets-after-rollback.txt).

## Completion status

Run the app locally with Go using the commands in [`README.md`](../README.md). Local source answers and tests are verified, but this is **not a complete cloud-execution submission**: Questions 2, 4–7, and 9–10 require actual output from IBM Container Registry, Docker, and a Kubernetes cluster.
