package services

import (
	. "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_products"
	"golang.org/x/net/context"
)

type ProductService struct {
}

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
