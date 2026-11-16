# IBM Cloud Kubernetes Guestbook — Assignment Answers

These answers reflect the current project sources and actual environment checks. Cloud command outputs that were not captured are explicitly marked BLOCKED; they must not be replaced with invented output.

## Question 1 — Updated Dockerfile

Actual project source (`Dockerfile`):

```dockerfile
FROM nginx:alpine

COPY index.html /usr/share/nginx/html/index.html

EXPOSE 80

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget -q -O /dev/null http://127.0.0.1/ || exit 1

CMD ["nginx", "-g", "daemon off;"]
```

Status: Source available. A Docker build and push were not run because Docker is not installed.

## Question 2 — IBM Cloud Container Registry image, tag v1

Command required: `ibmcloud cr images`

Status: **BLOCKED** — IBM Cloud CLI is not installed. No `crimages` output, registry namespace, or v1 tag was available to paste.

## Question 3 — Initial index.html / Guestbook - v1

Initial source snapshot preserved in `submission/evidence/task-03-default-index-html.txt`:

```html
<title>Guestbook - v1</title>
<h1>Guestbook - v1</h1>
```

Status: Source snapshot available. The current root `index.html` is the v2 application.

## Question 4 — Horizontal Pod Autoscaler details

Command required: `kubectl get hpa`

Status: **BLOCKED** — kubectl is not installed and no cluster output was available. No HPA table or replica count is claimed.

## Question 5 — Autoscaling / increased replicas

Commands required: `kubectl get hpa`, `kubectl get deployment`, and `kubectl get pods`.

Status: **BLOCKED** — kubectl and cluster access are unavailable. No CPU load test or increased replica count was observed.

## Question 6 — Docker push of updated image

Command pattern: `docker push us.icr.io/<actual-sn-labs-namespace>/guestbook:<updated-tag>`

Status: **BLOCKED** — Docker and IBM Cloud CLI are not installed; the authenticated registry namespace is unknown. No push output or SHA-256 digest exists to paste.

## Question 7 — Updated Kubernetes deployment applied

Command required: `kubectl apply -f deployment.yml`

Status: **BLOCKED** — kubectl and cluster access are unavailable. The image in `deployment.yml` still needs the actual IBM registry namespace. No apply output is claimed.

## Question 8 — Updated index.html / Guestbook - v2

Actual source in the root `index.html`:

```html
<title>Guestbook - v2</title>
<h1 id="page-title">Guestbook - v2</h1>
```

Status: Source is available. Local browser behavior was tested; Kubernetes serving was not.

## Question 9 — Deployment rollout history / CPU revision

Commands required: `kubectl rollout history deployment/guestbook` and `kubectl rollout history deployment/guestbook --revision=<actual-CPU-revision-number>`.

Status: **BLOCKED** — there is no accessible cluster Deployment or actual revision output. No CPU-related revision is claimed.

## Question 10 — ReplicaSets after deployment rollback

Command required: `kubectl get rs`

Status: **BLOCKED** — no cluster rollback was executed, so no ReplicaSet output is available.

## Submission status

Questions 1, 3, and 8 have source content. Questions 2, 4–7, and 9–10 require real IBM Cloud/Docker/Kubernetes execution evidence. This answer document intentionally contains no fabricated registry listing, digest, HPA count, rollout revision, or ReplicaSet result and is not a complete cloud-execution submission.
