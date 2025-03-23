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
	Extend(TypeFooter)
	Attributes(func() {
		Attribute("id", Int, fieldID)
		Attribute("mediaType", MediaType)
		Attribute("url", String)
		Attribute("filename", String)
		Attribute("mimeType", String)
		Attribute("size", Int64)
		Attribute("bucket", String)
		Attribute("key", String)
		Required("id", "mediaType", "url")
	})
})
