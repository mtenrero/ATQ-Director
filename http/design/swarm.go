package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("swarm", func() {
	BasePath("/swarm")
	Action("status", func() {
		Routing(GET("/"))
		Description("Response with the details of the swarm")
		Response("Swarm Error", func() {
			Description("Docker Swarm context Error Message")
			Status(503)
			Media(Swarm, "error")
		})
		Response("OK", func() {
			Description("Details of the Docker Swarm cluster")
			Status(200)
			Media(Swarm)
		})
	})
})

var Swarm = MediaType("application/atq.swarm+json", func() {
	Description("Swarm Details")
	Attributes(func() {
		Attribute("error", String, "Swarm Error Message")
		Attribute("joinTokens", JoinTokens)
	})
	View("default", func() {
		Attribute("joinTokens")
	})
	View("error", func() {
		Attribute("error")
	})
})

var JoinTokens = Type("JoinTokens", func() {
	Description("Docker Swarm Join Tokens")
	Attribute("worker", String)
	Attribute("manager", String)
})
