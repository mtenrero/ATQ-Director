package design

var _ = Resource("swarm", func() {
	Action("status", func() {
		Routing(GET("/"))
		Description("Response with the details of the swarm")
		Response("Status", func() {
			Description("Details of the Docker Swarm cluster")
			Status(200)
		})
	})
})
