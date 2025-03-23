package services

import (
	"context"
	"fmt"

	svcHttp "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/http/svc_healthcheck/server"
	svc "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svc_healthcheck"
	goahttp "goa.design/goa/v3/http"
)

type HealthcheckService struct{}

func (h *HealthcheckService) Check(ctx context.Context) (*svc.HealthcheckResponse, error) {
	Status := "OK"
	return &svc.HealthcheckResponse{
		Status: &Status,
	}, nil
}

func NewHealthcheckService() *HealthcheckService {
	return &HealthcheckService{}
}

func MountHealtcheckSVC(mux goahttp.Muxer, healthSvc *HealthcheckService) {
	req := goahttp.RequestDecoder
	res := goahttp.ResponseEncoder

	endpoints := svc.NewEndpoints(healthSvc)
	handler := svcHttp.New(endpoints, mux, req, res, nil, nil)
	svcHttp.Mount(mux, handler)
	go func() {
		for _, mount := range handler.Mounts {
			fmt.Printf("%q mounted on %s %s\n", mount.Method, mount.Verb, mount.Pattern)
		}
	}()
}
