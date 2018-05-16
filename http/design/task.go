package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("task", func() {
	BasePath("/task")
	Action("create", func() {
		Routing(PUT("/"))
		Description("Creates a new Task in the Swarm according with the config provided in the JSON body")
		Payload(TaskPayload)
		Response("OK", func() {
			Description("Task creation in progress")
			Status(200)
			Media(Task)
		})
		Response("Definition error", func() {
			Description("The Task definition has errors or it's not complete")
			Status(417)
			Media(Task)
		})
	})
	Action("inspect", func() {
		Routing(GET("/:id"))
		Description("Get Task's details")
		Params(func() {
			Param("id", UUID, "Task's UUID")
		})
		Response("OK", func() {
			Description("Successful response containing Task data in JSON format")
			Status(200)
			Media(Task)
		})
		Response("Task not identified", func() {
			Description("The given ID doesn't not exist")
			Status(404)
		})
		Response("Error Creating the Task", func() {
			Description("Response when the Task has not been created correctly")
			Status(500)
			Media(Task)
		})
	})
	Action("delete", func() {
		Routing(DELETE("/:id"))
		Description("Deletes the Task specified and its components")
		Response("OK", func() {
			Description("Successfuly deleted")
			Status(204)
		})
		Response("Task not identified", func() {
			Description("The given ID doesn't not exist")
			Status(404)
		})
		Response("Task could not be deleted", func() {
			Description("Docker Engine error deleting the Task generated container infrastructure")
			Status(500)
		})
	})
})

var TaskPayload = Type("TaskPayload", func() {
	Attribute("master", ServicePayload)
	Attribute("worker", ServicePayload)
	Attribute("waitCommand", WaitCommand)
	Attribute("delay", Integer, func() {
		Minimum(0)
	})
})

var Task = MediaType("application/atq.task+json", func() {
	Description("Task description")
	Reference(TaskPayload)
	Attributes(func() {
		Attribute("id", String, "Task ID")
		Attribute("status", func() {
			Enum("initializing", "started", "stopped", "finished", "errored")
			Description("Status of the Task")
		})

		Attribute("master")
		Attribute("worker")
		Attribute("waitCommand")
		Attribute("delay")
	})
	View("default", func() {
		Attribute("id")
		Attribute("status")
	})
	View("full", func() {
		Attribute("id")
		Attribute("status")
		Attribute("master")
		Attribute("worker")
		Attribute("waitCommand")
		Attribute("delay")
	})
})

var WaitCommand = Type("WaitCommand", func() {
	Description("Definition of a command to be executed")
	Attribute("command", String, "Command to be executed")
	Attribute("expectedResult", String, "Result expected")
	Attribute("timeout", Integer, "Maximum seconds to wait until succesfull response")
})
