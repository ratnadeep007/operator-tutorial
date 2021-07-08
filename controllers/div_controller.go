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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/ratnadeep/operator-tutorial/api/v1alpha1"
	basicv1alpha1 "github.com/ratnadeep/operator-tutorial/api/v1alpha1"
)

// DivReconciler reconciles a Div object
type DivReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=basic.calc.k8s,resources=divs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=basic.calc.k8s,resources=divs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=basic.calc.k8s,resources=divs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Div object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *DivReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	instance := &v1alpha1.Div{}

	err := r.Client.Get(ctx, req.NamespacedName, instance)
	if err != nil {
		fmt.Printf("Error getting instance: %s", err)
	}

	instance.Status.Output = instance.Spec.Num1 / instance.Spec.Num2
	err = r.Client.Status().Update(ctx, instance)
	if err != nil {
		fmt.Printf("Error updating instance: %s", err)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DivReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&basicv1alpha1.Div{}).
		Complete(r)
}
