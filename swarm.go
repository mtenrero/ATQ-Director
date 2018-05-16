package main

import (
	"encoding/json"

	"github.com/goadesign/goa"
	"github.com/mtenrero/ATQ-Director/app"
	"github.com/mtenrero/ATQ-Director/dockerMiddleware"
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
	inspect, err := dockerMiddleware.SwarmInspect()
	if err != nil {
		errr := err.Error()
		return ctx.Status([]byte(errr))
	}

	inspectJSON, _ := json.Marshal(inspect)

	return ctx.Status([]byte(inspectJSON))
}
