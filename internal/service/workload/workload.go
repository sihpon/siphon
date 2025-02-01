package workload

import (
	"context"

	"connectrpc.com/connect"
	workloadv1 "github.com/siphon/siphon/generated/workload/v1"
)

type WorkloadService struct{}

func (s *WorkloadService) ListWorkloads(
	context.Context,
	*connect.Request[workloadv1.ListWorkloadsRequest],
) (*connect.Response[workloadv1.ListWorkloadsResponse], error) {
	return connect.NewResponse(&workloadv1.ListWorkloadsResponse{
		Workloads: []*workloadv1.Workload{
			{
				Id:          "hoge",
				Name:        "hoge",
				Description: "",
				Version:     "v2.0.0",
				Status:      "",
				Url:         "",
			},
		},
	}), nil
}

func (s *WorkloadService) GetWorkload(
	context.Context,
	*connect.Request[workloadv1.GetWorkloadRequest],
) (*connect.Response[workloadv1.GetWorkloadResponse], error) {
	return connect.NewResponse(&workloadv1.GetWorkloadResponse{
		Workload: &workloadv1.Workload{},
	}), nil
}

func (s *WorkloadService) CreateWorkload(
	context.Context,
	*connect.Request[workloadv1.CreateWorkloadRequest],
) (*connect.Response[workloadv1.CreateWorkloadResponse], error) {
	return connect.NewResponse(&workloadv1.CreateWorkloadResponse{}), nil
}

func (s *WorkloadService) UpdateWorkload(
	context.Context,
	*connect.Request[workloadv1.UpdateWorkloadRequest],
) (*connect.Response[workloadv1.UpdateWorkloadResponse], error) {
	return connect.NewResponse(&workloadv1.UpdateWorkloadResponse{}), nil
}

func (s *WorkloadService) DeleteWorkload(
	context.Context,
	*connect.Request[workloadv1.DeleteWorkloadRequest],
) (*connect.Response[workloadv1.DeleteWorkloadResponse], error) {
	return connect.NewResponse(&workloadv1.DeleteWorkloadResponse{}), nil
}
