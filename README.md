# Guestbook – IBM Cloud Kubernetes project

A lightweight, responsive static Guestbook served by NGINX. Visitors can add and view messages; this assignment-sized implementation stores entries in the visitor's browser `localStorage` (not in a shared server-side database).

## Files

- `index.html` — Guestbook UI (currently v2).
- `Dockerfile` — NGINX image, port 80, and an HTTP health check.
- `deployment.yml` — single-replica Deployment with rolling updates, probes, and CPU/memory resources.
- `service.yml` — NodePort Service for cluster access.
- `hpa.yml` — CPU autoscaling from 1 to 4 replicas; requires a functioning metrics server.
- `submission/evidence/` — task-by-task evidence and explicit blockers.

## Configure the IBM Container Registry image

The deployment image is intentionally a placeholder because the authenticated IBM Cloud registry region and namespace must be discovered in the target account. After IBM Cloud access is available:

1. Authenticate with the IBM Cloud CLI (`ibmcloud login`) and select the intended account, region, and resource group.
2. Discover registry regions/namespaces with `ibmcloud cr region` and `ibmcloud cr namespace-list`; create a namespace if needed.
3. Use the actual registry hostname/region and namespace for `IMAGE`, for example `us.icr.io/<namespace>/guestbook`.
4. Replace the placeholder in `deployment.yml` with `${IMAGE}:v2`'s concrete value. Kubernetes manifests do not expand shell variables in image fields.

Example PowerShell commands (replace placeholders with discovered values):

```powershell
$Registry = "us.icr.io"
$Namespace = "<your-namespace>"
$Image = "$Registry/$Namespace/guestbook"
docker build -t "${Image}:v1" .
ibmcloud cr login
docker push "${Image}:v1"
docker build -t "${Image}:v2" .
docker push "${Image}:v2"
```

`ibmcloud cr login` uses the CLI's authenticated workflow. Do not put an API key or access token in source files or command arguments. Set the Deployment image to the exact pushed v2 reference before applying.

## Deploy

Ensure Docker, IBM Cloud CLI plus its Container Registry plugin, kubectl, an authenticated IBM Cloud account, a connected Kubernetes cluster, and a metrics server are available first.

```powershell
kubectl apply -f deployment.yml
kubectl apply -f service.yml
kubectl apply -f hpa.yml
kubectl rollout status deployment/guestbook
kubectl get deployments,pods,services,hpa
```

The NodePort Service is suitable for a course cluster; use the assigned IBM Cloud ingress/load-balancer arrangement if the course cluster requires a different exposure type.

## Demonstrate autoscaling and rollback

HPA requires `minReplicas >= 1`; zero is not a valid configured minimum. A live `kubectl get hpa` observation can show current replicas at zero only if the Deployment is momentarily scaled to zero before the HPA controller reconciles, and it is not guaranteed. Do not claim a zero-replica HPA or autoscaling demonstration without captured cluster output.

To demonstrate CPU scaling, generate load against the running service using a controlled load tool available in the cluster and observe `kubectl get hpa --watch`, pods, and Deployment replicas. The HPA maximum of four is a conservative starting point; adjust it only after confirming cluster capacity and assignment instructions.

For rollout history, make and document an intentional Deployment CPU-resource change, apply it, capture `kubectl rollout history deployment/guestbook`, and then perform a genuine `kubectl rollout undo deployment/guestbook`. Verify the restored ReplicaSets and pods before recording evidence.

## Current execution constraints

This checkout was empty at project start. The current execution environment has Git and GitHub CLI, but Docker, IBM Cloud CLI, and kubectl are not installed. No IBM authentication, registry namespace, cluster, image digest, deployment, autoscaling observation, or rollback is claimed. See `submission/FINAL_REPORT.md` and the evidence files for the exact blockers.
