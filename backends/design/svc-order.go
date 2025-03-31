package design

import (
	"github.com/adrisongomez/pti-ecommerce-site/backends/design/types"
	. "goa.design/goa/v3/dsl"
)

var PaginatedOrder = types.PaginatedResult("order-list", types.Order)

var _ = Service("order", func() {
	Error("NotFound")
	Error("BadInput")
	Error("Unauthorized")

	HTTP(func() {
		Path("/orders")
		Response("Unauthorized", StatusUnauthorized)
	})
	Method("create", func() {
		Payload(func() {
			Attribute("input", types.OrderInput)
			Required("input")
		})

		Result(types.Order)
		HTTP(func() {
			POST("")
			Body("input")
			Response(StatusCreated)
			Response("BadInput", StatusBadRequest)
		})
	})
	Method("Cancel", func() {
		Payload(func() {
			Attribute("orderId", Int)
			Required("orderId")
		})
		Result(Boolean)
		HTTP(func() {
			DELETE("/{orderId}")
			Param("orderId")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})
	Method("list", func() {
		Payload(func() {
			Attribute("pageSize", Int, "Record per page", func() {
				Minimum(10)
				Maximum(100)
				Default(10)
			})
			Attribute("q", String)
			Attribute("after", Int, "Start listing after this resource", func() {
				Default(0)
			})
		})
		Result(PaginatedOrder)
		HTTP(func() {
			GET("")
			Param("pageSize")
			Param("after")
			Response(StatusOK)
			Response("BadInput", StatusBadRequest)
		})
	})
	Method("show", func() {
		Payload(func() {
			Attribute("orderId", Int)
			Required("orderId")
		})
		Result(types.Order)
		HTTP(func() {
			GET("/{orderId}")
			Response(StatusOK)
		})
	})
})
