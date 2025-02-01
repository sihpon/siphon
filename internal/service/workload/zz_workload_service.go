package workload

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	proto "github.com/siphon/siphon/generated/workload/v1"
	"github.com/siphon/siphon/generated/workload/v1/workloadv1connect"
)

type Workloader interface {
	Create(ctx context.Context, request *proto.CreateRequest) (*proto.CreateResponse, error)
	Delete(ctx context.Context, request *proto.DeleteRequest) (*proto.DeleteResponse, error)
	Get(ctx context.Context, request *proto.GetRequest) (*proto.GetResponse, error)
	List(ctx context.Context, request *proto.ListRequest) (*proto.ListResponse, error)
	Update(ctx context.Context, request *proto.UpdateRequest) (*proto.UpdateResponse, error)
}

var _ workloadv1connect.WorkloadServiceHandler = &ZZWorkloadService{}

type ZZWorkloadService struct {
	Workloader
}

func NewWorkloadServiceHandler(handlers Workloader) (string, http.Handler) {
	return workloadv1connect.NewWorkloadServiceHandler(&ZZWorkloadService{Workloader: handlers})
}

// Create implements workloadv1connect.WorkloadServiceHandler.
func (z *ZZWorkloadService) Create(ctx context.Context, request *connect.Request[proto.CreateRequest]) (*connect.Response[proto.CreateResponse], error) {
	response, err := z.Workloader.Create(ctx, request.Msg)
	return connect.NewResponse(response), err
}

// Delete implements workloadv1connect.WorkloadServiceHandler.
func (z *ZZWorkloadService) Delete(ctx context.Context, request *connect.Request[proto.DeleteRequest]) (*connect.Response[proto.DeleteResponse], error) {
	response, err := z.Workloader.Delete(ctx, request.Msg)
	return connect.NewResponse(response), err
}

// Get implements workloadv1connect.WorkloadServiceHandler.
func (z *ZZWorkloadService) Get(ctx context.Context, request *connect.Request[proto.GetRequest]) (*connect.Response[proto.GetResponse], error) {
	response, err := z.Workloader.Get(ctx, request.Msg)
	return connect.NewResponse(response), err
}

// List implements workloadv1connect.WorkloadServiceHandler.
func (z *ZZWorkloadService) List(ctx context.Context, request *connect.Request[proto.ListRequest]) (*connect.Response[proto.ListResponse], error) {
	response, err := z.Workloader.List(ctx, request.Msg)
	return connect.NewResponse(response), err
}

// Update implements workloadv1connect.WorkloadServiceHandler.
func (z *ZZWorkloadService) Update(ctx context.Context, request *connect.Request[proto.UpdateRequest]) (*connect.Response[proto.UpdateResponse], error) {
	response, err := z.Workloader.Update(ctx, request.Msg)
	return connect.NewResponse(response), err
}
