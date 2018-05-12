//go:generate goagen bootstrap -d github.com/mtenrero/ATQ-Director/http/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/mtenrero/ATQ-Director/app"
)

func main() {
	// Create service
	service := goa.New("ATQ - Director")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "databind" controller
	c := NewDatabindController(service)
	app.MountDatabindController(service, c)
	// Mount "monitoring" controller
	c2 := NewMonitoringController(service)
	app.MountMonitoringController(service, c2)
	// Mount "swarm" controller
	c3 := NewSwarmController(service)
	app.MountSwarmController(service, c3)
	// Mount "task" controller
	c4 := NewTaskController(service)
	app.MountTaskController(service, c4)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
