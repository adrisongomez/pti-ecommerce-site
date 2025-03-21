package types

import (
	. "goa.design/goa/v3/dsl"
)

var ProductMediaInput = Type("ProductMediaInput", func() {
	Attribute("media_id", Int, "Media Id", Required, fieldID)
	Attribute("alt", String, "Alt for media")
})

var ProductMedia = ResultType("application/vnd.product-media+json", func() {
	Reference(ProductMediaInput)
	Attributes(func() {
		Attribute("url", String, "URL for the media")
	})
})
