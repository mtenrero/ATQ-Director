package design

var Service = Type("Service", func() {
	Attribute("image", String, func() {
		Description("Docker base image to attach to Service")
		Example(`hello-world`)
	})
	Attribute("args", []String, func() {
		Description("Arguments to be passed to the container")
		Example(`["-arg argvalue", ...]`)
	})
	Attribute("tty", bool, func() {
		Description("Interactive shell requirement")
		Example(`true`)
	})
	Attribute("alias", String, func() {
		Description("Service alias, this will identify the Service")
		Example(`ALIAS`)
	})
	Attribute("fileid",, func() {
		Description("ID of the Zipped contents that will be mounted and accesible inside the container, PREVIOUSLY UPLOADED")
	})
	Required("image", "alias")
})
