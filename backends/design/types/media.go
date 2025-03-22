package types

import (
	. "goa.design/goa/v3/dsl"
)

var MediaInput = Type("MediaInput", func() {
	Attribute("filename", String)
	Attribute("mimeType", String)
	Attribute("size", Int)
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
		Reference(MediaInput)
		Reference(TypeFooter)

		Required("id", "mediaType", "url")
	})
})
