package main

import (
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
		swarmError := app.AtqSwarmError{
			Error: &errr,
		}

		return ctx.SwarmErrorError(&swarmError)
	}

	swarm := &app.AtqSwarm{
		JoinTokens: &app.JoinTokens{
			Manager: &inspect.JoinTokens.Manager,
			Worker:  &inspect.JoinTokens.Worker,
		},
	}

	return ctx.OK(swarm)
}
