# Istio External Authorization Http Example

## Pre Requisites

* Kubernetes Cluster
* Istio, Istio Ingress Gateway

## Build the ext-authz example image

```bash
$ docker build ./ -t ext-authz:latest 
```

## Run on Kubernetes

```bash
$ kubectl create ns ext-authz
$ kubectl label ns ext-authz istio-injection=enabled
$ kubectl apply -f ./deployments
```

## See also

[Istio External Authorization](https://istio.io/latest/docs/tasks/security/authorization/authz-custom/)
