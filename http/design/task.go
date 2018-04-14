package design

var _ = Resource("task", func() {
	Action("create", func() {
		Routing(PUT("/"))
		Description("Creates a new Task in the Swarm according with the config provided in the JSON body")
		Payload(Task)
		Response("OK", func() {
			Description("Task creation in progress")
			Status(200)
			Media(TaskCreation)
		})
		Response("Definition error", func() {
			Description("The Task definition has errors or it's not complete")
			Status(417)
			Media(TaskCreation)
		})
	})
	Action("inspect", func() {
		Routing(GET("/:id"))
		Description("Get Task's details")
		Response("OK", func() {
			Description("Successful response containing Task data in JSON format")
			Status(200)
			Media(TaskCreation)
		})
		Response("Task not identified", func() {
			Description("The given ID doesn't not exist")
			Status(404)
		})
		Response("Error Creating the Task", func() {
			Description("Response when the Task has not been created correctly")
			Status(500)
			Media(TaskCreation)
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

var TaskCreation = Type("TaskCreation", func() {
	Description("Task Creation information")
	Attributes(func() {
		Attribute("id", String, "Internal Task ATQ ID")
		Attribute("status", func() {
			Enum("creating", "waitingWorkers", "running", "finished", "errored")
		})
		Attribute("error", String, "Error message if errored")
		Attribute("task", Task, "Task")
	})
})

var Task = Type("Task", func() {
	Description("Task definition")
	Attributes(func() {
		Attribute("type", func() {
			Enum("masterworker")
		})
		Attribute("masterworker", MasterWorkerTask, "MasterWorker Task definition")
	})
})

var MasterWorkerTask = Type("MasterWorkerTask", func() {
	Description("Master/Workers Task definition")
	Attribute("worker", func() {
		Attribute("service", Service, "WorkerService definition")
		Attribute("replicas", Integer, "Amount of workers to deploy in the Swarm cluster")
	})
	Attribute("master", func() {
		Attribute("service", Service, "MasterService definition")
		Attribute("delay", Integer, "Amount of seconds to wait until container creation")
		Attribute("waitCommand", WaitCommand, "WaitCommand definition to be launched on the server before master container creation")
	})
})

var WaitCommand = Type("WaitCommand", func() {
	Description("Definition of a command to be executed")
	Attribute("command", String, "Command to be executed")
	Attribute("expectedResult", String, "Result expected")
	Attribute("timeout", Integer, "Maximum seconds to wait until succesfull response")
})
