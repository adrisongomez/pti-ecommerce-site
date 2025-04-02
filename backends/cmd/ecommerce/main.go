package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/adrisongomez/pti-ecommerce-site/backends/databases/db"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils/auth"
	"github.com/adrisongomez/pti-ecommerce-site/backends/pkg/assistant"
	"github.com/adrisongomez/pti-ecommerce-site/backends/pkg/loggers"
	svc "github.com/adrisongomez/pti-ecommerce-site/backends/pkg/services"
	"go.uber.org/zap"
	goahttp "goa.design/goa/v3/http"
)

var (
	Day                  = time.Hour * 24
	Month                = Day * 24 * 30
	ACCESS_TOKEN_SECRET  = os.Getenv("ACCESS_SECRET")
	REFRESH_TOKEN_SECRET = os.Getenv("REFRESH_SECRET")
)

func main() {
	var (
		host   = os.Getenv("APP_HOST")
		port   = "3030"
		dburl  = os.Getenv("DATABASE_URL")
		logger = loggers.CreateLogger("ecommerce-api")
	)
	zap.ReplaceGlobals(logger)

	client := db.NewClient(db.WithDatasourceURL(dburl))
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
		accessTokenValidator  = auth.NewJWTValidator(&ACCESS_TOKEN_SECRET, client)
		refreshTokenValidator = auth.NewJWTValidator(&REFRESH_TOKEN_SECRET, client)
		passwordHasher        = &auth.PasswordHasher{}
		assistantClient       = assistant.NewAssitantService(client)
	)
	healthcheckSvc := svc.NewHealthcheckService()
	refreshAuthService := svc.NewAuthRefreshService(client, logger, accessTokenGenerator, refreshTokenGenerator, refreshTokenValidator)
	authService := svc.NewAuthService(logger, client, passwordHasher, accessTokenGenerator, refreshTokenGenerator, accessTokenValidator)
	productSvc := svc.NewProductService(client, accessTokenValidator)
	mediaSvc := svc.NewMediaService(client, accessTokenValidator)
	userSvc := svc.NewUserService(client, passwordHasher, accessTokenValidator)
	orderSvc := svc.NewOrderService(client, accessTokenValidator)
	addressSvc := svc.NewAddressService(client, accessTokenValidator)
	chatSvc := svc.NewChatService(client, accessTokenValidator, assistantClient)
	mux := goahttp.NewMuxer()
	svc.MountChatSVC(mux, chatSvc)
	svc.MountAddressSVC(mux, addressSvc)
	svc.MountOrderSVC(mux, orderSvc)
	svc.MountMediaSVC(mux, mediaSvc)
	svc.MountHealtcheckSVC(mux, healthcheckSvc)
	svc.MountProductSVC(mux, productSvc)
	svc.MountUserServiceSVC(mux, userSvc)
	svc.MountAuthSVC(mux, authService)
	svc.MountAuthRefreshSVC(mux, refreshAuthService)
	server := &http.Server{Addr: host + ":" + port, Handler: mux}

	logger.Info(fmt.Sprintf("Starting server on :%s\n", port))
	if err := server.ListenAndServe(); err != nil {
		logger.Fatal("Service shutdown due to error", zap.Any("error", err))
		log.Fatal(err)
	}
}
