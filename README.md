# API Status Operator

## Overview

This is a simple example of deploying rest api using operator in kubernetes.

## Quick Start

This quick start guide walks through the process of building the api-status-operator.

### Prerequisites


- [dep](https://golang.github.io/dep/docs/installation.html) version v0.5.0+.
- [go](https://golang.org/dl/) version v1.10+
- [docker](https://docs.docker.com/install/) version 17.03+
- Minikube v1.0.0+ cluster

### Install the Operator SDK CLI

First, checkout and install the operator-sdk CLI:

```
$ cd $GOPATH/src/github.com/operator-framework/operator-sdk
$ git checkout master // currently, there are no releases that include the test framework, so use the master for now
$ dep ensure
$ make install github.com/operator-framework/operator-sdk/commands/operator-sdk
```

Checkout this API Status Operator repository

```
$ mkdir $GOPATH/src/github.com/pratikjagrut
$ cd $GOPATH/src/github.com/pratikjagrut
$ git clone https://github.com/pratikjagrut/api-status-operator.git
$ cd api-status-operator
```

Vendor the dependencies

```
$ dep ensure
```

### Deploying this operator localy

```
$ kubectl apply -f deploy/crds/apistatus_v1alpha1_apistatus_crd.yaml
$ kubectl apply -f deploy/crds/apistatus_v1alpha1_apistatus_cr.yaml
$ operator-sdk up local
```
### Test the deployment
In another terminal
```
$ kubectl get pods

NAME                          READY     STATUS    RESTARTS   AGE
api-status-7b89bb67df-ks6b4   1/1       Running   0          6s
api-status-7b89bb67df-xqbrh   1/1       Running   0          6s
```
```
$ kubectl get deployments

NAME         READY     UP-TO-DATE   AVAILABLE   AGE
api-status   2/2       2            2           13s
```
```
$ kubectl get services

NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
api-status   NodePort    10.105.151.67   <none>        8080:30169/TCP   20s
kubernetes   ClusterIP   10.96.0.1       <none>        443/TCP          10h
```

```
$ minikube service api-status --url

http://ip:port
```

```
$ curl http://ip:port/api/status

[{"CommitID":"687c95e649f985c78a203e601bfb9a2822e0787c"}]
```

## Author

**Pratik Jagrut**
