package design

import (
	"github.com/adrisongomez/pti-ecommerce-site/backends/design/securities"
	. "goa.design/goa/v3/dsl"
)

var _ = Service("authRefresh", func() {
	HTTP(func() { Path("/auth") })
	Error("NotValidToken")
	Security(securities.JWTAuth)
	Method("refresh", func() {
		Security(securities.JWTAuth)
		Payload(func() {
			Token("token")
			Required("token")
		})
		HTTP(func() {
			POST("/refresh")
			Response(StatusOK)
			Response("NotValidToken", StatusConflict)
		})
		Result(Creds)
	})
})
