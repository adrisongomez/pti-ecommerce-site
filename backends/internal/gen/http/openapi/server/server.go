// Code generated by goa v3.20.0, DO NOT EDIT.
//
// openapi HTTP server
//
// Command:
// $ goa gen github.com/adrisongomez/pti-ecommerce-site/backends/design -o
// ./internal

package server

import (
	"context"
	"net/http"
	"path"

	openapi "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/openapi"
	goahttp "goa.design/goa/v3/http"
)

// Server lists the openapi service endpoint HTTP handlers.
type Server struct {
	Mounts                     []*MountPoint
	InternalGenHTTPOpenapiJSON http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the openapi service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *openapi.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
	fileSystemInternalGenHTTPOpenapiJSON http.FileSystem,
) *Server {
	if fileSystemInternalGenHTTPOpenapiJSON == nil {
		fileSystemInternalGenHTTPOpenapiJSON = http.Dir(".")
	}
	fileSystemInternalGenHTTPOpenapiJSON = appendPrefix(fileSystemInternalGenHTTPOpenapiJSON, "/internal/gen/http")
	return &Server{
		Mounts: []*MountPoint{
			{"Serve internal/gen/http/openapi.json", "GET", "/api/openapi.json"},
		},
		InternalGenHTTPOpenapiJSON: http.FileServer(fileSystemInternalGenHTTPOpenapiJSON),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "openapi" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return openapi.MethodNames[:] }

// Mount configures the mux to serve the openapi endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountInternalGenHTTPOpenapiJSON(mux, http.StripPrefix("/api", h.InternalGenHTTPOpenapiJSON))
}

// Mount configures the mux to serve the openapi endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// appendFS is a custom implementation of fs.FS that appends a specified prefix
// to the file paths before delegating the Open call to the underlying fs.FS.
type appendFS struct {
	prefix string
	fs     http.FileSystem
}

// Open opens the named file, appending the prefix to the file path before
// passing it to the underlying fs.FS.
func (s appendFS) Open(name string) (http.File, error) {
	switch name {
	}
	return s.fs.Open(path.Join(s.prefix, name))
}

// appendPrefix returns a new fs.FS that appends the specified prefix to file paths
// before delegating to the provided embed.FS.
func appendPrefix(fsys http.FileSystem, prefix string) http.FileSystem {
	return appendFS{prefix: prefix, fs: fsys}
}

// MountInternalGenHTTPOpenapiJSON configures the mux to serve GET request made
// to "/api/openapi.json".
func MountInternalGenHTTPOpenapiJSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/api/openapi.json", h.ServeHTTP)
}
