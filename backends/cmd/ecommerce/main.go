package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/adrisongomez/pti-ecommerce-site/backends/databases/db"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils/auth"
	"github.com/adrisongomez/pti-ecommerce-site/backends/pkg/loggers"
	svc "github.com/adrisongomez/pti-ecommerce-site/backends/pkg/services"
	"go.uber.org/zap"
	goahttp "goa.design/goa/v3/http"
)

var (
	Day                  = time.Hour * 24
	Month                = Day * 24 * 30
	ACCESS_TOKEN_SECRET  = "SECRET"
	REFRESH_TOKEN_SECRET = "SECRET2"
)

func main() {
	var (
		port   = "3030"
		logger = loggers.CreateLogger("ecommerce-api")
		client = db.NewClient()
	)

	zap.ReplaceGlobals(logger)

	if err := client.Prisma.Connect(); err != nil {
		logger.Error("Error connecting to prisma server", zap.Error(err))
		panic(err)
	}

	defer func() {
		logger.Sync()
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	var (
		accessTokenGenerator  = &auth.JWTGenerator{Secret: &ACCESS_TOKEN_SECRET, ExpirationBandwith: Day}
		refreshTokenGenerator = &auth.JWTGenerator{Secret: &REFRESH_TOKEN_SECRET, ExpirationBandwith: Month}
		accessTokenValidator  = &auth.JWTValidator{Secret: &ACCESS_TOKEN_SECRET}
		refreshTokenValidator = &auth.JWTValidator{Secret: &REFRESH_TOKEN_SECRET}
		passwordHasher        = &auth.PasswordHasher{}
	)

	refreshAuthService := svc.NewAuthRefreshService(client, logger, accessTokenGenerator, refreshTokenGenerator, refreshTokenValidator)
	authService := svc.NewAuthService(logger, client, passwordHasher, accessTokenGenerator, refreshTokenGenerator, accessTokenValidator)
	healthcheckSvc := svc.NewHealthcheckService()
	productSvc := svc.NewProductService(client)
	mediaSvc := svc.NewMediaService(client)
	userSvc := svc.NewUserService(client, passwordHasher)
	mux := goahttp.NewMuxer()
	svc.MountMediaSVC(mux, mediaSvc)
	svc.MountHealtcheckSVC(mux, healthcheckSvc)
	svc.MountProductSVC(mux, productSvc)
	svc.MountUserServiceSVC(mux, userSvc)
	svc.MountAuthSVC(mux, authService)
	svc.MountAuthRefreshSVC(mux, refreshAuthService)
	server := &http.Server{Addr: ":" + port, Handler: mux}

	logger.Info(fmt.Sprintf("Starting server on :%s\n", port))
	if err := server.ListenAndServe(); err != nil {
		logger.Fatal("Service shutdown due to error", zap.Any("error", err))
		log.Fatal(err)
	}
}
