package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("ATQ - Director", func() { // "My API" is the name of the API used in docs
	Title("ATQ Test Orchestration Director")                         // Documentation title
	Description("REST Interface for Test queuing and orchestration") // Longer documentation description
	Host("atq.mtenrero.com")                                         // Host used by Swagger and clients
	Scheme("http")                                                   // HTTP scheme used by Swagger and clients
	BasePath("/api")                                                 // Base path to all API endpoints
	Consumes("application/json")                                     // Media types supported by the API
	Produces("application/json")                                     // Media types generated by the API
})
