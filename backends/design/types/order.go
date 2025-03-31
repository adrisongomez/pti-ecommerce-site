package types

import . "goa.design/goa/v3/dsl"

var OrderLineItemInput = Type("OrderLineItemInput", func() {
	Attribute("productId", String)
	Attribute("price", Int)
	Attribute("quantity", Int)
	Required("productId", "price", "quantity")
})

var OrderInput = Type("OrderInput", func() {
	Attribute("email", String)
	Attribute("userId", String)
	Attribute("lineItems", ArrayOf(OrderLineItemInput))
	Attribute("addressId", String)
	Required("email", "lineItems", "addressId")
})

var Order = ResultType("application/vnd.order+json", func() {
	ContentType("application/json")
	Attributes(func() {
		Attribute("id", Int, fieldID)
		Attribute("email", String)
		Attribute("user", User)
		Attribute("lineItems", ArrayOf(OrderLineItem))
		Attribute("address", Address)
		Attribute("totalPrice", Int)
		Attribute("createdAt", String, "Date of creation", fieldDatetime)
		Attribute("updatedAt", String, "Last update date", fieldDatetime)
		Required("id", "email", "user", "address", "lineItems", "createdAt", "totalPrice")
	})
})

var OrderLineItem = ResultType("application/vnd.order-line-item+json", func() {
	ContentType("application/json")
	Attributes(func() {
		Attribute("id", Int, fieldID)
		Attribute("productId", Int, fieldID)
		Attribute("price", String)
		Attribute("quantity", Int)
		Attribute("createdAt", String, "Date of creation", fieldDatetime)
		Attribute("updatedAt", String, "Last update date", fieldDatetime)
		Required("id", "price", "quantity", "createdAt")
	})
})
