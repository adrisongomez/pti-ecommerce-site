package types

import (
	. "goa.design/goa/v3/dsl"
)

var ProductVariantInput = Type("ProductVariantInput", func() {
	Attribute("colorName", String, "Color variant option")
	Attribute("colorHex", String, "Color in HEX value that would be used on the variant picker")
	Attribute("price", Int, "Price on cents")

	Required("colorName", "price")
})

var ProductVariant = ResultType("application/vnd.product-variant+json", func() {
	Description("Definition of product variants")
	Attributes(func() {
		Attribute("id", Int, fieldID)
		Attribute("colorName", String, "Color variant option")
		Attribute("colorHex", String, "Color in HEX value that would be used on the variant picker")
		Attribute("price", Int, "Price on cents")
		Attribute("featureMediaId", Int, "ProductMedia which would be focus when a variant is picked by the user")
		Reference(TypeFooter)
		Required("colorName", "price", "id")
	})
})
