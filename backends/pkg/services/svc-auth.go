package services

import (
	"context"
	"fmt"

	"github.com/adrisongomez/pti-ecommerce-site/backends/databases/db"
	. "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/auth"
	authhttp "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/http/auth/server"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils"
	auth "github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils/auth"
	"go.uber.org/zap"
	"goa.design/goa/v3/http"
	"goa.design/goa/v3/security"
)

type AuthService struct {
	client                *db.PrismaClient
	hasher                *auth.PasswordHasher
	accessTokenGenerator  *auth.JWTGenerator
	refreshTokenGenerator *auth.JWTGenerator

	*zap.Logger
	*auth.JWTValidator
}

var (
	InvalidCredential   = fmt.Errorf("Credentials")
	InternalServerError = fmt.Errorf("Internal Server Error")
	CantParseToken      = fmt.Errorf("Cant parse token")
)

func NewAuthService(
	logger *zap.Logger,
	client *db.PrismaClient,
	hasher *auth.PasswordHasher,
	accessTokenGenerator, refreshTokenGenerator *auth.JWTGenerator,
	tokenValidator *auth.JWTValidator,
) Service {
	return &AuthService{client, hasher, accessTokenGenerator, refreshTokenGenerator, logger, tokenValidator}
}

func (a *AuthService) BasicAuth(ctx context.Context, user, pass string, schema *security.BasicScheme) (context.Context, error) {
	a.Info("Validating basic auth", zap.String("user", user))
	userDB, err := a.client.User.FindUnique(db.User.Email.Equals(user)).Exec(ctx)
	if err != nil {
		a.Error("Error on trying to get user from db", zap.Error(err))
		if db.IsErrNotFound(err) {
			return nil, InvalidCredential
		}
		return nil, InternalServerError
	}
	if !a.hasher.Validate(pass, userDB.PasswordHash) {
		a.Error("Password not match")
		return nil, InvalidCredential
	}
	return ctx, nil
}

func (a AuthService) MapUserDBToOutput(model db.UserModel) *User {
	user := User{
		ID:        model.ID,
		FirstName: model.FirstName,
		LastName:  nil,
		Email:     model.Email,
		Role:      UserRole(model.Role),
		CreatedAt: model.CreatedAt.String(),
	}
	if value, ok := model.LastName(); ok {
		user.LastName = &value
	}

	if value, ok := model.UpdatedAt(); ok {
		user.UpdatedAt = utils.StringRef(value.String())
	}
	return &user
}

func (a *AuthService) Me(ctx context.Context, input *MePayload) (*User, error) {
	a.Info("current user", zap.Any("ctx", ctx.Value(auth.UserCtxKey)))
	if value, ok := ctx.Value(auth.UserCtxKey).(*db.UserModel); ok {
		return a.MapUserDBToOutput(*value), nil
	}
	return nil, MakeNotFound(fmt.Errorf("No current user"))
}

func (a *AuthService) Signup(ctx context.Context, input *UserRegistrationInput) (res *Creds, err error) {
	a.Info("Signup got called with", zap.String("email", input.FirstName), zap.String("firstName", input.FirstName))
	hashedPassword, err := a.hasher.Hash(input.Password)
	if err != nil {
		return nil, err
	}
	change := []db.UserSetParam{}

	if input.LastName != nil {
		change = append(
			change,
			db.User.LastName.Set(*input.LastName),
		)
	}
	user, err := a.client.User.CreateOne(
		db.User.FirstName.Set(input.FirstName),
		db.User.Email.Set(input.Email),
		db.User.PasswordHash.Set(hashedPassword),
		change...,
	).Exec(ctx)

	if err != nil {

		if field, err := db.IsErrUniqueConstraint(err); err {
			err := fmt.Errorf("Error on user trying to violate unique constraint: %v", field)
			a.Error("Error on user trying to violate unique constraint", zap.Error(err))
			return nil, err
		}
		a.Error("Error trying to create user from db", zap.Error(err))
		return nil, err
	}
	accessToken, err := a.accessTokenGenerator.GenerateToken(user.ID, auth.GetScopesByRole(user.Role))
	if err != nil {
		a.Error("Error generate accessToken", zap.Error(err))
		return nil, err
	}
	refreshToken, err := a.refreshTokenGenerator.GenerateToken(user.ID, auth.RefreshTokenScope)
	if err != nil {
		a.Error("Error generate refreshToken", zap.Error(err))
		return nil, err
	}
	creds := Creds{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return &creds, nil
}

func (a *AuthService) Login(ctx context.Context, payload *LoginPayload) (*Creds, error) {
	a.Info("Login got called with", zap.String("email", payload.Email))
	userdb, err := a.client.User.FindUnique(db.User.Email.Equals(payload.Email)).Exec(ctx)
	if err != nil {
		a.Error("Error getting user from db", zap.Error(err))
		return nil, MakeNotValidCrendentials(fmt.Errorf("Credentials not valid"))
	}
	if !a.hasher.Validate(payload.Password, userdb.PasswordHash) {
		a.Info("Error password are not valid")
		return nil, MakeNotValidCrendentials(fmt.Errorf("Credentials not valid"))
	}
	accessToken, err := a.accessTokenGenerator.GenerateToken(userdb.ID, auth.GetScopesByRole(userdb.Role))
	if err != nil {
		a.Error("Error generate accessToken", zap.Error(err))
		return nil, err
	}
	refreshToken, err := a.refreshTokenGenerator.GenerateToken(userdb.ID, auth.RefreshTokenScope)
	if err != nil {
		a.Error("Error generate refreshToken", zap.Error(err))
		return nil, err
	}
	return &Creds{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func MountAuthSVC(mux http.Muxer, svc Service) {
	endpoints := NewEndpoints(svc)
	req := http.RequestDecoder
	res := http.ResponseEncoder

	handler := authhttp.New(endpoints, mux, req, res, nil, nil)
	authhttp.Mount(mux, handler)

	go func() {
		for _, mount := range handler.Mounts {
			zap.L().Info(fmt.Sprintf("%q mounted on %s %s\n", mount.Method, mount.Verb, mount.Pattern))
		}
	}()
}
