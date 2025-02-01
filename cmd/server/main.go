package main

import (
	"net/http"

	"github.com/siphon/siphon/gen/system/v1/systemv1connect"
	"github.com/siphon/siphon/gen/workload/v1/workloadv1connect"
	"github.com/siphon/siphon/internal/service/system"
	"github.com/siphon/siphon/internal/service/workload"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle(systemv1connect.NewSystemServiceHandler(&system.SystemService{}))
	mux.Handle(workloadv1connect.NewWorkloadServiceHandler(&workload.WorkloadService{}))
	http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
