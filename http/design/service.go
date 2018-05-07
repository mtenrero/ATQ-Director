package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var ServicePayload = Type("ServicePayload", func() {
	Attribute("image", String, func() {
		Description("Docker base image to attach to Service")
		Example(`hello-world`)
	})
	Attribute("replicas", Integer, func() {
		Description("Amount of replicas to be deployed. (1 by default)")
	})
	Attribute("args", ArrayOf(String), func() {
		Description("Arguments to be passed to the container")
	})
	Attribute("tty", Boolean, func() {
		Description("Interactive shell requirement")
		Example(true)
	})
	Attribute("alias", String, func() {
		Description("Service alias, this will identify the Service")
		Example(`ALIAS`)
	})
	Attribute("fileid", UUID, func() {
		Description("ID of the Zipped contents that will be mounted and accesible inside the container, PREVIOUSLY UPLOADED")
	})
	Required("image", "alias")
})
