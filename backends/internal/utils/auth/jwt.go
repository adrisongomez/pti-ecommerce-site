package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/adrisongomez/pti-ecommerce-site/backends/databases/db"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/auth"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"goa.design/goa/v3/security"
)

type JWTGenerator struct {
	Secret             *string
	ExpirationBandwith time.Duration
}

func (j *JWTGenerator) GenerateToken(userID int, scopes []string) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(j.ExpirationBandwith)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
			Subject:   fmt.Sprintf("%v", userID),
			Audience:  jwt.ClaimStrings{"user"},
			Issuer:    "ecommerce-api",
		},
		UserID: userID,
		Scopes: scopes,
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString([]byte(*j.Secret))
}

type Claims struct {
	UserID int      `json:"uid"`
	Scopes []string `json:"scopes"`
	jwt.RegisteredClaims
}

func (c *Claims) IsExpired() bool {
	return time.Now().After(c.ExpiresAt.Time)
}

type JWTValidator struct {
	secret *string
	client *db.PrismaClient
}

const (
	UserCtxKey   = "user"
	ClaimsCtxKey = "claims"
)

var (
	UnauthorizedError = fmt.Errorf("Unathorized")
	MalformToken      = fmt.Errorf("Token is malform")
)

func (j *JWTValidator) JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error) {
	claim, err := j.Parse(token)
	if err != nil {
		zap.L().Error("Claims is expired", zap.Error(err))
		return nil, err
	}

	if claim.IsExpired() {
		zap.L().Error("Claims is expired", zap.Error(fmt.Errorf("Token is expired")))
		return nil, UnauthorizedError
	}

	if !hasRequiredScopes(claim.Scopes, schema.RequiredScopes) {
		err := fmt.Errorf("Not required scope: %v", schema.RequiredScopes)
		zap.L().Info("Scope checking failed", zap.Error(err))
		return nil, auth.MakeUnauthorized(err)
	}

	userDB, err := j.client.User.FindUnique(db.User.ID.Equals(claim.UserID)).Exec(ctx)

	if err != nil {
		if db.IsErrNotFound(err) {
			return nil, auth.MakeUnauthorized(fmt.Errorf("User does not exists anymore"))
		}
		return nil, err
	}

	ctx = context.WithValue(ctx, UserCtxKey, userDB)
	ctx = context.WithValue(ctx, ClaimsCtxKey, claim)
	return ctx, nil
}

func (j *JWTValidator) Parse(token string) (*Claims, error) {
	claims := &Claims{}
	parsedToken, err := jwt.ParseWithClaims(
		token,
		claims,
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected token method: %v", t.Header["alg"])
			}
			return []byte(*j.secret), nil
		})

	if err != nil {
		return nil, err
	}
	if !parsedToken.Valid {
		return nil, MalformToken
	}
	return claims, nil
}

func NewJWTValidator(secret *string, client *db.PrismaClient) *JWTValidator {
	return &JWTValidator{secret, client}
}
