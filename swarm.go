package main

import (
	"github.com/goadesign/goa"
	"github.com/mtenrero/ATQ-Director/app"
)

// SwarmController implements the swarm resource.
type SwarmController struct {
	*goa.Controller
}

// NewSwarmController creates a swarm controller.
func NewSwarmController(service *goa.Service) *SwarmController {
	return &SwarmController{Controller: service.NewController("SwarmController")}
}

// Status runs the status action.
func (c *SwarmController) Status(ctx *app.StatusSwarmContext) error {
	// SwarmController_Status: start_implement

	// Put your logic here

	return nil
	// SwarmController_Status: end_implement
}
