package main

import (
	"fmt"
	"log"
	"net/http"

	product_genhttp "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/http/svc_products/server"
	product_gen "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_products"
	svc "github.com/adrisongomez/pti-ecommerce-site/backends/libs/services"
	goahttp "goa.design/goa/v3/http"

	openapisvf "goa.design/examples/files/gen/http/openapi/server"
)

func main() {
	var (
		port = "3030"
	)

	product_svc := svc.NewProductService()
	endpoints := product_gen.NewEndpoints(product_svc)
	mux := goahttp.NewMuxer()
	req := goahttp.RequestDecoder
	res := goahttp.ResponseEncoder
	openapiServer := openapisvf.New(nil, mux, req, res, nil, nil, nil, nil, nil, nil)

	openapisvf.Mount(mux, openapiServer)

	handler := product_genhttp.New(endpoints, mux, req, res, nil, nil)

	product_genhttp.Mount(mux, handler)

	server := &http.Server{Addr: ":" + port, Handler: mux}

	for _, mount := range handler.Mounts {
		fmt.Printf("%q mounted on %s %s\n", mount.Method, mount.Verb, mount.Pattern)
	}

	fmt.Printf("Starting server on :%s\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
