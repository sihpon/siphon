package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/siphon/siphon/generated/system/v1/systemv1connect"
	"github.com/siphon/siphon/internal/model"
	"github.com/siphon/siphon/internal/service/system"
	"github.com/siphon/siphon/internal/service/version"
	"github.com/siphon/siphon/internal/service/workload"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func Start(client client.Client, schema *runtime.Scheme) error {
	db, err := SetupDatabase()
	if err != nil {
		return err
	}
	err = AutoMigrate(db)
	if err != nil {
		return err
	}

	log.Println("Starting server on localhost:8080")
	mux := http.NewServeMux()
	mux.Handle(systemv1connect.NewSystemServiceHandler(&system.SystemService{}))
	mux.Handle(workload.NewWorkloadServiceHandler(workload.NewWorkloadService(db)))
	mux.Handle(version.NewVersionServiceHandler(version.NewVersionService(db, client, schema)))
	corsHandler := cors.AllowAll().Handler(h2c.NewHandler(mux, &http2.Server{}))
	http.ListenAndServe(
		"localhost:8080",
		corsHandler,
	)

	return nil
}

func SetupDatabase() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"local",
		"local",
		"localhost",
		3306,
		"local",
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) (err error) {
	if err := db.AutoMigrate(
		&model.Workload{},
		&model.Version{},
	); err != nil {
		return err
	}

	return nil
}
