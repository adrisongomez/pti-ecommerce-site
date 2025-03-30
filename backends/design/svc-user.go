package design

import (
	"github.com/adrisongomez/pti-ecommerce-site/backends/design/types"
	. "goa.design/goa/v3/dsl"
)

var PaginatedUser = types.PaginatedResult("user-list", types.User)

var _ = Service(servicePrefix+"user", func() {
	HTTP(func() {
		Path("/users")
	})
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
	Method("show", func() {
		Payload(func() {
			Attribute("userId", Int)
			Required("userId")
		})
		Result(types.User)
		HTTP(func() {
			GET("/{userId}")
			Param("userId")
			Response(StatusOK)
		})
	})
	Method("create", func() {
		Result(types.User)
		Payload(types.UserCreateInput)
		HTTP(func() {
			POST("")
			Response(StatusCreated)
		})
	})
	Method("update", func() {
		Payload(func() {
			Attribute("payload", types.UserCreateInput)
			Attribute("userId", Int)
			Required("payload", "userId")
		})
		Result(types.User)
		HTTP(func() {
			PUT("/{userId}")
			Param("userId")
			Response(StatusOK)
		})
	})

	Method("delete", func() {
		Payload(func() {
			Attribute("userId", Int)
			Required("userId")
		})
		Result(Boolean)
		HTTP(func() {
			DELETE("/{userId}")
			Param("userId")
			Response(StatusAccepted)
		})

	})
})
