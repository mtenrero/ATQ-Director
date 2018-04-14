package design

var _ = Resource("monitoring", func() {
	Action("ping", func() {
		Routing(GET("/ping"))
		Description("Endpoint for pinging and healthcheck purposes")
		Response("Alive", func() {
			Status(200)
		})
	})
})
