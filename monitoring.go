package main

import (
	"github.com/goadesign/goa"
	"github.com/mtenrero/ATQ-Director/app"
)

// MonitoringController implements the monitoring resource.
type MonitoringController struct {
	*goa.Controller
}

// NewMonitoringController creates a monitoring controller.
func NewMonitoringController(service *goa.Service) *MonitoringController {
	return &MonitoringController{Controller: service.NewController("MonitoringController")}
}

// Ping runs the ping action.
func (c *MonitoringController) Ping(ctx *app.PingMonitoringContext) error {
	// MonitoringController_Ping: start_implement

	// Put your logic here

	return nil
	// MonitoringController_Ping: end_implement
}
