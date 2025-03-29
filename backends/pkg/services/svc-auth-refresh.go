package services

import (
	"context"
	"fmt"

	"github.com/adrisongomez/pti-ecommerce-site/backends/databases/db"
	. "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/auth_refresh"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/http/auth_refresh/server"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils/auth"
	"go.uber.org/zap"
	"goa.design/goa/v3/http"
)

type AuthRefreshService struct {
	client                *db.PrismaClient
	accessTokenGenerator  *auth.JWTGenerator
	refreshTokenGenerator *auth.JWTGenerator

	*auth.JWTValidator
	*zap.Logger
}

func NewAuthRefreshService(
	client *db.PrismaClient,
	logger *zap.Logger,
	accessTokenGenerator, refreshTokenGenerator *auth.JWTGenerator,
	validator *auth.JWTValidator,
) Service {
	return &AuthRefreshService{
		client,
		accessTokenGenerator,
		refreshTokenGenerator,
		validator,
		logger,
	}
}

func (a *AuthRefreshService) Refresh(ctx context.Context, input *RefreshPayload) (res *Creds, err error) {
	return nil, nil
}

func MountAuthRefreshSVC(mux http.Muxer, svc Service) {
	endpoints := NewEndpoints(svc)
	req := http.RequestDecoder
	res := http.ResponseEncoder

	handler := server.New(endpoints, mux, req, res, nil, nil)
	server.Mount(mux, handler)

	go func() {
		for _, mount := range handler.Mounts {
			zap.L().Info(fmt.Sprintf("%q mounted on %s %s\n", mount.Method, mount.Verb, mount.Pattern))
		}
	}()
}
