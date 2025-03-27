package types

import (
	. "goa.design/goa/v3/dsl"
)

var UserRegistrationInput = Type("UserRegistrationInput", func() {
	Attribute("firstName", String)
	Attribute("lastName", String)
	Attribute("email", String, func() {
		Format("email")
	})
	Attribute("password", String)

	Required("firstName", "email", "password")
})

var UserCreateInput = Type("UserCreateInput", func() {
	Attribute("firstName", String)
	Attribute("lastName", String)
	Attribute("email", String, func() {
		Format("email")
	})
	Attribute("role", UserRole, func() {
		Default("CUSTOMER")
	})

	Required("firstName", "email")
})

var User = ResultType("application/vnd.user+json", func() {
	ContentType("application/json")
	Attributes(func() {
		Attribute("id", Int, fieldID)
		Attribute("firstName", String)
		Attribute("lastName", String)
		Attribute("email", String, func() {
			Format("email")
		})
		Attribute("role", UserRole, func() {
			Default("CUSTOMER")
		})
		Attribute("createdAt", String, "Date of creation", fieldDatetime)
		Attribute("updatedAt", String, "Last update date", fieldDatetime)

		Required("id", "firstName", "email", "createdAt", "role")
	})

})

var UserRole = Type("UserRole", String, func() {
	Enum("CUSTOMER", "ADMIN")
})
