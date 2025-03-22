package design

import (
	"github.com/adrisongomez/pti-ecommerce-site/backends/design/types"
	. "goa.design/goa/v3/dsl"
)

var PaginatedVendor = types.PaginatedResult("vendor-list", types.Vendor)

var _ = Service(servicePrefix+"-vendor", func() {
	Description("The product service perform CRUD over the vendor resource")

	HTTP(func() {
		Path("/vendors")
	})

	Error("NotFound")
	Error("BadRequest")
	Error("Conflict")

	Method("list", func() {
		Description("List vendors")
		Payload(func() {
			Attribute("pageSize", Int, "Record per page", func() {
				Minimum(10)
				Maximum(100)
				Default(10)
			})
			Attribute("after", Int, "Start listing after this resource")
		})

		Result(PaginatedVendor)

		HTTP(func() {
			GET("")
			Param("pageSize")
			Param("after")
			Response(StatusOK)
			Response("BadRequest", StatusBadRequest)
		})
	})

	Method("create", func() {
		Description("Create a new product")

		Payload(types.VendorInput)
		Result(types.Vendor)

		HTTP(func() {
			POST("")
			Response(StatusCreated)
		})
	})

	Method("deleteById", func() {
		Description("Create a new product")

		Payload(func() {
			Attribute("vendorId", Int, "Unique product identifier")
			Required("vendorId")
		})
		Result(Boolean)

		HTTP(func() {
			PUT("/{vendorId}")
			Param("vendorId")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})
})
