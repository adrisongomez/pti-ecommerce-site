package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("openapi", func() {
	Files("/openapi.json", "internal/gen/http/openapi.json")
})
