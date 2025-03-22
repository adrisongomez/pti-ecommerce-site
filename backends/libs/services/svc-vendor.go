package services

import (
	"context"
	"fmt"

	svchttp "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/http/svc_vendor/server"
	svc "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_vendor"
	goahttp "goa.design/goa/v3/http"
)

type VendorService struct{}

func (v *VendorService) List(ctx context.Context, payload *svc.ListPayload) (*svc.VendorList, error) {
	return nil, nil
}

func (v *VendorService) Create(ctx context.Context, input *svc.VendorInput) (*svc.Vendor, error) {
	return nil, nil
}
func (v *VendorService) DeleteByID(ctx context.Context, input *svc.DeleteByIDPayload) (bool, error) {
	return false, nil
}

func NewVendorService() *VendorService {
	return &VendorService{}
}

func MountVendorSVC(mux goahttp.Muxer, vendorSvc *VendorService) {
	endpoints := svc.NewEndpoints(vendorSvc)
	req := goahttp.RequestDecoder
	res := goahttp.ResponseEncoder

	handler := svchttp.New(endpoints, mux, req, res, nil, nil)
	svchttp.Mount(mux, handler)

	go func() {
		for _, mount := range handler.Mounts {
			fmt.Printf("%q mounted on %s %s\n", mount.Method, mount.Verb, mount.Pattern)
		}
	}()
}
