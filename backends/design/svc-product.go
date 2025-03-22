package design

import (
	"github.com/adrisongomez/pti-ecommerce-site/backends/design/types"
	. "goa.design/goa/v3/dsl"
)

var _ = Service(servicePrefix+"-products", func() {
	Description("The product service perform CRUD operation over the product resource")

	HTTP(func() {
		Path("/products")
	})

	Error("NotFound")
	Error("BadRequest")
	Error("Conflict")

	Method("listProduct", func() {
		Description("List products")
		Payload(func() {
			Attribute("pageSize", Int, "Record per page", func() {
				Minimum(10)
				Maximum(100)
				Default(10)
			})
			Attribute("after", Int, "Start listing after this resource")
		})

		Result(types.ProductPaginated)

		HTTP(func() {
			GET("")
			Param("pageSize")
			Param("after")
			Response(StatusOK)
			Response("BadRequest", StatusBadRequest)
		})
	})

	Method("getProductById", func() {
		Description("Get a product by its id")

		Payload(func() {
			Attribute("productId", Int, "Unique product identifier")
			Required("productId")
		})
		Result(types.Product)

		HTTP(func() {
			GET("/{productId}")
			Param("productId")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})

	Method("createProduct", func() {
		Description("Create a new product")

		Payload(types.ProductInput)
		Result(types.Product)

		HTTP(func() {
			POST("")
			Response(StatusCreated)
			Response("Conflict", StatusConflict)
		})
	})

	Method("updateProductById", func() {
		Description("Create a new product")

		Payload(func() {
			Attribute("productId", Int, "Unique product identifier")
			Attribute("payload", types.ProductInput)
			Required("productId")
		})
		Result(types.Product)

		HTTP(func() {
			PUT("/{productId}")
			Param("productId")
			Body("payload")
			Response(StatusOK)
			Response("Conflict", StatusConflict)
		})
	})

	Method("deleteProductById", func() {
		Description("Create a new product")

		Payload(func() {
			Attribute("productId", Int, "Unique product identifier")
			Required("productId")
		})
		Result(Boolean)

		HTTP(func() {
			PUT("/{productId}")
			Param("productId")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})
})
