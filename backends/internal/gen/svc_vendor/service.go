// Code generated by goa v3.20.0, DO NOT EDIT.
//
// svc-vendor service
//
// Command:
// $ goa gen github.com/adrisongomez/pti-ecommerce-site/backends/design -o
// ./internal

package svcvendor

import (
	"context"

	svcvendorviews "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_vendor/views"
	goa "goa.design/goa/v3/pkg"
)

// The product service perform CRUD over the vendor resource
type Service interface {
	// List vendors
	List(context.Context, *ListPayload) (res *VendorList, err error)
	// Create a new product
	Create(context.Context, *VendorInput) (res *Vendor, err error)
	// Create a new product
	DeleteByID(context.Context, *DeleteByIDPayload) (res bool, err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "ecommerce"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "svc-vendor"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"list", "create", "deleteById"}

// DeleteByIDPayload is the payload type of the svc-vendor service deleteById
// method.
type DeleteByIDPayload struct {
	// Unique product identifier
	VendorID int
}

// ListPayload is the payload type of the svc-vendor service list method.
type ListPayload struct {
	// Record per page
	PageSize int
	// Start listing after this resource
	After *int
}

// Pagination information
type PageInfo struct {
	// The starting cursor for pagination
	StartCursor string
	// The ending cursor for pagination
	EndCursor string
	// Indicates if there are more results available
	HasMore bool
	// Total number of resources available
	TotalResource int
}

// Vendor is the result type of the svc-vendor service create method.
type Vendor struct {
	// Key ID
	ID   *int
	Name string
}

type VendorCollection []*Vendor

// VendorInput is the payload type of the svc-vendor service create method.
type VendorInput struct {
	Name string
}

// VendorList is the result type of the svc-vendor service list method.
type VendorList struct {
	// Data
	Data VendorCollection
	// Pagination information
	PageInfo *PageInfo
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "NotFound", false, false, false)
}

// MakeBadRequest builds a goa.ServiceError from an error.
func MakeBadRequest(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "BadRequest", false, false, false)
}

// MakeConflict builds a goa.ServiceError from an error.
func MakeConflict(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "Conflict", false, false, false)
}

// NewVendorList initializes result type VendorList from viewed result type
// VendorList.
func NewVendorList(vres *svcvendorviews.VendorList) *VendorList {
	return newVendorList(vres.Projected)
}

// NewViewedVendorList initializes viewed result type VendorList from result
// type VendorList using the given view.
func NewViewedVendorList(res *VendorList, view string) *svcvendorviews.VendorList {
	p := newVendorListView(res)
	return &svcvendorviews.VendorList{Projected: p, View: "default"}
}

// NewVendor initializes result type Vendor from viewed result type Vendor.
func NewVendor(vres *svcvendorviews.Vendor) *Vendor {
	return newVendor(vres.Projected)
}

// NewViewedVendor initializes viewed result type Vendor from result type
// Vendor using the given view.
func NewViewedVendor(res *Vendor, view string) *svcvendorviews.Vendor {
	p := newVendorView(res)
	return &svcvendorviews.Vendor{Projected: p, View: "default"}
}

// newVendorList converts projected type VendorList to service type VendorList.
func newVendorList(vres *svcvendorviews.VendorListView) *VendorList {
	res := &VendorList{}
	if vres.Data != nil {
		res.Data = newVendorCollection(vres.Data)
	}
	if vres.PageInfo != nil {
		res.PageInfo = newPageInfo(vres.PageInfo)
	}
	return res
}

// newVendorListView projects result type VendorList to projected type
// VendorListView using the "default" view.
func newVendorListView(res *VendorList) *svcvendorviews.VendorListView {
	vres := &svcvendorviews.VendorListView{}
	if res.Data != nil {
		vres.Data = newVendorCollectionView(res.Data)
	}
	if res.PageInfo != nil {
		vres.PageInfo = newPageInfoView(res.PageInfo)
	}
	return vres
}

// newVendorCollection converts projected type VendorCollection to service type
// VendorCollection.
func newVendorCollection(vres svcvendorviews.VendorCollectionView) VendorCollection {
	res := make(VendorCollection, len(vres))
	for i, n := range vres {
		res[i] = newVendor(n)
	}
	return res
}

// newVendorCollectionView projects result type VendorCollection to projected
// type VendorCollectionView using the "default" view.
func newVendorCollectionView(res VendorCollection) svcvendorviews.VendorCollectionView {
	vres := make(svcvendorviews.VendorCollectionView, len(res))
	for i, n := range res {
		vres[i] = newVendorView(n)
	}
	return vres
}

// newVendor converts projected type Vendor to service type Vendor.
func newVendor(vres *svcvendorviews.VendorView) *Vendor {
	res := &Vendor{
		ID: vres.ID,
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	return res
}

// newVendorView projects result type Vendor to projected type VendorView using
// the "default" view.
func newVendorView(res *Vendor) *svcvendorviews.VendorView {
	vres := &svcvendorviews.VendorView{
		ID:   res.ID,
		Name: &res.Name,
	}
	return vres
}

// newPageInfo converts projected type PageInfo to service type PageInfo.
func newPageInfo(vres *svcvendorviews.PageInfoView) *PageInfo {
	res := &PageInfo{}
	if vres.StartCursor != nil {
		res.StartCursor = *vres.StartCursor
	}
	if vres.EndCursor != nil {
		res.EndCursor = *vres.EndCursor
	}
	if vres.HasMore != nil {
		res.HasMore = *vres.HasMore
	}
	if vres.TotalResource != nil {
		res.TotalResource = *vres.TotalResource
	}
	return res
}

// newPageInfoView projects result type PageInfo to projected type PageInfoView
// using the "default" view.
func newPageInfoView(res *PageInfo) *svcvendorviews.PageInfoView {
	vres := &svcvendorviews.PageInfoView{
		StartCursor:   &res.StartCursor,
		EndCursor:     &res.EndCursor,
		HasMore:       &res.HasMore,
		TotalResource: &res.TotalResource,
	}
	return vres
}
