// Code generated by goa v3.20.0, DO NOT EDIT.
//
// svc-healthcheck HTTP client types
//
// Command:
// $ goa gen github.com/adrisongomez/pti-ecommerce-site/backends/design -o
// ./internal

package client

import (
	svchealthcheckviews "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_healthcheck/views"
)

// CheckResponseBody is the type of the "svc-healthcheck" service "check"
// endpoint HTTP response body.
type CheckResponseBody struct {
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// NewCheckHealthcheckResponseOK builds a "svc-healthcheck" service "check"
// endpoint result from a HTTP "OK" response.
func NewCheckHealthcheckResponseOK(body *CheckResponseBody) *svchealthcheckviews.HealthcheckResponseView {
	v := &svchealthcheckviews.HealthcheckResponseView{
		Status: body.Status,
	}

	return v
}
