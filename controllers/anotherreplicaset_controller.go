/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	replicasetv1alpha1 "github.com/LalatenduMohanty/another-replicaset-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

var log = logf.Log.WithName("another_replica_set")

// AnotherReplicaSetReconciler reconciles a AnotherReplicaSet object
type AnotherReplicaSetReconciler struct {
	Client client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=replicaset.my.domain,resources=anotherreplicasets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=replicaset.my.domain,resources=anotherreplicasets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=replicaset.my.domain,resources=anotherreplicasets/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the AnotherReplicaSet object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *AnotherReplicaSetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	reqLogger := logf.Log.WithName("another-replicaset-controller")
	reqLogger.Info("Reconciling another-replicaset controller")

	// Fetch the another replica set value from the instance
	instance := &replicasetv1alpha1.AnotherReplicaSet{}
	err := r.Client.Get(context.TODO(), req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			reqLogger.Info("Request object not found")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		reqLogger.Info("Error reading the object, requeuing")
		return ctrl.Result{}, err
	}
	//return ctrl.Result{}, nil

	// check the value of Replicas in AnotherReplicaSetSpec of the in-cluser resource
	// get actual pod count from the in-cluster custom resource
	replicas := instance.Spec.Replicas
	reqLogger.Info(fmt.Sprintf("Replica count in the custom resource: %d", replicas))

	//Get actual number of pods
	reqLogger.Info(fmt.Sprintf("Trying to get the actual in cluster pod count"))

	inClusterPodList := &corev1.PodList{}
	labelSelector := labels.SelectorFromSet(labelsForApp(instance.Name))
	listOps := &client.ListOptions{Namespace: instance.Namespace, LabelSelector: labelSelector}
	err = r.Client.List(context.TODO(), inClusterPodList, listOps)
	if err != nil {
		reqLogger.Error(err, "failed to list pods", "Namespace",
			instance.Namespace, "Name", instance.Name)
		return reconcile.Result{}, err
	}
	podNames := getPodNames(inClusterPodList.Items)
	reqLogger.Info("got pod names", "podNames", podNames)

	//return ctrl.Result{}, nil
	return reconcile.Result{RequeueAfter: time.Second * 10}, nil
}

func getPodNames(pods []corev1.Pod) []string {
	var podNames []string
	for _, pod := range pods {
		podNames = append(podNames, pod.Name)
	}
	return podNames
}

// labelsForApp returns the labels for selecting the resources
// belonging to the given AppService CR name.
func labelsForApp(name string) map[string]string {
	return map[string]string{"app": "test"}
}

// SetupWithManager sets up the controller with the Manager.
func (r *AnotherReplicaSetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&replicasetv1alpha1.AnotherReplicaSet{}).
		Complete(r)
}
