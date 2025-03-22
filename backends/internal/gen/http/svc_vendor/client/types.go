// Code generated by goa v3.20.0, DO NOT EDIT.
//
// svc-vendor HTTP client types
//
// Command:
// $ goa gen github.com/adrisongomez/pti-ecommerce-site/backends/design -o
// ./internal

package client

import (
	svcvendor "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_vendor"
	svcvendorviews "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_vendor/views"
	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "svc-vendor" service "create" endpoint
// HTTP request body.
type CreateRequestBody struct {
	Name string `form:"name" json:"name" xml:"name"`
}

// ListResponseBody is the type of the "svc-vendor" service "list" endpoint
// HTTP response body.
type ListResponseBody struct {
	// Data
	Data VendorCollectionResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	// Pagination information
	PageInfo *PageInfoResponseBody `form:"pageInfo,omitempty" json:"pageInfo,omitempty" xml:"pageInfo,omitempty"`
}

// CreateResponseBody is the type of the "svc-vendor" service "create" endpoint
// HTTP response body.
type CreateResponseBody struct {
	// Key ID
	ID   *int    `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// ListBadRequestResponseBody is the type of the "svc-vendor" service "list"
// endpoint HTTP response body for the "BadRequest" error.
type ListBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// DeleteByIDNotFoundResponseBody is the type of the "svc-vendor" service
// "deleteById" endpoint HTTP response body for the "NotFound" error.
type DeleteByIDNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// VendorCollectionResponseBody is used to define fields on response body types.
type VendorCollectionResponseBody []*VendorResponseBody

// VendorResponseBody is used to define fields on response body types.
type VendorResponseBody struct {
	// Key ID
	ID   *int    `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// PageInfoResponseBody is used to define fields on response body types.
type PageInfoResponseBody struct {
	// The starting cursor for pagination
	StartCursor *int `form:"startCursor,omitempty" json:"startCursor,omitempty" xml:"startCursor,omitempty"`
	// The ending cursor for pagination
	EndCursor *int `form:"endCursor,omitempty" json:"endCursor,omitempty" xml:"endCursor,omitempty"`
	// Indicates if there are more results available
	HasMore *bool `form:"hasMore,omitempty" json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// Total number of resources available
	TotalResource *int `form:"totalResource,omitempty" json:"totalResource,omitempty" xml:"totalResource,omitempty"`
}

// NewCreateRequestBody builds the HTTP request body from the payload of the
// "create" endpoint of the "svc-vendor" service.
func NewCreateRequestBody(p *svcvendor.VendorInput) *CreateRequestBody {
	body := &CreateRequestBody{
		Name: p.Name,
	}
	return body
}

// NewListVendorListOK builds a "svc-vendor" service "list" endpoint result
// from a HTTP "OK" response.
func NewListVendorListOK(body *ListResponseBody) *svcvendorviews.VendorListView {
	v := &svcvendorviews.VendorListView{}
	if body.Data != nil {
		v.Data = make([]*svcvendorviews.VendorView, len(body.Data))
		for i, val := range body.Data {
			v.Data[i] = unmarshalVendorResponseBodyToSvcvendorviewsVendorView(val)
		}
	}
	if body.PageInfo != nil {
		v.PageInfo = unmarshalPageInfoResponseBodyToSvcvendorviewsPageInfoView(body.PageInfo)
	}

	return v
}

// NewListBadRequest builds a svc-vendor service list endpoint BadRequest error.
func NewListBadRequest(body *ListBadRequestResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewCreateVendorCreated builds a "svc-vendor" service "create" endpoint
// result from a HTTP "Created" response.
func NewCreateVendorCreated(body *CreateResponseBody) *svcvendorviews.VendorView {
	v := &svcvendorviews.VendorView{
		ID:   body.ID,
		Name: body.Name,
	}

	return v
}

// NewDeleteByIDNotFound builds a svc-vendor service deleteById endpoint
// NotFound error.
func NewDeleteByIDNotFound(body *DeleteByIDNotFoundResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateListBadRequestResponseBody runs the validations defined on
// list_BadRequest_response_body
func ValidateListBadRequestResponseBody(body *ListBadRequestResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateDeleteByIDNotFoundResponseBody runs the validations defined on
// deleteById_NotFound_response_body
func ValidateDeleteByIDNotFoundResponseBody(body *DeleteByIDNotFoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateVendorCollectionResponseBody runs the validations defined on
// VendorCollectionResponseBody
func ValidateVendorCollectionResponseBody(body VendorCollectionResponseBody) (err error) {
	for _, e := range body {
		if e != nil {
			if err2 := ValidateVendorResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateVendorResponseBody runs the validations defined on VendorResponseBody
func ValidateVendorResponseBody(body *VendorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID != nil {
		if *body.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.id", *body.ID, 1, true))
		}
	}
	return
}

// ValidatePageInfoResponseBody runs the validations defined on
// Page-InfoResponseBody
func ValidatePageInfoResponseBody(body *PageInfoResponseBody) (err error) {
	if body.StartCursor == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("startCursor", "body"))
	}
	if body.EndCursor == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("endCursor", "body"))
	}
	if body.HasMore == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("hasMore", "body"))
	}
	if body.TotalResource == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("totalResource", "body"))
	}
	return
}
