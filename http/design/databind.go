package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("databind", func() {
	Action("upload", func() {
		Routing(POST("/"))
		Description("Upload new zipped file for later usage with a Task")
		Payload(UploadPayload)
		MultipartForm()
		Response("OK", func() {
			Description("The file was uploaded succesfully")
			Status(200)
			Media(Upload)
		})
		Response("Upload error", func() {
			Description("Response when there are an error uploading the file")
			Status(500)
			Media(Upload, "error")
		})
		Response("The file doesn't have an accepted compression", func() {
			Description("The file doesn't have a valid extension")
			Status(415)
			Media(Upload, "error")
		})
	})

	Action("list", func() {
		Routing(GET("/list"))
		Description("List of uploaded and available files")
		Response("OK", CollectionOf(Upload))
		Response(NoContent)
	})
})

var UploadPayload = Type("UploadPayload", func() {
	Attribute("file", File, "Zipped File")
})

var Upload = MediaType("application/atq.databind.upload+json", func() {
	Reference(UploadPayload)
	Description("User upload files response")
	Attributes(func() {
		Attribute("id", UUID, "Upload ID")
		Attribute("error", String, "Error message if errored")
	})
	View("default", func() {
		Attribute("id")
	})
	View("error", func() {
		Attribute("error")
	})
})
