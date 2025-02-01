package workload

import (
	"context"

	"connectrpc.com/connect"
	workloadv1 "github.com/siphon/siphon/generated/workload/v1"
)

type WorkloadService struct{}

func (s *WorkloadService) List(
	context.Context,
	*connect.Request[workloadv1.ListRequest],
) (*connect.Response[workloadv1.ListResponse], error) {
	return connect.NewResponse(&workloadv1.ListResponse{
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

func (s *WorkloadService) Get(
	context.Context,
	*connect.Request[workloadv1.GetRequest],
) (*connect.Response[workloadv1.GetResponse], error) {
	return connect.NewResponse(&workloadv1.GetResponse{
		Workload: &workloadv1.Workload{},
	}), nil
}

func (s *WorkloadService) Create(
	context.Context,
	*connect.Request[workloadv1.CreateRequest],
) (*connect.Response[workloadv1.CreateResponse], error) {
	return connect.NewResponse(&workloadv1.CreateResponse{}), nil
}

func (s *WorkloadService) Update(
	context.Context,
	*connect.Request[workloadv1.UpdateRequest],
) (*connect.Response[workloadv1.UpdateResponse], error) {
	return connect.NewResponse(&workloadv1.UpdateResponse{}), nil
}

func (s *WorkloadService) Delete(
	context.Context,
	*connect.Request[workloadv1.DeleteRequest],
) (*connect.Response[workloadv1.DeleteResponse], error) {
	return connect.NewResponse(&workloadv1.DeleteResponse{}), nil
}
