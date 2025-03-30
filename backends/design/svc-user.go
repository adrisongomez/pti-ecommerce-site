package design

import (
	"github.com/adrisongomez/pti-ecommerce-site/backends/design/securities"
	"github.com/adrisongomez/pti-ecommerce-site/backends/design/types"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils/auth"
	. "goa.design/goa/v3/dsl"
)

var PaginatedUser = types.PaginatedResult("user-list", types.User)

var _ = Service(servicePrefix+"user", func() {
	HTTP(func() {
		Path("/users")
		Response("Unauthorized", StatusUnauthorized)
	})
	Error("ErrNotFound")
	Error("Unauthorized")
	Method("list", func() {
		Security(securities.JWTAuth, func() {
			Scope(auth.UsersReads)
		})
		Payload(func() {
			Attribute("pageSize", Int, "Record per page", func() {
				Minimum(10)
				Maximum(100)
				Default(10)
			})
			Attribute("after", Int, "Start listing after this resource", func() {
				Default(0)
			})
			Token("token")
			Required("token")
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
		Security(securities.JWTAuth, func() {
			Scope(auth.UsersReads)
		})
		Payload(func() {
			Attribute("userId", Int)
			Token("token")
			Required("userId", "token")
		})
		Result(types.User)
		HTTP(func() {
			GET("/{userId}")
			Param("userId")
			Response(StatusOK)
			Response("ErrNotFound", StatusNotFound)
		})
	})
	Method("create", func() {
		Result(types.User)
		Payload(func() {
			Attribute("input", types.UserCreateInput)
			Token("token")
			Required("token", "input")
		})
		Security(securities.JWTAuth, func() {
			Scope(auth.UsersWrite)
		})
		HTTP(func() {
			POST("")
			Body("input")
			Response(StatusCreated)
		})
	})
	Method("update", func() {
		Payload(func() {
			Attribute("payload", types.UserCreateInput)
			Attribute("userId", Int)
			Token("token")
			Required("payload", "userId", "token")
		})
		Security(securities.JWTAuth, func() {
			Scope(auth.UsersWrite)
		})
		Result(types.User)
		HTTP(func() {
			PUT("/{userId}")
			Body("payload")
			Param("userId")
			Response(StatusOK)
		})
	})

	Method("delete", func() {
		Payload(func() {
			Attribute("userId", Int)
			Token("token")

			Required("userId", "token")
		})
		Result(Boolean)
		Security(securities.JWTAuth, func() {
			Scope(auth.UsersWrite)
		})
		HTTP(func() {
			DELETE("/{userId}")
			Param("userId")
			Response(StatusAccepted)
		})

	})
})
