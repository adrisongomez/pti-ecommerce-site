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

	Error("ErrNotFound")
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
			Response(StatusOK)
			Response("ErrNotFound", StatusNotFound)
		})
	})

	Method("createProduct", func() {
		Description("Create a new product")
		Payload(types.ProductCreateInput)
		Result(types.Product)
		HTTP(func() {
			POST("")
			Response(StatusCreated)
			Response("Conflict", StatusConflict)
		})
	})

	Method("updateProductById", func() {
		Description("Update a product by id")

		Payload(func() {
			Attribute("productId", Int, "Unique product identifier")
			Attribute("payload", types.ProductUpdateInput)
			Required("productId")
		})
		Result(types.Product)

		HTTP(func() {
			PUT("/{productId}")
			Param("productId")
			Response(StatusOK)
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
			DELETE("/{productId}")
			Param("productId")
			Response(StatusOK)
			Response("ErrNotFound", StatusNotFound)
		})
	})

	Method("addVariant", func() {
		Description("Add a new product variant to a given product")
		Payload(func() {
			Attribute("productId", Int, "Unique product identifier")
			Required("productId")

			Attribute("payload", types.ProductVariantInput)
		})
		Result(types.Product)
		HTTP(func() {
			POST("/{productId}/variants")
			Param("productId")
			Response(StatusCreated)
			Response("ErrNotFound", StatusNotFound)
		})
	})

	Method("removeVariant", func() {
		Description("Remove a given product variant")
		Payload(func() {
			Attribute("productId", Int, "Product ID")
			Attribute("variantId", Int, "Product ID")
			Required("productId", "variantId")
		})
		Result(types.Product)

		HTTP(func() {
			DELETE("/{productId}/variants/{variantId}")
			Param("productId")
			Param("variantId")
			Response(StatusOK)
			Response("ErrNotFound", StatusNotFound)
		})
	})

	Method("addMedia", func() {
		Description("Add a new product media to a given product")
		Payload(func() {
			Attribute("productId", Int, "Unique product identifier")
			Attribute("payload", types.ProductMediaInput)

			Required("productId")
		})
		Result(types.Product)
		HTTP(func() {
			POST("/{productId}/product-medias")
			Param("productId")
			Response(StatusCreated)
			Response("BadRequest", StatusBadRequest)
		})
	})

	Method("removeMedia", func() {
		Description("Remove a product media from a given product")
		Payload(func() {
			Attribute("productId", Int, "Unique product identifier")
			Attribute("productMediaId", Int)
			Required("productId", "productMediaId")
		})
		Result(types.Product)
		HTTP(func() {
			POST("/{productId}/product-medias/{productMediaId}")
			Param("productId")
			Param("productMediaId")
			Response(StatusOK)
			Response("ErrNotFound", StatusNotFound)
		})
	})

})
