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
	Attribute("fileid", String, func() {
		Description("ID of the Zipped contents that will be mounted and accesible inside the container, PREVIOUSLY UPLOADED")
	})
	Required("image", "alias")
})

var ServiceResponse = MediaType("application/atq.service+json", func() {
	Description("Created Relevant Service Information")
	Attributes(func() {
		Attribute("id", String, "Docker Service internal identifier")
		Attribute("fileId", String, "ATQ FileID if exists")
		Attribute("image", String, "Docker Image Name")
		Attribute("alias", String, "ATQ Service internal alias")
		Attribute("replicas", Integer, "Amount of Replicas")
		Attribute("tty", Boolean, "Interactive Shell")
		Attribute("args", ArrayOf(String), "Arguments passed to the containers")
	})

	View("default", func() {
		Attribute("id")
		Attribute("alias")
		Attribute("fileId")
	})

	View("minimal", func() {
		Attribute("id")
	})

	View("full", func() {
		Attribute("id")
		Attribute("fileId")
		Attribute("image")
		Attribute("alias")
		Attribute("replicas")
		Attribute("tty")
		Attribute("args")
	})
})
