package main

import (
	"os"

	"github.com/mtenrero/ATQ-Director/configLoader"
	"github.com/mtenrero/ATQ-Director/dockerMiddleware"
	"github.com/sirupsen/logrus"
)

// LOGGER is the Logrus logger global instance
var LOGGER *logrus.Logger

// CONFIG holds the application configuration
var CONFIG configLoader.ControllerConfig

func init() {
	logrus.SetOutput(os.Stdout)
	LOGGER = logrus.New()
}

func main() {
	CONFIG, err := configLoader.LoadControllerConfigYaml("controller-config.yaml")
	if err != nil {
		LOGGER.Fatal(err)
	}

	LOGGER.Info(CONFIG.Port)

	swarmChecks()
}

func swarmChecks() {
	if !dockerMiddleware.IsSwarmNode() {
		_, error := dockerMiddleware.SwarmInspect()
		LOGGER.Error(error)
	}
}
