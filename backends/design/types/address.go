package types

import (
	. "goa.design/goa/v3/dsl"
)

var AddressInput = Type("AddressInput", func() {
	Attribute("addressLine1", String)
	Attribute("addressLine2", String)
	Attribute("city", String)
	Attribute("state", String)
	Attribute("country", String)
	Attribute("zipCode", String)
})

var Address = ResultType("application/vnd.address+json", func() {
	ContentType("application/json")
	Attributes(func() {
		Attribute("id", Int, fieldID)
		Attribute("user", User)
		Attribute("addressLine1", String)
		Attribute("addressLine2", String)
		Attribute("city", String)
		Attribute("state", String)
		Attribute("country", String)
		Attribute("zipCode", String)
		Attribute("createdAt", String, "Date of creation", fieldDatetime)
		Attribute("updatedAt", String, "Last update date", fieldDatetime)

		Required("id", "user", "addressLine1", "city", "state", "country", "zipCode", "createdAt")
	})
})
