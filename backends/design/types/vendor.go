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
		Attribute("name")

		Reference(TypeFooter)
		Required("name")
	})
})
