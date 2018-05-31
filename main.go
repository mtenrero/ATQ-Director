//go:generate goagen bootstrap -d github.com/mtenrero/ATQ-Director/http/design

package main

import (
	"os"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/mtenrero/ATQ-Director/app"
	"github.com/mtenrero/ATQ-Director/configLoader"
	"github.com/mtenrero/ATQ-Director/persistance"
)

// Persistance is the Global Persistance Instance
var Persistance *persistance.Persistance

const persistancePath = "./storage/datastore.atq"

func main() {
	// Create service
	service := goa.New("ATQ - Director")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Initialize Persistance Datastore
	var err error
	Persistance, err = persistance.InitPersistance(persistancePath)
	if err != nil {
		service.LogError("Error initializing datastore", "datastoreErr", err)
		os.Exit(-200)
	}

	// Load config
	config, err := configLoader.LoadControllerConfigYaml("./controller-config.yaml")
	if err != nil {
		service.LogError("Error loading config file", "configFile", err)
		os.Exit(-45)
	}

	// Mount "databind" controller
	c := NewDatabindController(service, Persistance)
	app.MountDatabindController(service, c)
	// Mount "monitoring" controller
	c2 := NewMonitoringController(service)
	app.MountMonitoringController(service, c2)
	// Mount "swarm" controller
	c3 := NewSwarmController(service)
	app.MountSwarmController(service, c3)
	// Mount "task" controller
	c4 := NewTaskController(service, Persistance)
	app.MountTaskController(service, c4)

	// Start service
	if err := service.ListenAndServe(":" + config.Port); err != nil {
		service.LogError("startup", "err", err)
	}

}
