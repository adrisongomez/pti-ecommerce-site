package design

import (
	"github.com/adrisongomez/pti-ecommerce-site/backends/design/securities"
	"github.com/adrisongomez/pti-ecommerce-site/backends/design/types"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils/auth"
	. "goa.design/goa/v3/dsl"
)

var PaginatedAddress = types.PaginatedResult("address-list", types.Address)

var _ = Service("address", func() {
	Error("Unauthorized")
	Error("NotFound")
	HTTP(func() {
		Path("/addresses")
		Response("Unauthorized", StatusUnauthorized)
	})
	Method("create", func() {
		Security(securities.JWTAuth, func() {
			Scope(auth.OrderWrite)
		})
		Payload(func() {
			Attribute("input", types.AddressInput)
			Token("token")
			Required("input", "token")
		})
		Result(types.Address)
		HTTP(func() {
			POST("")
			Body("input")
			Response(StatusCreated)
		})
	})
	Method("Delete", func() {
		Result(Boolean)
		Security(securities.JWTAuth, func() {
			Scope(auth.OrderWrite)
		})
		Payload(func() {
			Token("token")
			Attribute("addressId", Int)
			Required("addressId", "token")
		})
		HTTP(func() {
			DELETE("/{addressId}")
			Param("addressId")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})
	Method("Show", func() {
		Result(types.Address)
		Security(securities.JWTAuth, func() {
			Scope(auth.OrderRead)
		})
		Payload(func() {
			Token("token")
			Attribute("addressId", Int)
			Required("addressId", "token")
		})
		HTTP(func() {
			GET("/{addressId}")
			Param("addressId")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})
	Method("List", func() {
		Result(PaginatedAddress)
		Security(securities.JWTAuth, func() {
			Scope(auth.OrderRead)
		})
		Payload(func() {
			Token("token")
			Attribute("pageSize", Int, "Record per page", func() {
				Minimum(10)
				Maximum(100)
				Default(10)
			})
			Attribute("after", Int, "Start listing after this resource", func() {
				Default(0)
			})
			Required("token")
		})
		HTTP(func() {
			GET("")
			Param("pageSize")
			Param("after")
			Response(StatusOK)
		})
	})
})
