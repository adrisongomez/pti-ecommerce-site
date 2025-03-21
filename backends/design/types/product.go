package types

import (
	. "goa.design/goa/v3/dsl"
)

var ProductInput = Type("ProductInput", func() {
	Attribute("title", String, "Title", Required)
	Attribute("handle", String, "Handle", Required)
	Attribute("description", String, "Product description", Required)
	Attribute("medias", ArrayOf(ProductMediaInput), "Product media input", Required)
})

var Product = ResultType("application/vnd.product+json", func() {
	Description("Product information")
	ContentType("application/json")
	Reference(TypeFooter)

	Attributes(func() {
		Attribute("title", String, "Title", Required)
		Attribute("handle", String, "Handle", Required)
		Attribute("description", String, "Product description", Required)
		Attribute("id", fieldID)
		Attribute("medias", ArrayOf(ProductMedia), "Product medias")
	})
})
