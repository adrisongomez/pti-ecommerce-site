package types

import (
	. "goa.design/goa/v3/dsl"
)

var ProductVariantInput = Type("ProductVariant", func () {
	Attribute("name", String, "Variant name", String)
	Attribute("price", Int, "Price on cent of a product variant", Required)
	Attribute("media_id", Int)
})


