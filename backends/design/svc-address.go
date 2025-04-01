package design

import (
	"github.com/adrisongomez/pti-ecommerce-site/backends/design/types"
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
		Payload(func() {
			Attribute("input", types.AddressInput)
			Required("input")
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
		Payload(func() {
			Attribute("addressId", Int)
			Required("addressId")
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
		Payload(func() {
			Attribute("addressId", Int)
			Required("addressId")
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
		HTTP(func() {
			GET("")
			Param("pageSize")
			Param("after")
			Response(StatusOK)
		})
	})
})
