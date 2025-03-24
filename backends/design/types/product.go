package types

import (
	. "goa.design/goa/v3/dsl"
)

var ProductUpdateInput = Type("ProductUpdateInput", func() {
	Attribute("title", String, "Title's product")
	Attribute("description", String, "Product description")
	Attribute("handle", String, "Last part of the url which use to idepntify the user")
	Attribute("status", ProductStatus, "Product's status", func() {
		Default("DRAFT")
	})
	Attribute("tags", ArrayOf(String), "Product tags")
	Attribute("vendorId", Int, "Vendor's product")
	Required("title", "description", "tags", "vendorId")
})

var ProductCreateInput = Type("ProductInput", func() {
	Attribute("title", String, "Title's product")
	Attribute("description", String, "Product description")
	Attribute("handle", String, "Last part of the url which use to idepntify the user")
	Attribute("status", ProductStatus, "Product's status", func() {
		Default("DRAFT")
	})
	Attribute("tags", ArrayOf(String), "Product tags")
	Attribute("vendorId", Int, "Vendor's product")
	Attribute("variants", ArrayOf(ProductVariantInput), "Product variants")
	Attribute("medias", ArrayOf(ProductMediaInput))

	Required("title", "description", "tags", "vendorId", "variants")
})

var ProductStatus = Type("ProductStatus", String, func() {
	Description("Define the status of product on the site")
	Enum("ACTIVE", "DRAFT")
})

var Product = ResultType("application/vnd.product+json", func() {
	Description("Product information")
	ContentType("application/json")
	Extend(TypeFooter)

	Attributes(func() {
		Attribute("id", Int, fieldID)
		Attribute("title", String, "Title")
		Attribute("description", String, "Product description")
		Attribute("handle", String, "Handle")
		Attribute("vendor", Vendor)
		Attribute("tags", ArrayOf(String), "Product tags")
		Attribute("status", ProductStatus, "The product's status on ecommerce site", func() {
			Default("DRAFT")
		})
		Attribute("variants", ArrayOf(ProductVariant))
		Attribute("medias", ArrayOf(ProductMedia))
		Required("id", "title", "handle", "description", "status", "variants", "medias")
	})
})

var ProductPaginated = PaginatedResult("products-list", Product)
