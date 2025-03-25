package types

import (
	. "goa.design/goa/v3/dsl"
)

var ProductMediaInput = Type("ProductMediaInput", func() {
	Attribute("mediaId", Int, "ID of the media record where the resource has being upload")
	Attribute("sortNumber", Int, "Position on the images of the product")
	Attribute("alt", String, "Alt text that would show in case the image does not render")

	Required("mediaId", "sortNumber")
})

var ProductMedia = ResultType("application/vnd.product-media+json", func() {
	Description("Instance of media in a product")
	ContentType("application/json")
	Attributes(func() {
		Attribute("id", Int, fieldID)
		Attribute("url", String, "URL to the media")
		Attribute("mediaType", MediaType)
		Attribute("mediaId", Int, "ID of the media record where the resource has being upload")
		Attribute("sortNumber", Int, "Position on the images of the product")
		Attribute("alt", String, "Alt text that would show in case the image does not render")
		Attribute("createdAt", String, "Date of creation", fieldDatetime)
		Attribute("updatedAt", String, "Last update date", fieldDatetime)

		Required(
			"id",
			"url",
			"mediaType",
			"mediaId",
			"sortNumber",
			"createdAt",
		)
	})
})

var MediaType = Type("MediaType", String, func() {
	Description("Type of the media")
	Enum("IMAGE", "VIDEO", "UNKNWON")
})
