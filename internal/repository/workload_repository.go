package repository

import (
	"context"

	v1 "github.com/siphon/siphon/api/v1"
	"github.com/siphon/siphon/internal/model"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type WorkloadRepository interface {
	CreateOrUpdate(ctx context.Context, version *model.Workload) error
	Find(ctx context.Context, id string) (*model.Workload, error)
	All(ctx context.Context) ([]*model.Workload, error)
	Delete(ctx context.Context, id string) error
}

type WorkloadRepositoryImpl struct {
	Client client.Client
}

func (v *WorkloadRepositoryImpl) Create(ctx context.Context, workload *v1.SiphonManagedWorkload) error {
	return v.Client.Create(ctx, workload)
}

func (v *WorkloadRepositoryImpl) All(ctx context.Context) ([]*v1.SiphonManagedWorkload, error) {
	workloadList := &v1.SiphonManagedWorkloadList{}
	if err := v.Client.List(ctx, workloadList); err != nil {
		return nil, err
	}

	list := make([]*v1.SiphonManagedWorkload, 0, len(workloadList.Items))
	for _, workload := range workloadList.Items {
		list = append(list, &workload)
	}

	return list, nil
}
