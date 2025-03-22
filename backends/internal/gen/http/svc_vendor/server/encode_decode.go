// Code generated by goa v3.20.0, DO NOT EDIT.
//
// svc-vendor HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/adrisongomez/pti-ecommerce-site/backends/design -o
// ./internal

package server

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strconv"

	svcvendorviews "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_vendor/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeListResponse returns an encoder for responses returned by the
// svc-vendor list endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res := v.(*svcvendorviews.VendorList)
		enc := encoder(ctx, w)
		body := NewListResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListRequest returns a decoder for requests sent to the svc-vendor list
// endpoint.
func DecodeListRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			pageSize int
			after    int
			err      error
		)
		qp := r.URL.Query()
		{
			pageSizeRaw := qp.Get("pageSize")
			if pageSizeRaw == "" {
				pageSize = 10
			} else {
				v, err2 := strconv.ParseInt(pageSizeRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("pageSize", pageSizeRaw, "integer"))
				}
				pageSize = int(v)
			}
		}
		if pageSize < 10 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("pageSize", pageSize, 10, true))
		}
		if pageSize > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("pageSize", pageSize, 100, false))
		}
		{
			afterRaw := qp.Get("after")
			if afterRaw != "" {
				v, err2 := strconv.ParseInt(afterRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("after", afterRaw, "integer"))
				}
				after = int(v)
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewListPayload(pageSize, after)

		return payload, nil
	}
}

// EncodeListError returns an encoder for errors returned by the list
// svc-vendor endpoint.
func EncodeListError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "BadRequest":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewListBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateResponse returns an encoder for responses returned by the
// svc-vendor create endpoint.
func EncodeCreateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res := v.(*svcvendorviews.Vendor)
		enc := encoder(ctx, w)
		body := NewCreateResponseBody(res.Projected)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateRequest returns a decoder for requests sent to the svc-vendor
// create endpoint.
func DecodeCreateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body CreateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCreateVendorInput(&body)

		return payload, nil
	}
}

// EncodeDeleteByIDResponse returns an encoder for responses returned by the
// svc-vendor deleteById endpoint.
func EncodeDeleteByIDResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(bool)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeDeleteByIDRequest returns a decoder for requests sent to the
// svc-vendor deleteById endpoint.
func DecodeDeleteByIDRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			vendorID int
			err      error

			params = mux.Vars(r)
		)
		{
			vendorIDRaw := params["vendorId"]
			v, err2 := strconv.ParseInt(vendorIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("vendorId", vendorIDRaw, "integer"))
			}
			vendorID = int(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewDeleteByIDPayload(vendorID)

		return payload, nil
	}
}

// EncodeDeleteByIDError returns an encoder for errors returned by the
// deleteById svc-vendor endpoint.
func EncodeDeleteByIDError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "NotFound":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDeleteByIDNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalSvcvendorviewsVendorViewToVendorResponseBody builds a value of type
// *VendorResponseBody from a value of type *svcvendorviews.VendorView.
func marshalSvcvendorviewsVendorViewToVendorResponseBody(v *svcvendorviews.VendorView) *VendorResponseBody {
	if v == nil {
		return nil
	}
	res := &VendorResponseBody{
		ID:   v.ID,
		Name: *v.Name,
	}

	return res
}

// marshalSvcvendorviewsPageInfoViewToPageInfoResponseBody builds a value of
// type *PageInfoResponseBody from a value of type *svcvendorviews.PageInfoView.
func marshalSvcvendorviewsPageInfoViewToPageInfoResponseBody(v *svcvendorviews.PageInfoView) *PageInfoResponseBody {
	if v == nil {
		return nil
	}
	res := &PageInfoResponseBody{
		StartCursor:   *v.StartCursor,
		EndCursor:     *v.EndCursor,
		HasMore:       *v.HasMore,
		TotalResource: *v.TotalResource,
	}

	return res
}
