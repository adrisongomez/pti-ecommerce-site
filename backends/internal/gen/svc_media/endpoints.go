// Code generated by goa v3.20.0, DO NOT EDIT.
//
// svc-media endpoints
//
// Command:
// $ goa gen github.com/adrisongomez/pti-ecommerce-site/backends/design -o
// ./internal

package svcmedia

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "svc-media" service endpoints.
type Endpoints struct {
	List       goa.Endpoint
	GetByID    goa.Endpoint
	Create     goa.Endpoint
	DeleteByID goa.Endpoint
}

// NewEndpoints wraps the methods of the "svc-media" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		List:       NewListEndpoint(s),
		GetByID:    NewGetByIDEndpoint(s),
		Create:     NewCreateEndpoint(s),
		DeleteByID: NewDeleteByIDEndpoint(s),
	}
}

// Use applies the given middleware to all the "svc-media" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
	e.GetByID = m(e.GetByID)
	e.Create = m(e.Create)
	e.DeleteByID = m(e.DeleteByID)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "svc-media".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ListPayload)
		res, err := s.List(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedMediaList(res, "default")
		return vres, nil
	}
}

// NewGetByIDEndpoint returns an endpoint function that calls the method
// "getById" of service "svc-media".
func NewGetByIDEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetByIDPayload)
		res, err := s.GetByID(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedMedia(res, "default")
		return vres, nil
	}
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "svc-media".
func NewCreateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*MediaInput)
		res, err := s.Create(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedMedia(res, "default")
		return vres, nil
	}
}

// NewDeleteByIDEndpoint returns an endpoint function that calls the method
// "deleteById" of service "svc-media".
func NewDeleteByIDEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DeleteByIDPayload)
		return s.DeleteByID(ctx, p)
	}
}
