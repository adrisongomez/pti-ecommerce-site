package types

import (
	. "goa.design/goa/v3/dsl"
)

var RegisterUserInput = Type("RegisterUserInput", func() {
	Attribute("firstName", String)
	Attribute("lastName", String)
	Attribute("email", String, func() {
		Format("email")
	})
	Attribute("password", String)

	Required("firstName", "email", "password")
})

var CreateUserInput = Type("CreateUserInput", func() {
	Attribute("firstName", String)
	Attribute("lastName", String)
	Attribute("email", String, func() {
		Format("email")
	})
	Attribute("role", UserRole, func() {
		Default("USER")
	})
	Attribute("password", String)

	Required("firstName", "email", "password")
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
		Attribute("role", UserRole)
		Attribute("createdAt", String, "Date of creation", fieldDatetime)
		Attribute("updatedAt", String, "Last update date", fieldDatetime)

		Required("id", "firstName", "email", "createdAt", "role")
	})

})

var UserRole = Type("UserRole", String, func() {
	Enum("CUSTOMER", "ADMIN")
})
