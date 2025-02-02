package controller

import (
	"context"
	"fmt"

	v1 "github.com/siphon/siphon/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *SiphonReconciler) Reconciler(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	deployments := &appsv1.DeploymentList{}
	if err := r.List(ctx, deployments, client.MatchingFields{"metadata.annotations[siphon-managed]": "true"}); err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to list Deployments: %w", err)
	}

	info := make([]v1.DeploymentInfo, 0, len(deployments.Items))
	for _, deployment := range deployments.Items {
		info = append(info, v1.DeploymentInfo{
			Name:      deployment.Name,
			Namespace: deployment.Namespace,
			Replicas:  deployment.Status.Replicas,
			Status:    deployment.Status.String(),
		})
	}

	deploymentStore := &v1.Siphon{ObjectMeta: metav1.ObjectMeta{
		Name:      "siphpon-store",
		Namespace: "default",
	}}

	_, err := ctrl.CreateOrUpdate(ctx, r.Client, deploymentStore, func() error {
		deploymentStore.Spec.Deployments = info
		return nil
	})
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to create DeploymentStore: %w", err)
	}

	return ctrl.Result{}, nil
}
