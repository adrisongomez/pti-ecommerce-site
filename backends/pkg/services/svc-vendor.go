package services

import (
	"context"
	"fmt"

	"github.com/adrisongomez/pti-ecommerce-site/backends/databases/db"
	svchttp "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/http/svc_vendor/server"
	svc "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_vendor"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils"
	"go.uber.org/zap"
	goahttp "goa.design/goa/v3/http"
)

type VendorService struct {
	client *db.PrismaClient
}

func MapVendorToVendorResponse(vendor db.VendorModel) *svc.Vendor {
	return &svc.Vendor{
		ID:   &vendor.ID,
		Name: vendor.Name,
	}
}

func (v *VendorService) count(ctx context.Context) (int, error) {
	var resp []struct {
		Count db.BigInt `json:"count"`
	}
	err := v.client.Prisma.QueryRaw("SELECT count(*) FROM project.vendors").Exec(ctx, &resp)
	if err != nil {
		return 0, err
	}
	if len(resp) == 0 {
		return 0, nil
	}
	count := int(resp[0].Count)
	return count, nil

}

func (v *VendorService) List(ctx context.Context, payload *svc.ListPayload) (*svc.VendorList, error) {
	cursor := v.client.Vendor.FindMany()
	vendors, err := cursor.Take(payload.PageSize).Skip(payload.After).Exec(ctx)
	if err != nil {
		return nil, err
	}

	var vendorList svc.VendorCollection = []*svc.Vendor{}
	for _, vendor := range vendors {
		vendorList = append(vendorList, MapVendorToVendorResponse(vendor))
	}
	count, err := v.count(ctx)

	if err != nil {
		return nil, err
	}
	nextPageCursor := utils.MinInt(count, payload.After+payload.PageSize)
	pageInfo := &svc.PageInfo{
		StartCursor:   payload.After,
		EndCursor:     nextPageCursor,
		TotalResource: count,
		HasMore:       nextPageCursor < count,
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
	response := MapVendorToVendorResponse(*createdVendor)
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
			zap.L().Info(fmt.Sprintf("%q mounted on %s %s\n", mount.Method, mount.Verb, mount.Pattern))
		}
	}()
}
