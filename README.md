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

### Steps to run the operator

* Before running the operator , check if the api group already exists
```
$  kubectl api-resources --api-group=replicaset.my.domain
```
* make install

```
$ kubectl api-resources --api-group=replicaset.my.domain
NAME                 SHORTNAMES   APIVERSION                      NAMESPACED   KIND
anotherreplicasets                replicaset.my.domain/v1alpha1   true         AnotherReplicaSet
```
## Other Examples

* https://github.com/hrishin/podset-operator
* https://github.com/kubernetes/sample-controller