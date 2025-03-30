package design

import (
	"github.com/adrisongomez/pti-ecommerce-site/backends/design/securities"
	"github.com/adrisongomez/pti-ecommerce-site/backends/design/types"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils/auth"
	. "goa.design/goa/v3/dsl"
)

var _ = Service(servicePrefix+"-products", func() {
	Description("The product service perform CRUD operation over the product resource")

	HTTP(func() {
		Path("/products")
		Response("Unauthorized", StatusUnauthorized)
	})

	Error("ErrNotFound")
	Error("Unauthorized")
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
			Attribute("after", Int, "Start listing after this resource", func() {
				Default(0)
			})
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
			Response("ErrNotFound", StatusNotFound)
			Response(StatusOK)
		})
	})

	Method("createProduct", func() {
		Security(securities.JWTAuth, func() {
			Scope(auth.ProductsWrite)
		})
		Description("Create a new product")
		Payload(func() {
			Token("token")
			Attribute("input", types.ProductCreateInput)
			Required("token", "input")
		})
		Result(types.Product)
		HTTP(func() {
			POST("")
			Body("input")
			Response(StatusCreated)
			Response("Conflict", StatusBadRequest)
		})
	})

	Method("updateProductById", func() {
		Description("Update a product by id")

		Security(securities.JWTAuth, func() {
			Scope(auth.ProductsWrite)
		})
		Payload(func() {
			Attribute("productId", Int, "Unique product identifier")
			Attribute("payload", types.ProductUpdateInput)
			Token("token")
			Required("productId", "payload", "token")
		})
		Result(types.Product)

		HTTP(func() {
			Body("payload")
			PUT("/{productId}")
			Param("productId")
			Response("ErrNotFound", StatusNotFound)
			Response(StatusOK)
		})
	})

	Method("deleteProductById", func() {
		Description("Delete a product")

		Security(securities.JWTAuth, func() {
			Scope(auth.ProductsWrite)
		})
		Payload(func() {
			Token("token")
			Attribute("productId", Int, "Unique product identifier")
			Required("productId", "token")
		})
		Result(Boolean)

		HTTP(func() {
			DELETE("/{productId}")
			Param("productId")
			Response("ErrNotFound", StatusNotFound)
			Response(StatusOK)
		})
	})
})
