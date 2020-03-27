# For using informer

## what does it do

a tool to monitor pods in a specified namespace from a k8s cluster, and print every event

* learned from `https://github.com/kubernetes/client-go/blob/master/examples/workqueue/main.go`

## how to build

* go version >= 1.14
* use go module

```bash
go build
```

## how to use

* a k8s cluster exists
* have the right access to the cluster
* if to be executed out of cluster, prepare the kube config file at first
* executed in cluster or out of cluster

```bash
Usage of ./using-informer:
  -in
        whether it is in the cluster
  -k string
        kube config file (default "${HOME}/.kube/config")
  -n string
        namespace (default "default")
```

## example

* execute `./using-informer -n istio`
* delete a pod in the k8s cluster 
* stdout:
```
for using informer
+++++++++++++++++++++++++++++
Add a pod:  istio/ratings-v1-6c9dbf6b45-992jz
Add a pod:  istio/productpage-v1-85b9bf9cd7-kjrvw
Add a pod:  istio/reviews-v3-67b4988599-kvpkq
Add a pod:  istio/reviews-v2-568c7c9d8f-trjz2
Add a pod:  istio/details-v1-78d78fbddf-q9r4d
Add a pod:  istio/reviews-v1-564b97f875-2jxgp
Update a pod: istio/details-v1-78d78fbddf-q9r4d =>  Running
Add a pod:  istio/details-v1-78d78fbddf-hgqf6
Update a pod: istio/details-v1-78d78fbddf-hgqf6 =>  Pending
Update a pod: istio/details-v1-78d78fbddf-hgqf6 =>  Pending
Update a pod: istio/details-v1-78d78fbddf-hgqf6 =>  Pending
Update a pod: istio/details-v1-78d78fbddf-hgqf6 =>  Running
Update a pod: istio/details-v1-78d78fbddf-hgqf6 =>  Running
Update a pod: istio/details-v1-78d78fbddf-q9r4d =>  Running
Update a pod: istio/details-v1-78d78fbddf-q9r4d =>  Running
Delete a pod:  istio/details-v1-78d78fbddf-q9r4d
```