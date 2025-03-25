package types

import (
	. "goa.design/goa/v3/dsl"
)

var MediaInput = Type("MediaInput", func() {
	Attribute("filename", String)
	Attribute("mimeType", String)
	Attribute("size", Int64)
	Attribute("bucket", String)
	Attribute("key", String)

	Required("filename", "mimeType", "size", "bucket", "key")
})

var Media = ResultType("application/vnd.media+json", func() {
	ContentType("application/json")
	Attributes(func() {
		Attribute("id", Int, fieldID)
		Attribute("mediaType", MediaType)
		Attribute("url", String)
		Attribute("filename", String)
		Attribute("mimeType", String)
		Attribute("size", Int64)
		Attribute("bucket", String)
		Attribute("key", String)
		Attribute("createdAt", String, "Date of creation", fieldDatetime)
		Attribute("updatedAt", String, "Last update date", fieldDatetime)

		Required(
			"id",
			"mediaType",
			"url",
			"createdAt",
			"bucket",
			"filename",
			"mimeType",
			"size",
			"key",
		)
	})
})
