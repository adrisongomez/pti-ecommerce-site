package design

import (
	"github.com/adrisongomez/pti-ecommerce-site/backends/design/securities"
	"github.com/adrisongomez/pti-ecommerce-site/backends/design/types"
	. "goa.design/goa/v3/dsl"
)

var Creds = Type("Creds", func() {
	Attribute("accessToken", String, "Access JWT Token", func() {
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
	})
	Attribute("refreshToken", String, "Refresh JWT Token", func() {
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
	})
	Required("accessToken", "refreshToken")
})

var BasicAuth = BasicAuthSecurity("basic", func() {
	Scope("token:generation")
})

var _ = Service("auth", func() {
	HTTP(func() { Path("/auth") })
	Error("NotValidCrendentials")
	Error("BadInput")
	Error("Unproccesable")
	Method("login", func() {
		Security(BasicAuth)
		Payload(func() {
			Username("email", String, func() {
				Example("test@example.com")
			})
			Password("password", String, func() {
				Example("password")
			})
			Required("email", "password")
		})
		Result(Creds)
		HTTP(func() {
			POST("/login")
			Response(StatusOK)
			Response("NotValidCrendentials", StatusUnauthorized)
		})
	})
	Method("signup", func() {
		Payload(types.UserRegistrationInput)
		HTTP(func() {
			POST("/signup")
			Response(StatusOK)
			Response("BadInput", StatusBadRequest)
		})
		Result(Creds)
	})
	Method("me", func() {
		Security(securities.JWTAuth)
		Payload(func() {
			Token("token")
			Required("token")
		})
		Result(types.User)
		HTTP(func() {
			GET("/me")
			Response(StatusOK)
			Response("NotValidCrendentials", StatusUnauthorized)
			Response("Unproccesable", StatusUnprocessableEntity)
		})
	})
})
