package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adrisongomez/pti-ecommerce-site/backends/databases/db"
	"github.com/adrisongomez/pti-ecommerce-site/backends/pkg/loggers"
	svc "github.com/adrisongomez/pti-ecommerce-site/backends/pkg/services"
	"go.uber.org/zap"
	goahttp "goa.design/goa/v3/http"
)

func main() {
	var (
		port = "3030"
	)
	logger := loggers.CreateLogger("ecommerce-api")
	client := db.NewClient()

	zap.ReplaceGlobals(logger)

	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	defer func() {
		logger.Sync()
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	productSvc := svc.NewProductService(client)
	healthcheckSvc := svc.NewHealthcheckService()
	mediaSvc := svc.NewMediaService(client)
	mux := goahttp.NewMuxer()
	svc.MountMediaSVC(mux, mediaSvc)
	svc.MountHealtcheckSVC(mux, healthcheckSvc)
	svc.MountProductSVC(mux, productSvc)
	// middleware.Debug(mux, fmt.Fprint)
	server := &http.Server{Addr: ":" + port, Handler: mux}

	logger.Info(fmt.Sprintf("Starting server on :%s\n", port))
	if err := server.ListenAndServe(); err != nil {
		logger.Fatal("Service shutdown due to error", zap.Any("error", err))
		log.Fatal(err)
	}
}
