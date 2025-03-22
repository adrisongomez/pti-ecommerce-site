// Code generated by goa v3.20.0, DO NOT EDIT.
//
// svc-vendor HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/adrisongomez/pti-ecommerce-site/backends/design -o
// ./internal

package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	svcvendor "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_vendor"
	svcvendorviews "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_vendor/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "svc-vendor" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListSvcVendorPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("svc-vendor", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the svc-vendor
// list server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*svcvendor.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("svc-vendor", "list", "*svcvendor.ListPayload", v)
		}
		values := req.URL.Query()
		values.Add("pageSize", fmt.Sprintf("%v", p.PageSize))
		if p.After != nil {
			values.Add("after", fmt.Sprintf("%v", *p.After))
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeListResponse returns a decoder for responses returned by the
// svc-vendor list endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeListResponse may return the following errors:
//   - "BadRequest" (type *goa.ServiceError): http.StatusBadRequest
//   - error: internal error
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("svc-vendor", "list", err)
			}
			p := NewListVendorListOK(&body)
			view := "default"
			vres := &svcvendorviews.VendorList{Projected: p, View: view}
			if err = svcvendorviews.ValidateVendorList(vres); err != nil {
				return nil, goahttp.ErrValidationError("svc-vendor", "list", err)
			}
			res := svcvendor.NewVendorList(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ListBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("svc-vendor", "list", err)
			}
			err = ValidateListBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("svc-vendor", "list", err)
			}
			return nil, NewListBadRequest(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("svc-vendor", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "svc-vendor" service "create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateSvcVendorPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("svc-vendor", "create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the svc-vendor
// create server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*svcvendor.VendorInput)
		if !ok {
			return goahttp.ErrInvalidType("svc-vendor", "create", "*svcvendor.VendorInput", v)
		}
		body := NewCreateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("svc-vendor", "create", err)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the
// svc-vendor create endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body CreateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("svc-vendor", "create", err)
			}
			p := NewCreateVendorCreated(&body)
			view := "default"
			vres := &svcvendorviews.Vendor{Projected: p, View: view}
			if err = svcvendorviews.ValidateVendor(vres); err != nil {
				return nil, goahttp.ErrValidationError("svc-vendor", "create", err)
			}
			res := svcvendor.NewVendor(vres)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("svc-vendor", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteByIDRequest instantiates a HTTP request object with method and
// path set to call the "svc-vendor" service "deleteById" endpoint
func (c *Client) BuildDeleteByIDRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		vendorID int
	)
	{
		p, ok := v.(*svcvendor.DeleteByIDPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("svc-vendor", "deleteById", "*svcvendor.DeleteByIDPayload", v)
		}
		vendorID = p.VendorID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteByIDSvcVendorPath(vendorID)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("svc-vendor", "deleteById", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeDeleteByIDResponse returns a decoder for responses returned by the
// svc-vendor deleteById endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeDeleteByIDResponse may return the following errors:
//   - "NotFound" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodeDeleteByIDResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body bool
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("svc-vendor", "deleteById", err)
			}
			return body, nil
		case http.StatusNotFound:
			var (
				body DeleteByIDNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("svc-vendor", "deleteById", err)
			}
			err = ValidateDeleteByIDNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("svc-vendor", "deleteById", err)
			}
			return nil, NewDeleteByIDNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("svc-vendor", "deleteById", resp.StatusCode, string(body))
		}
	}
}

// unmarshalVendorResponseBodyToSvcvendorviewsVendorView builds a value of type
// *svcvendorviews.VendorView from a value of type *VendorResponseBody.
func unmarshalVendorResponseBodyToSvcvendorviewsVendorView(v *VendorResponseBody) *svcvendorviews.VendorView {
	if v == nil {
		return nil
	}
	res := &svcvendorviews.VendorView{
		ID:   v.ID,
		Name: v.Name,
	}

	return res
}

// unmarshalPageInfoResponseBodyToSvcvendorviewsPageInfoView builds a value of
// type *svcvendorviews.PageInfoView from a value of type *PageInfoResponseBody.
func unmarshalPageInfoResponseBodyToSvcvendorviewsPageInfoView(v *PageInfoResponseBody) *svcvendorviews.PageInfoView {
	if v == nil {
		return nil
	}
	res := &svcvendorviews.PageInfoView{
		StartCursor:   v.StartCursor,
		EndCursor:     v.EndCursor,
		HasMore:       v.HasMore,
		TotalResource: v.TotalResource,
	}

	return res
}
