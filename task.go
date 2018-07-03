package main

import (
	"github.com/goadesign/goa"
	"github.com/mtenrero/ATQ-Director/app"
	"github.com/mtenrero/ATQ-Director/dockerMiddleware"
	"github.com/mtenrero/ATQ-Director/persistance"
)

// TaskController implements the task resource.
type TaskController struct {
	*goa.Controller
	*persistance.Persistance
}

// NewTaskController creates a task controller.
func NewTaskController(service *goa.Service, persistance *persistance.Persistance) *TaskController {
	return &TaskController{
		Controller:  service.NewController("TaskController"),
		Persistance: persistance,
	}
}

// Create runs the create action.
func (c *TaskController) Create(ctx *app.CreateTaskContext) error {
	task, err := dockerMiddleware.TaskMasterWorker(ctx.Payload, c.Persistance)
	if err != nil {
		errr := err.Error()
		errorResponse := app.AtqTask{
			Status: &errr,
		}
		return ctx.DefinitionError(&errorResponse)
	}

	return ctx.OKFull(task)
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
	status, err := c.Persistance.ReadTask(ctx.ID)

	if err != nil {
		return ctx.TaskNotIdentified()
	}

	res := &app.AtqTask{}
	res.ID = &ctx.ID
	res.Status = status.Status
	return ctx.OK(res)
}
