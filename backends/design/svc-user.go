package design

import (
	"github.com/adrisongomez/pti-ecommerce-site/backends/design/types"
	. "goa.design/goa/v3/dsl"
)

var PaginatedUser = types.PaginatedResult("user-list", types.User)

var _ = Service(servicePrefix+"user", func() {
	HTTP(func() { Path("/users") })
	Method("list", func() {
		Payload(func() {
			Attribute("pageSize", Int, "Record per page", func() {
				Minimum(10)
				Maximum(100)
				Default(10)
			})
			Attribute("after", Int, "Start listing after this resource", func() {
				Default(0)
			})
		})
		Result(PaginatedUser)
		HTTP(func() {
			GET("")
			Param("pageSize")
			Param("after")
			Response(StatusOK)
		})
	})
	Method("upsert", func() {

	})
})
