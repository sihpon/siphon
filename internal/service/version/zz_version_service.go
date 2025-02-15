package version

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	proto "github.com/siphon/siphon/generated/version/v1"
	"github.com/siphon/siphon/generated/version/v1/versionv1connect"
)

type Versioner interface {
	Create(ctx context.Context, request *proto.CreateRequest) (*proto.CreateResponse, error)
	Delete(ctx context.Context, request *proto.DeleteRequest) (*proto.DeleteResponse, error)
	Get(ctx context.Context, request *proto.GetRequest) (*proto.GetResponse, error)
	List(ctx context.Context, request *proto.ListRequest) (*proto.ListResponse, error)
	Update(ctx context.Context, request *proto.UpdateRequest) (*proto.UpdateResponse, error)
}

var _ versionv1connect.VersionServiceHandler = &ZZVersionService{}

type ZZVersionService struct {
	Versioner
}

func NewVersionServiceHandler(handlers Versioner) (string, http.Handler) {
	return versionv1connect.NewVersionServiceHandler(&ZZVersionService{Versioner: handlers})
}

// Create implements versionv1connect.WorkloadServiceHandler.
func (z *ZZVersionService) Create(ctx context.Context, request *connect.Request[proto.CreateRequest]) (*connect.Response[proto.CreateResponse], error) {
	response, err := z.Versioner.Create(ctx, request.Msg)
	return connect.NewResponse(response), err
}

// Delete implements versionv1connect.WorkloadServiceHandler.
func (z *ZZVersionService) Delete(ctx context.Context, request *connect.Request[proto.DeleteRequest]) (*connect.Response[proto.DeleteResponse], error) {
	response, err := z.Versioner.Delete(ctx, request.Msg)
	return connect.NewResponse(response), err
}

// Get implements versionv1connect.WorkloadServiceHandler.
func (z *ZZVersionService) Get(ctx context.Context, request *connect.Request[proto.GetRequest]) (*connect.Response[proto.GetResponse], error) {
	response, err := z.Versioner.Get(ctx, request.Msg)
	return connect.NewResponse(response), err
}

// List implements versionv1connect.WorkloadServiceHandler.
func (z *ZZVersionService) List(ctx context.Context, request *connect.Request[proto.ListRequest]) (*connect.Response[proto.ListResponse], error) {
	response, err := z.Versioner.List(ctx, request.Msg)
	return connect.NewResponse(response), err
}

// Update implements versionv1connect.WorkloadServiceHandler.
func (z *ZZVersionService) Update(ctx context.Context, request *connect.Request[proto.UpdateRequest]) (*connect.Response[proto.UpdateResponse], error) {
	response, err := z.Versioner.Update(ctx, request.Msg)
	return connect.NewResponse(response), err
}
