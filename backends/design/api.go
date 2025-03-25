package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var (
	domain        = "localhost"
	baseAPIDir    = "/api"
	servicePrefix = "svc"
)

var _ = API("ecommerce", func() {
	Title("Ecommerce API")
	Description("Ecommerce API Design to Programming The Internet Class at BIU University")
	Contact(func() {
		Name("Adrison Gomez")
		Email("info@adrisongomez.me")
		URL("https://github.com/adrisongomez")
	})

	cors.Origin("/.*localhost.*/", func() {
		cors.Methods("GET", "POST", "PUT", "DELETE")
	})

	Server("ecommerce", func() {
		Host("localhost", func() {
			URI("http://localhost:3030")
		})
	})

	HTTP(func() {
		Consumes("application/json")
		Path(baseAPIDir)
	})
})
