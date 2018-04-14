package design

var _ = Resource("databind", func() {
	Action("upload", func() {
		Routing(POST("/"))
		Description("Upload new zipped file for later usage with a Task")
		Response("OK", func() {
			Description("The file was uploaded succesfully")
			Status(200)
			Media(Upload)
		})
		Response("Upload error", func() {
			Description("Response when there are an error uploading the file")
			Status(500)
			Media(UploadError)
		})
		Response("The file doesn't have an accepted compression", func() {
			Description("The file doesn't have a valid extension")
			Status(415)
			Media(UploadError)
		})
	})

	Action("list", func() {
		Routing(GET("/list"))
		Description("List of uploaded and available files")
		Response("OK", CollectionOf(Upload))
		Response(NoContent)
	})
})

var Upload = Type("Upload", func() {
	Attribute("id", String, func() {
		Description("Internal file identifier")
	})
})

var UploadError = Type("UploadError", func() {
	Attribute("error", String, func() {
		Description("Error description")
	})
})
