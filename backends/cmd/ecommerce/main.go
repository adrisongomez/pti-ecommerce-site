package main

import (
	"fmt"
	"log"
	"net/http"

	svc "github.com/adrisongomez/pti-ecommerce-site/backends/libs/services"
	goahttp "goa.design/goa/v3/http"
)

func main() {
	var (
		port = "3030"
	)

	productSvc := svc.NewProductService()
	healthcheckSvc := svc.NewHealthcheckService()
	vendorSvc := svc.NewVendorService()
	mux := goahttp.NewMuxer()

	svc.MountHealtcheckSVC(mux, healthcheckSvc)
	svc.MountProductSVC(mux, productSvc)
	svc.MountVendorSVC(mux, vendorSvc)

	server := &http.Server{Addr: ":" + port, Handler: mux}

	fmt.Printf("Starting server on :%s\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
