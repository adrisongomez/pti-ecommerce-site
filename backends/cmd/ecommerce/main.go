package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adrisongomez/pti-ecommerce-site/backends/databases/db"
	svc "github.com/adrisongomez/pti-ecommerce-site/backends/libs/services"
	goahttp "goa.design/goa/v3/http"
)

func main() {
	var (
		port = "3030"
	)
	client := db.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	productSvc := svc.NewProductService()
	healthcheckSvc := svc.NewHealthcheckService()
	vendorSvc := svc.NewVendorService(client)
	mediaSvc := svc.NewMediaService(client)
	mux := goahttp.NewMuxer()
	svc.MountMediaSVC(mux, mediaSvc)
	svc.MountHealtcheckSVC(mux, healthcheckSvc)
	svc.MountProductSVC(mux, productSvc)
	svc.MountVendorSVC(mux, vendorSvc)

	server := &http.Server{Addr: ":" + port, Handler: mux}

	fmt.Printf("Starting server on :%s\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
