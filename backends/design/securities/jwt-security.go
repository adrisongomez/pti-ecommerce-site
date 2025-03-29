package securities

import (
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils/auth"
	. "goa.design/goa/v3/dsl"
)

var JWTAuth = JWTSecurity("accessToken", func() {
	Description("JWT authentication with scopes")
	for _, scope := range append(auth.AdminScope, auth.AccessTokenGeneration) {
		Scope(string(scope))
	}
})
