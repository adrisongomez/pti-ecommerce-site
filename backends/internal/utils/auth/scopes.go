package auth

import (
	"github.com/adrisongomez/pti-ecommerce-site/backends/databases/db"
	"go.uber.org/zap"
)

type Scope = string

const (
	// RefreshToke Access Scopen
	AccessTokenGeneration Scope = "access_token:generation"

	// Customer access scope
	UserMe           Scope = "user:read_personal"
	UserMeWrite      Scope = "user:write_personal"
	OrderWrite       Scope = "order:write"
	OrderRead        Scope = "order:read"
	ChatSessionWrite Scope = "chat_session:write"
	ChatSessionRead  Scope = "chat_session:read"

	// Admin access scope
	ProductsWrite Scope = "products:write"
	UsersReads    Scope = "users:reads"
	UsersWrite    Scope = "users:write"
	MediasReads   Scope = "medias:reads"
	MediasWrite   Scope = "medias:write"
)

var RefreshTokenScope = []string{
	AccessTokenGeneration,
}

var CustomerScope = []string{
	UserMe,
	UserMeWrite,
	OrderWrite,
	OrderRead,
	ChatSessionWrite,
	ChatSessionRead,
}

var AdminScope = append([]string{
	ProductsWrite,
	UsersReads,
	UsersWrite,
	MediasReads,
	MediasWrite,
}, CustomerScope...)

func GetScopesByRole(role db.UserRole) []string {
	zap.L().Info("Getting scope for user role", zap.String("role", string(role)))
	switch role {
	case db.UserRoleCustomer:
		return CustomerScope
	case db.UserRoleAdmin:
		return AdminScope
	}
	return []string{}
}

func hasRequiredScopes(tokenScopes, requiredScopes []string) bool {
	zap.L().Info("Validating scopes", zap.Any("token", tokenScopes), zap.Any("requireds", requiredScopes))
	scopeMap := make(map[string]bool)
	for _, scope := range tokenScopes {
		scopeMap[scope] = true
	}

	for _, required := range requiredScopes {
		if !scopeMap[required] {
			return false
		}
	}
	return true
}
