# Another ReplicaSet Controller/Operator


## Commands to generate this project

```
$ go mod init another-replicaset-operator

$ kubebuilder version
  Version: main.version{KubeBuilderVersion:"3.1.0", KubernetesVendor:"1.19.2", GitCommit:"92e0349ca7334a0a8e5e499da4fb077eb524e94a", BuildDate:"2021-05-27T17:54:28Z", GoOs:"linux", GoArch:"amd64"}

$ kubebuilder init --domain my.domain
$ kubebuilder create api --group replicaset --version v1alpha1 --kind AnotherReplicaSet

```
