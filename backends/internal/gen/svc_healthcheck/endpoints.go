// Code generated by goa v3.20.0, DO NOT EDIT.
//
// svc-healthcheck endpoints
//
// Command:
// $ goa gen github.com/adrisongomez/pti-ecommerce-site/backends/design -o
// ./internal

package svchealthcheck

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "svc-healthcheck" service endpoints.
type Endpoints struct {
	Check goa.Endpoint
}

// NewEndpoints wraps the methods of the "svc-healthcheck" service with
// endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Check: NewCheckEndpoint(s),
	}
}

// Use applies the given middleware to all the "svc-healthcheck" service
// endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Check = m(e.Check)
}

// NewCheckEndpoint returns an endpoint function that calls the method "check"
// of service "svc-healthcheck".
func NewCheckEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		res, err := s.Check(ctx)
		if err != nil {
			return nil, err
		}
		vres := NewViewedHealthcheckResponse(res, "default")
		return vres, nil
	}
}
