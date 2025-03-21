// Code generated by goa v3.20.0, DO NOT EDIT.
//
// svc-healthcheck HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/adrisongomez/pti-ecommerce-site/backends/design -o
// ./internal

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	svchealthcheck "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_healthcheck"
	svchealthcheckviews "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_healthcheck/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildCheckRequest instantiates a HTTP request object with method and path
// set to call the "svc-healthcheck" service "check" endpoint
func (c *Client) BuildCheckRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CheckSvcHealthcheckPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("svc-healthcheck", "check", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeCheckResponse returns a decoder for responses returned by the
// svc-healthcheck check endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeCheckResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body CheckResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("svc-healthcheck", "check", err)
			}
			p := NewCheckHealthcheckResponseOK(&body)
			view := "default"
			vres := &svchealthcheckviews.HealthcheckResponse{Projected: p, View: view}
			if err = svchealthcheckviews.ValidateHealthcheckResponse(vres); err != nil {
				return nil, goahttp.ErrValidationError("svc-healthcheck", "check", err)
			}
			res := svchealthcheck.NewHealthcheckResponse(vres)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("svc-healthcheck", "check", resp.StatusCode, string(body))
		}
	}
}
