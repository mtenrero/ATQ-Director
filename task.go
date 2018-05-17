package main

import (
	"github.com/goadesign/goa"
	"github.com/mtenrero/ATQ-Director/app"
	"github.com/mtenrero/ATQ-Director/dockerMiddleware"
)

// TaskController implements the task resource.
type TaskController struct {
	*goa.Controller
}

// NewTaskController creates a task controller.
func NewTaskController(service *goa.Service) *TaskController {
	return &TaskController{Controller: service.NewController("TaskController")}
}

// Create runs the create action.
func (c *TaskController) Create(ctx *app.CreateTaskContext) error {
	task, err := dockerMiddleware.TaskMasterWorker(ctx.Payload)
	if err != nil {
		errr := err.Error()
		errorResponse := app.AtqTask{
			Status: &errr,
		}
		return ctx.DefinitionError(&errorResponse)
	}

	return ctx.OK(task)
	// TaskController_Create: end_implement
}

// Delete runs the delete action.
func (c *TaskController) Delete(ctx *app.DeleteTaskContext) error {
	// TaskController_Delete: start_implement

	// Put your logic here

	return nil
	// TaskController_Delete: end_implement
}

// Inspect runs the inspect action.
func (c *TaskController) Inspect(ctx *app.InspectTaskContext) error {
	// TaskController_Inspect: start_implement

	// Put your logic here

	res := &app.AtqTask{}
	return ctx.OK(res)
	// TaskController_Inspect: end_implement
}
