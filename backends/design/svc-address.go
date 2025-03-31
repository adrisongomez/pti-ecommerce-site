package design

import (
	"github.com/adrisongomez/pti-ecommerce-site/backends/design/types"
	. "goa.design/goa/v3/dsl"
)

var _ = Service("address", func() {
	Error("NotFound")
	Error("BadInput")
	Error("Unauthorized")

	HTTP(func() {
		Path("/addresses")
		Response("NotFound", StatusNotFound)
	})

	Method("show", func() {
		Payload(func() {
			Attribute("addressId", Int)
		})
		Result(types.Address)
		HTTP(func() {
			GET("/{addressId}")
			Param("addressId")
			Response(StatusOK)
			Response("NotFound", StatusOK)
		})
	})
	Method("upsert", func() {
		Payload(func() {
			Attribute("addressId", Int)
			Required("addressId")
		})
		Result(types.Address)
		HTTP(func() {
			POST("/{addressId}")
			Param("addressId")
			Response(StatusCreated)
			Response("BadInput", StatusOK)
		})
	})
})
