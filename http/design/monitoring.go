package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("monitoring", func() {
	BasePath("/monitoring")
	Action("ping", func() {
		Routing(GET("/ping"))
		Description("Endpoint for pinging and healthcheck purposes")
		Response("Alive", func() {
			Description("Pong")
			Status(200)
		})
	})
})
