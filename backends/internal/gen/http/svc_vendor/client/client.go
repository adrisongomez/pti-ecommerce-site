// Code generated by goa v3.20.0, DO NOT EDIT.
//
// svc-vendor client HTTP transport
//
// Command:
// $ goa gen github.com/adrisongomez/pti-ecommerce-site/backends/design -o
// ./internal

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the svc-vendor service endpoint HTTP clients.
type Client struct {
	// List Doer is the HTTP client used to make requests to the list endpoint.
	ListDoer goahttp.Doer

	// Create Doer is the HTTP client used to make requests to the create endpoint.
	CreateDoer goahttp.Doer

	// DeleteByID Doer is the HTTP client used to make requests to the deleteById
	// endpoint.
	DeleteByIDDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the svc-vendor service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		ListDoer:            doer,
		CreateDoer:          doer,
		DeleteByIDDoer:      doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// List returns an endpoint that makes HTTP requests to the svc-vendor service
// list server.
func (c *Client) List() goa.Endpoint {
	var (
		encodeRequest  = EncodeListRequest(c.encoder)
		decodeResponse = DecodeListResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildListRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("svc-vendor", "list", err)
		}
		return decodeResponse(resp)
	}
}

// Create returns an endpoint that makes HTTP requests to the svc-vendor
// service create server.
func (c *Client) Create() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateRequest(c.encoder)
		decodeResponse = DecodeCreateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildCreateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("svc-vendor", "create", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteByID returns an endpoint that makes HTTP requests to the svc-vendor
// service deleteById server.
func (c *Client) DeleteByID() goa.Endpoint {
	var (
		decodeResponse = DecodeDeleteByIDResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildDeleteByIDRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteByIDDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("svc-vendor", "deleteById", err)
		}
		return decodeResponse(resp)
	}
}
