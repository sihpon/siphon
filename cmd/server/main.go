package server

import (
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/siphon/siphon/generated/system/v1/systemv1connect"
	"github.com/siphon/siphon/internal/container"
	"github.com/siphon/siphon/internal/service/system"
	"github.com/siphon/siphon/internal/service/version"
	"github.com/siphon/siphon/internal/service/workload"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func Start(container container.Containerer) error {
	log.Println("Starting server on localhost:8080")
	mux := http.NewServeMux()
	mux.Handle(systemv1connect.NewSystemServiceHandler(&system.SystemService{}))
	mux.Handle(workload.NewWorkloadServiceHandler(workload.NewWorkloadService(container)))
	mux.Handle(version.NewVersionServiceHandler(version.NewVersionService(container)))
	corsHandler := cors.AllowAll().Handler(h2c.NewHandler(mux, &http2.Server{}))

	http.ListenAndServe(
		"localhost:8080",
		corsHandler,
	)

	return nil
}
