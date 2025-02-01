package main

import (
	"fmt"
	"net/http"

	"github.com/siphon/siphon/generated/system/v1/systemv1connect"
	"github.com/siphon/siphon/internal/model"
	"github.com/siphon/siphon/internal/service/system"
	"github.com/siphon/siphon/internal/service/workload"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := SetupDatabase()
	if err != nil {
		panic(err)
	}
	err = AutoMigrate(db)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle(systemv1connect.NewSystemServiceHandler(&system.SystemService{}))
	mux.Handle(workload.NewWorkloadServiceHandler(workload.NewWorkloadService(db)))
	http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
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
	); err != nil {
		return err
	}

	return nil
}
