package system

import (
	"context"

	"connectrpc.com/connect"
	systemv1 "github.com/siphon/siphon/generated/system/v1"
)

type SystemService struct{}

func (s *SystemService) Health(
	context.Context,
	*connect.Request[systemv1.HealthRequest],
) (*connect.Response[systemv1.HealthResponse], error) {
	return connect.NewResponse(&systemv1.HealthResponse{
		Status: "OK",
	}), nil
}
