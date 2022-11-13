# Another ReplicaSet Controller/Operator

The objective is to implement a replicaset controller similar to [Kube replicaset controller](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/) using [kubebuilder quick start](https://book.kubebuilder.io/quick-start.html). The process of implementating a controller will help the developer better understanding the challenges and understanding of how Kubernetes controllers work.

## Steps to generate this project

```
$ go mod init another-replicaset-operator

$ kubebuilder version
  Version: main.version{KubeBuilderVersion:"3.1.0", KubernetesVendor:"1.19.2", GitCommit:"92e0349ca7334a0a8e5e499da4fb077eb524e94a", BuildDate:"2021-05-27T17:54:28Z", GoOs:"linux", GoArch:"amd64"}

$ kubebuilder init --domain my.domain
$ kubebuilder create api --group replicaset --version v1alpha1 --kind AnotherReplicaSet

```

## Steps after changing the API

* If you are editing the API definitions, generate the manifests such as Custom Resources (CRs) or Custom Resource Defintions (CRDs) using
```
make manifests
```

## Steps to run the operator

* Before running the operator , check if the api group already exists
```
kubectl api-resources --api-group=replicaset.my.domain
```
* Install the CRDs into the cluster:
```
make install
```
* Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running)

```
make run
```
* Check the api-resources (optional)
```
$ kubectl api-resources --api-group=replicaset.my.domain
NAME                 SHORTNAMES   APIVERSION                      NAMESPACED   KIND
anotherreplicasets                replicaset.my.domain/v1alpha1   true         AnotherReplicaSet
```
* Apply the custom resource in another terminal which has KUBECONFIG configured for the kube environment
```
$ kubectl apply -f resources/cr.yaml
```

## Uninstall CRDs

* To delete your CRDs from the cluster

```
make uninstall
```

## Undeploy controller

* Undeploy the controller to the cluster
```
make undeploy
```
