package types

import (
	. "goa.design/goa/v3/dsl"
)

var VendorInput = Type("VendorInput", func() {
	Attribute("name", String)
	Required("name")
})

var Vendor = ResultType("application/vnd.vendor+json", func() {
	Attributes(func() {
		Attribute("id", Int, fieldID)
		Attribute("name")

		Attribute("createdAt", String, "Date of creation", fieldDatetime)
		Attribute("updatedAt", String, "Last update date", fieldDatetime)
		Required("name", "createdAt")
	})
})
