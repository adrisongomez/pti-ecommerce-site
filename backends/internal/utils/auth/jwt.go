package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

type JWTValidator struct {
	Secret *string
}

var (
	UnauthorizedError = fmt.Errorf("Unathorized")
	MalformToken      = fmt.Errorf("Token is malform")
)

func (j *JWTValidator) JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error) {
	claims, err := j.Parse(token)
	if err != nil {
		return nil, err
	}
	if hasRequiredScopes(claims.Scopes, schema.RequiredScopes) {
		return ctx, nil
	}
	return nil, UnauthorizedError
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
			return []byte(*j.Secret), nil
		})

	if err != nil {
		return nil, err
	}
	if !parsedToken.Valid {
		return nil, MalformToken
	}
	return claims, nil
}
