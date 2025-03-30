package design

import (
	"github.com/adrisongomez/pti-ecommerce-site/backends/design/securities"
	types "github.com/adrisongomez/pti-ecommerce-site/backends/design/types"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils/auth"
	. "goa.design/goa/v3/dsl"
)

var PaginatedMedia = types.PaginatedResult("media-list", types.Media)
var CreateMediaResponse = ResultType("application/vnd.create-media-response+json", func() {
	Attributes(func() {
		Attribute("media", types.Media)
		Attribute("uploadUrl", String)
		Required("media", "uploadUrl")
	})
})

var _ = Service(servicePrefix+"-media", func() {
	Description("Service perform CRUDs over media resource")
	HTTP(func() {
		Path("/medias")
		Response("Unauthorized", StatusUnauthorized)
	})

	Error("Unauthorized")
	Error("NotFound")
	Error("BadRequest")

	Method("list", func() {
		Description("Create a media record")
		Security(securities.JWTAuth, func() {
			Scope(auth.MediasReads)
		})

		Result(PaginatedMedia)
		Payload(func() {
			Token("token")
			Attribute("pageSize", Int, "Record per page", func() {
				Minimum(10)
				Maximum(100)
				Default(10)
			})
			Attribute("after", Int, "Start listing after this resource", func() {
				Default(0)
			})
			Attribute("bucket", String, "S3 bucket where data is store", func() {
				Default("")
			})
			Required("token")
		})

		HTTP(func() {
			GET("")
			Param("pageSize")
			Param("after")
			Param("bucket")
			Response(StatusOK)
			Response("BadRequest", StatusBadGateway)
		})
	})

	Method("getById", func() {
		Description("Get a media by id")
		Security(securities.JWTAuth, func() {
			Scope(auth.MediasReads)
		})

		Payload(func() {
			Token("token")
			Attribute("mediaId", Int)
			Required("token", "mediaId")
		})

		Result(types.Media)
		HTTP(func() {
			GET("/{mediaId}")
			Param("mediaId")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})

	Method("create", func() {
		Description("Create a media record")
		Payload(func() {
			Attribute("input", types.MediaInput)
			Token("token")
			Required("token", "input")
		})
		Result(CreateMediaResponse)
		Security(securities.JWTAuth, func() {
			Scope(auth.MediasWrite)
		})

		HTTP(func() {
			POST("")
			Body("input")
			Response(StatusCreated)
			Response("BadRequest", StatusBadRequest)
		})
	})

	Method("deleteById", func() {
		Description("Create a media record")
		Payload(types.MediaInput)
		Result(Boolean)
		Security(securities.JWTAuth, func() {
			Scope(auth.MediasWrite)
		})

		Payload(func() {
			Token("token")
			Attribute("mediaId", Int)
			Required("token", "mediaId")
		})

		HTTP(func() {
			DELETE("/{mediaId}")
			Param("mediaId")
			Response(StatusCreated)
			Response("NotFound", StatusNotFound)
		})
	})
})
