package types

import (
	. "goa.design/goa/v3/dsl"
)

var fieldID = func() {
	Description("Key ID")
	Minimum(1)
	Example(10)
}

var fieldDatetime = func() {
	Description("Datetime")
	Format("date-time")
}

var TypeFooter = Type("TypeFooter", func() {
	Attribute("created_at", String, "Date of creation", fieldDatetime)
	Attribute("updated_at", String, "Last update date", fieldDatetime)
})
