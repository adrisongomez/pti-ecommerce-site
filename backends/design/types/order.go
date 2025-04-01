package types

import . "goa.design/goa/v3/dsl"

var OrderLineItemInput = Type("OrderLineItemInput", func() {
	Attribute("productId", Int)
	Attribute("price", Int)
	Attribute("quantity", Int)
	Required("productId", "price", "quantity")
})

var OrderInput = Type("OrderInput", func() {
	Attribute("email", String)
	Attribute("userId", Int)
	Attribute("lineItems", ArrayOf(OrderLineItemInput))
	Attribute("addressId", Int)
	Attribute("totalPrice", Int64, func() {
		Default(0)
	})
	Required("email", "lineItems", "addressId", "userId", "totalPrice")
})

var Order = ResultType("application/vnd.order+json", func() {
	ContentType("application/json")
	Attributes(func() {
		Attribute("id", Int, fieldID)
		Attribute("email", String)
		Attribute("user", User)
		Attribute("lineItems", ArrayOf(OrderLineItem))
		Attribute("address", Address)
		Attribute("totalPrice", String)
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
