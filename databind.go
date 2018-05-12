package main

import (
	"github.com/goadesign/goa"
	"github.com/mtenrero/ATQ-Director/app"
)

// DatabindController implements the databind resource.
type DatabindController struct {
	*goa.Controller
}

// NewDatabindController creates a databind controller.
func NewDatabindController(service *goa.Service) *DatabindController {
	return &DatabindController{Controller: service.NewController("DatabindController")}
}

// List runs the list action.
func (c *DatabindController) List(ctx *app.ListDatabindContext) error {
	// DatabindController_List: start_implement

	// Put your logic here

	res := app.AtqDatabindUploadCollection{}
	return ctx.OK(res)
	// DatabindController_List: end_implement
}

// Upload runs the upload action.
func (c *DatabindController) Upload(ctx *app.UploadDatabindContext) error {
	// DatabindController_Upload: start_implement

	// Put your logic here

	res := &app.AtqDatabindUpload{}
	return ctx.OK(res)
	// DatabindController_Upload: end_implement
}
