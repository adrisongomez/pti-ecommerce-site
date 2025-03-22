package services

import (
	"context"
	"fmt"

	. "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_products"

	product_genhttp "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/http/svc_products/server"
	product_gen "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_products"
	goahttp "goa.design/goa/v3/http"
)

type ProductService struct{}

func (p *ProductService) ListProduct(ctx context.Context, payload *ListProductPayload) (*ProductsList, error) {
	return nil, nil
}

func (p *ProductService) GetProductByID(ctx context.Context, payload *GetProductByIDPayload) (*Product, error) {
	return nil, nil
}

func (p *ProductService) CreateProduct(ctx context.Context, payload *ProductInput) (*Product, error) {
	return nil, nil
}
func (p *ProductService) UpdateProductByID(ctx context.Context, payload *UpdateProductByIDPayload) (*Product, error) {
	return nil, nil
}
func (p *ProductService) DeleteProductByID(ctx context.Context, payload *DeleteProductByIDPayload) (bool, error) {
	return false, nil
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func MountProductSVC(mux goahttp.Muxer, svc *ProductService) {
	endpoints := product_gen.NewEndpoints(svc)
	req := goahttp.RequestDecoder
	res := goahttp.ResponseEncoder

	handler := product_genhttp.New(endpoints, mux, req, res, nil, nil)
	product_genhttp.Mount(mux, handler)

	go func() {
		for _, mount := range handler.Mounts {
			fmt.Printf("%q mounted on %s %s\n", mount.Method, mount.Verb, mount.Pattern)
		}
	}()
}
