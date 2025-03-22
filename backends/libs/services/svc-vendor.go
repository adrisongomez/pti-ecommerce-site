package services

import (
	"context"
	"fmt"

	"github.com/adrisongomez/pti-ecommerce-site/backends/databases/db"
	svchttp "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/http/svc_vendor/server"
	svc "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_vendor"
	goahttp "goa.design/goa/v3/http"
)

type VendorService struct {
	client *db.PrismaClient
}

func mapVendorToVendorResponse(vendor db.VendorModel) *svc.Vendor {
	return &svc.Vendor{
		ID:   &vendor.ID,
		Name: vendor.Name,
	}
}

func (v *VendorService) List(ctx context.Context, payload *svc.ListPayload) (*svc.VendorList, error) {
	cursor := v.client.Vendor.FindMany().Take(payload.PageSize).Skip(payload.After)
	vendors, err := cursor.Exec(ctx)
	if err != nil {
		return nil, err
	}
	var vendorList svc.VendorCollection = []*svc.Vendor{}
	for _, vendor := range vendors {
		vendorList = append(vendorList, mapVendorToVendorResponse(vendor))
	}
	pageInfo := &svc.PageInfo{
		StartCursor:   payload.After,
		EndCursor:     payload.After + payload.PageSize,
		TotalResource: 100,
	}

	response := &svc.VendorList{
		Data:     vendorList,
		PageInfo: pageInfo,
	}
	return response, nil
}

func (v *VendorService) Create(ctx context.Context, input *svc.VendorInput) (*svc.Vendor, error) {
	createdVendor, err := v.client.Vendor.CreateOne(
		db.Vendor.Name.Set(input.Name),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	response := mapVendorToVendorResponse(*createdVendor)
	return response, nil
}
func (v *VendorService) DeleteByID(ctx context.Context, input *svc.DeleteByIDPayload) (bool, error) {
	_, err := v.client.Vendor.FindUnique(
		db.Vendor.ID.Equals(input.VendorID),
	).Delete().Exec(ctx)

	if err != nil {
		return false, err
	}
	return true, nil
}

func NewVendorService(client *db.PrismaClient) *VendorService {
	return &VendorService{
		client: client,
	}
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
