package design

import (
	. "goa.design/goa/v3/dsl"
)

var HealthcheckResponse = ResultType("application/vnd.healthcheck-response+json", func() {
	Attributes(func() {
		Attribute("status", String)
	})
})

var _ = Service(servicePrefix+"-healthcheck", func() {
	Method("check", func() {
		Result(HealthcheckResponse)
		HTTP(func() {
			GET("/healthcheck")
			Response(StatusOK)
		})
	})
})
